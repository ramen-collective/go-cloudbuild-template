package dataloader

import (
	"context"
	"net/http"

	"github.com/ramen-collective/go-cloudbuild-template/internal/repository"
)

type contextKey string

const loadersKey contextKey = "dataloaders"

// Loaders struct contains all app dataloaders
type Loaders struct {
	User UserLoader
}

// Middleware represents the dataloader middleware
type Middleware struct {
	repositories *repository.Container
}

// NewMiddleware instantiates a new Middleware
func NewMiddleware(repositories *repository.Container) *Middleware {
	return &Middleware{
		repositories: repositories,
	}
}

// Handle attach a Loaders instance to the request context
func (m Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := BindLoadersToContext(r.Context(), m.repositories)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

// BindLoadersToContext add loaders instance to a given context
func BindLoadersToContext(ctx context.Context, repositories *repository.Container) context.Context {
	return context.WithValue(ctx, loadersKey, &Loaders{
		User: NewUserLoader(repositories.User),
	})
}

// For retrieves the Loaders instance for the given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
