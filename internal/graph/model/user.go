package model

import (
	"github.com/ramen-collective/go-cloudbuild-template/internal/repository"
	"github.com/ramen-collective/go-cloudbuild-template/pkg/util/gqlutil"
)

// User represents the User type in GraphQL
type User struct {
	UUID gqlutil.UUID `json:"uuid"`
	Name string       `json:"name"`
}

// NewUserFromDB create a new User from a repository.User struct
func NewUserFromDB(user repository.User) *User {
	return &User{
		UUID: gqlutil.UUID(user.UUID),
		Name: user.Name,
	}
}

// UserEdge represents the UserEdge type in GraphQL
type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

// NewUserEdgesFromModel creates new []*UserEdge from a []*User
func NewUserEdgesFromModel(users []*User) []*UserEdge {
	edges := make([]*UserEdge, len(users))
	for i, user := range users {
		edges[i] = &UserEdge{
			Cursor: gqlutil.EncodeUUIDCursor(user.UUID.String()),
			Node:   user,
		}
	}
	return edges
}
