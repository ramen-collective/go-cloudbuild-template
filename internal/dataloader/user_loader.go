//go:generate go run github.com/vektah/dataloaden UserByUUIDLoader string "*github.com/ramen-collective/go-cloudbuild-template/internal/graph/model.User"
package dataloader

import (
	"time"

	"github.com/ramen-collective/go-cloudbuild-template/internal/graph/model"

	"github.com/ramen-collective/go-cloudbuild-template/internal/repository"
)

// UserLoader struct contains all User dataloaders
type UserLoader struct {
	GetByUUID UserByUUIDLoader
}

// NewUserLoader returns a new instance of UserLoader
func NewUserLoader(userRepository repository.UserRepositoryInterface) UserLoader {
	return UserLoader{
		GetByUUID: UserByUUIDLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch:    batchUserByUUID(userRepository),
		},
	}
}

func batchUserByUUID(userRepository repository.UserRepositoryInterface) func(keys []string) ([]*model.User, []error) {
	return func(uuids []string) ([]*model.User, []error) {
		dbUsers, err := userRepository.GetByUUIDs(uuids)
		if err != nil {
			return nil, []error{err}
		}
		userByID := map[string]*model.User{}
		for _, user := range dbUsers {
			userByID[user.UUID] = model.NewUserFromDB(user)
		}
		users := make([]*model.User, len(uuids))
		for i, uuid := range uuids {
			users[i] = userByID[uuid]
		}
		return users, nil
	}
}
