package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/ramen-collective/go-cloudbuild-template/internal/client"
	"github.com/ramen-collective/go-cloudbuild-template/internal/dataloader"
	"github.com/ramen-collective/go-cloudbuild-template/internal/graph/generated"
	"github.com/ramen-collective/go-cloudbuild-template/internal/graph/model"
	"github.com/ramen-collective/go-cloudbuild-template/internal/repository"
	"github.com/ramen-collective/go-cloudbuild-template/pkg/util/gqlutil"
	"github.com/ramen-collective/go-cloudbuild-template/pkg/util/translation"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

//"github.com/Celtcoste/server-graphql/graph/generated"

const (
	complexityExtensionKey = "cost" // named cost to be the same as PIPE
	queryComplexityLimit   = 600
)

// NewGraphQLServer instanciates a new gqlgen server from a config and a logger.
func NewGraphQLServer(config generated.Config, repositories *repository.Container) *handler.Server {
	ApplyGraphQLCustomComplexityCalculation(&config)
	DefineDirectives(&config, repositories)
	server := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	HandleGraphQLServerError(server)
	server.Use(extension.FixedComplexityLimit(queryComplexityLimit))
	var mb int64 = 1 << 20
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})
	server.Use(extension.Introspection{})
	return server
}

// SetupGraphQLRoutes register HTTP handlers related to GraphQL/gqlgen.
func SetupGraphQLRoutes(
	//configuration *authentication.Configuration,
	repositories *repository.Container,
	servePlayground bool,
	server *handler.Server,
	router *mux.Router,
) {
	if servePlayground {
		router.Handle("/", playground.Handler("GraphQL playground", "/query")).Methods("GET")
	}
	gqlrouter := router.PathPrefix("/query").Subrouter()
	//authMiddlware := authentication.NewMiddleware(configuration, repositories.User)
	gqlrouter.Use(mux.MiddlewareFunc(client.NewClientMiddleware()))
	//gqlrouter.Use(authMiddlware.Handle)
	gqlrouter.Use(translation.ParseAcceptLanguageMiddleware)
	dataloaderMiddlware := dataloader.NewMiddleware(repositories)
	gqlrouter.Use(dataloaderMiddlware.Handle)
	//gqlrouter.Use(mux.MiddlewareFunc(tracing.NewSpanMiddleware("graphql_query")))
	gqlrouter.Handle("", server).Methods("POST")
	restRouter := router.PathPrefix("/auth").Subrouter()
	restRouter.Handle("/create-user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//CreateUserRest(w, r, repositories)
	})).Methods("PUT")
	//restRouter.Use(authMiddlware.Handle)
}

// HandleGraphQLServerError defines how graphql errors should be
// handled and logged.
func HandleGraphQLServerError(server *handler.Server) {
	server.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		span := trace.SpanFromContext(ctx)
		operationContext := graphql.GetOperationContext(ctx)
		span.SetAttributes(
			attribute.String(
				"/graphql/operation/name",
				operationContext.OperationName,
			),
			attribute.String(
				"/graphql/path",
				graphql.GetPath(ctx).String(),
			),
		)
		for key, value := range operationContext.Variables {
			span.SetAttributes(
				attribute.String(
					fmt.Sprintf("/graphql/variables/%s", key),
					fmt.Sprintf("%v", value),
				),
			)
		}
		return next(ctx)
	})
	server.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		stats := extension.GetComplexityStats(ctx)
		if stats != nil {
			graphql.RegisterExtension(ctx, complexityExtensionKey, stats.Complexity)
		}
		return next(ctx)
	})
	server.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		switch e := err.(type) {
		case error:
			return gqlutil.InternalServerError(e)
		case string:
			return gqlutil.InternalServerErrorf(e)
		default:
			return gqlutil.InternalServerErrorf("%v", e)
		}
	})
	server.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		span := trace.SpanFromContext(ctx)
		span.RecordError(err)
		span.SetStatus(codes.Error, "graphql error")
		oc := graphql.GetOperationContext(ctx)
		path := graphql.GetPath(ctx).String()
		code := gqlutil.GetErrorCode(err)
		span.SetAttributes(
			attribute.String(
				"/graphql/error/code",
				string(code),
			),
		)
		errFields := []interface{}{
			"code", gqlutil.GetErrorCode(err),
			"operation_name", oc.OperationName,
			"path", path,
			"variables", oc.Variables,
		}
		errMessage := fmt.Sprintf("GraphQL: %s", err)
		log.Println("Error = ", errMessage, " - ", errFields)
		return gqlutil.ErrorPresenter(ctx, err)
	})
}

// DefineDirectives implements Graph directives
func DefineDirectives(config *generated.Config, repositories *repository.Container) {
	config.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		/*user := authentication.UserForContext(ctx)
		if user == nil {
			return nil, gqlutil.UnauthenticatedErrorf("Cannot perform operation for unauthenticated user")
		}*/
		return next(ctx)
	}

	config.Directives.RateLimitByUser = func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
		limit int,
		duration int) (interface{}, error) {
		/*user := authentication.UserForContext(ctx)
		if user == nil {
			return nil, gqlutil.UnauthenticatedErrorf("Cannot perform operation for unauthenticated user")
		}*/
		return next(ctx)
	}

	config.Directives.Locale = func(ctx context.Context, obj interface{}, next graphql.Resolver, lang model.Languages) (interface{}, error) {
		return next(context.WithValue(ctx, "locale", &lang))
	}
}

// ApplyGraphQLCustomComplexityCalculation handles specific query complexity
// calculation.
// Complexity values:
// - Basic field: 1
// - Field with extra sql queries: 1 + 3 per additional query
// - Paginated lists: multiply by number of item requested
func ApplyGraphQLCustomComplexityCalculation(config *generated.Config) {
	const extraSQLQueryCost = 3
	/*paginationComplexityFunc := func(childComplexity int, first int, after *string) int {
		return extraSQLQueryCost + childComplexity*first
	}*/

	// Game
	//config.Complexity.Game.Contents = paginationComplexityFunc

	// Medias
	/*config.Complexity.Media.Urls = func(childComplexity int) int {
		return 3 + childComplexity // generating signed urls
	}*/

	// Query
	/*config.Complexity.Query.Collections = func(childComplexity int, mode *model.CollectionMode, first int, after *string) int {
		return extraSQLQueryCost + childComplexity*first
	}*/

	// PrivateUser
	//config.Complexity.PrivateUser.InProgressContents = paginationComplexityFunc

	// Recommendations
	/*config.Complexity.Recommendations.ContentsByTags = func(childComplexity int, filter []string, exclude []string, prioritize []string, first int, after *string) int {
		return 3*extraSQLQueryCost + childComplexity*first // 3 extra sql queries
	}*/

	// Tags
	/*config.Complexity.Tag.Contents = func(childComplexity int, orderBy *model.TagContentOrder, first int, after *string) int {
	  	return extraSQLQueryCost + childComplexity*first
	  }
	  config.Complexity.Tag.Collections = func(childComplexity int, mode *model.CollectionMode, orderBy *model.TagCollectionOrder, first int, after *string) int {
	  	return extraSQLQueryCost + childComplexity*first
	  }*/
}
