package model

import (
	"github.com/ramen-collective/go-cloudbuild-template/pkg/util/gqlutil"
)

// Language represents the Language type in GraphQL
type Language struct {
	ID        gqlutil.ID `json:"id"`
	Name      string     `json:"name"`
	ImageUUID gqlutil.UUID
}

// NewLanguageFromDB create a new Language from a repository.Language struct
/*func NewLanguageFromDB(language repository.Language) *Language {
	return &Language{
		ID:        gqlutil.ID(language.ID),
		Name:      language.Name,
		ImageUUID: language.ImageUUID,
	}
}*/
