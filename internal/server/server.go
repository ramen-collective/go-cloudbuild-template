package server

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/ramen-collective/go-cloudbuild-template/internal/util"
	"github.com/ramen-collective/go-cloudbuild-template/pkg/cloud/cloudLogger"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// ApiTemplateServer is the application entrypoint
type ApiTemplateServer struct {
	configuration *Configuration
	Router        *mux.Router
}

// NewApiTemplateServer create a server instance with a router
func NewApiTemplateServer(configuration *Configuration) *ApiTemplateServer {
	router := mux.NewRouter()
	router.Use(cloudLogger.ParseTraceContextMiddleware)
	return &ApiTemplateServer{
		configuration: configuration,
		Router:        router,
	}
}

// Start the server and wait for gracefull shutdown
func (server *ApiTemplateServer) Start(ctx context.Context) {
	// New h2s server. As we want to use http/2 but cloud run
	// terminates tls before hitting the container, we must handle
	// requests in cleartext (h2c)
	// @see https://cloud.google.com/run/docs/configuring/http2#before_you_configure
	addr := server.configuration.Addr()
	h2server := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(server.Router, &http2.Server{}),
	}
	go func() {
		if err := h2server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	errChan := make(chan error, 10)
	go func() {
		util.RunHealthServer(flag.String("health", "0.0.0.0:"+server.configuration.HealthPort, "Health service address."), errChan)
	}()
	log.Printf("api-template server listening @ http://%s/", addr)
	// NOTE: Block and listen to interruption signals
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINT)
	log.Println("Catch signal:", <-exit)
	// NOTE: Shutdown the server with a given timeout before kill
	ctx, cancel := context.WithTimeout(ctx, server.configuration.ShutdownTimeout)
	defer cancel()
	if err := h2server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v\n", err)
	}
	log.Printf("api-template server stopped")
}
