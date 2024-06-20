//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/ramen-collective/go-cloudbuild-template/internal/repository"
	"github.com/ramen-collective/go-cloudbuild-template/pkg/util/translation"
)

// This file will not be regenerated automatically.

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	Repositories *repository.Container
	Translator   translation.TranslatorInterface
}

// NewResolver instantiate a Resolver
func NewResolver(repos *repository.Container, translator translation.TranslatorInterface) *Resolver {
	return &Resolver{
		Repositories: repos,
		Translator:   translator,
	}
}
