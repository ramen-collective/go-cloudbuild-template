package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	_insertUserQuery = `
		INSERT INTO user (uuid, name, created_at, updated_at)
		VALUES (?, ?, ?, ?);`
	_selectAllUserUUIDsQuery = `
		SELECT uuid
		FROM user
		ORDER BY created_at DESC;`
	_selectUserByUUIDQuery = `
		SELECT uuid, name, created_at, updated_at
		FROM user
		WHERE uuid = ?
		LIMIT 1;`
	_selectUserByNameQuery = `
		SELECT uuid, name, created_at, updated_at
		FROM user
		WHERE name = ?
		LIMIT 1;`
	_selectUsersByUUIDsQuery = `
		SELECT uuid, name, created_at, updated_at
		FROM user
		WHERE uuid IN (?);`
)

// UserRepositoryInterface should be implemented by UserRepository
type UserRepositoryInterface interface {
	Create(name string) (User, error)
	DeleteByUUID(uuid string) (bool, error)
	GetByUUID(uuid string) (User, error)
	GetByName(name string) (User, error)
	GetByUUIDs(uuids []string) ([]User, error)
	GetAllUUIDs() ([]string, error)
	UpdateByUUID(uuid string, name *string) (bool, error)
}

// User struct reflects database user table.
type User struct {
	UUID         string    `db:"uuid" json:"uuid"`
	Name         string    `db:"name" json:"name"`
	CreationDate time.Time `db:"created_at" json:"creationDate"`
	UpdateDate   time.Time `db:"updated_at" json:"updateDate"`
}

// UserRepository handle user data access
type UserRepository struct {
	DB *sqlx.DB
}

// NewUserRepository instantiate a new UserRepository
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Create an user entry in the database.
func (r *UserRepository) Create(
	name string,
) (User, error) {
	id := uuid.New()
	now := time.Now().UTC()
	var err error
	_, err = r.DB.Exec(
		_insertUserQuery,
		id,
		name,
		now,
		now,
	)
	user := User{
		UUID:         id.String(),
		Name:         name,
		CreationDate: now,
		UpdateDate:   now,
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteByUUID delete the user denoted by the given identifier.
func (r *UserRepository) DeleteByUUID(uuid string) (bool, error) {
	return DeleteByUUID(r.DB, "user", uuid)
}

// GetByUUID retrieve one User by its uuid from the database
func (r *UserRepository) GetByUUID(uuid string) (User, error) {
	user := User{}
	err := r.DB.Get(&user, _selectUserByUUIDQuery, uuid)
	return user, err
}

// GetByName retrieve one User by its pseudo from the database
func (r *UserRepository) GetByName(name string) (User, error) {
	user := User{}
	err := r.DB.Get(&user, _selectUserByNameQuery, name)
	return user, err
}

// GetByUUIDs retrieve multiple Users by given uuids from the database
func (r *UserRepository) GetByUUIDs(uuids []string) ([]User, error) {
	query, args, err := sqlx.In(_selectUsersByUUIDsQuery, uuids)
	if err != nil {
		return nil, err
	}
	var users []User
	err = r.DB.Select(&users, query, args...)

	return users, err
}

// GetAllUUIDs retrieve all User UUIDs from the database
func (r *UserRepository) GetAllUUIDs() ([]string, error) {
	var userUUIDs []string
	err := r.DB.Select(&userUUIDs, _selectAllUserUUIDsQuery)
	return userUUIDs, err
}

// UpdateByUUID update the user denoted by the given identifier.
func (r *UserRepository) UpdateByUUID(uuid string, name *string) (bool, error) {
	fields := make(map[string]interface{})

	if name != nil {
		fields["name"] = name
	}

	fields["updated_at"] = time.Now().UTC()
	return UpdateByUUID(r.DB, "user", uuid, fields)
}
