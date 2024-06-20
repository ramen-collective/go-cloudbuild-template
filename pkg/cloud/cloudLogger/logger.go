package cloudLogger

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"cloud.google.com/go/compute/metadata"
	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

var traceContextKey contextKey = "traceContext"

const (
	messageKey         string = "message"
	levelKey           string = "severity"
	traceContextHeader string = "X-Cloud-Trace-Context"
)

// LoggerInterface should be implemented by Logger
type LoggerInterface interface {
	WrapTraceContext(ctx context.Context) *zap.SugaredLogger
}

// Logger represents a zap logger
type Logger struct {
	*zap.Logger
	googleProjectID string
}

// TraceContext represents a trace
type TraceContext struct {
	TraceID   string
	SpanID    string
	IsSampled bool
}

// NewLogger instanciates a new app logger.
// It relies on the gcp metadata service to know if we're on a gcp
// env and to retrieve the project id.
func NewLogger() (*Logger, error) {
	var config zap.Config
	var err error
	var projectID string
	if metadata.OnGCE() {
		projectID, err = metadata.ProjectID()
		if err != nil {
			return nil, fmt.Errorf("Logger error: %v", err)
		}
		config = zap.NewProductionConfig()
		config.EncoderConfig.MessageKey = messageKey
		config.EncoderConfig.LevelKey = levelKey
	} else {
		projectID = ""
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}
	zapLogger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("Logger error: %v", err)
	}
	return &Logger{
		Logger:          zapLogger,
		googleProjectID: projectID,
	}, nil
}

// WrapTraceContext adds trace request context to the logger
func (l *Logger) WrapTraceContext(ctx context.Context) *zap.SugaredLogger {
	traceContext, _ := ctx.Value(traceContextKey).(*TraceContext)
	fields := zapdriver.TraceContext(traceContext.TraceID, traceContext.SpanID, traceContext.IsSampled, l.googleProjectID)
	setFields := l.With(fields...)
	return setFields.Sugar()
}

// ParseTraceContext parses X-Cloud-Trace-Context HTTP header
func ParseTraceContext(r *http.Request) *TraceContext {
	traceContextHeader := r.Header.Get(traceContextHeader)
	// Parse cloud trace context header that looks like:
	// "X-Cloud-Trace-Context: 105445aa7843bc8bf206b120001000/1;o=1"
	matches := regexp.MustCompile(`([a-f\d]+)?(?:/([a-f\d]+))?(?:;o=(\d))?`).FindStringSubmatch(traceContextHeader)
	traceID, spanID, isSampled := matches[1], matches[2], matches[3] == "1"
	return &TraceContext{
		TraceID:   traceID,
		SpanID:    spanID,
		IsSampled: isSampled,
	}
}

// ParseTraceContextMiddleware parse X-Cloud-Trace-Context HTTP header
// and store a TraceContext into context
func ParseTraceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceContext := ParseTraceContext(r)
		ctx := context.WithValue(r.Context(), traceContextKey, traceContext)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
