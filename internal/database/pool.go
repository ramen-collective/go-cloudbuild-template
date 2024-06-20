package database

import (
	_ "github.com/go-sql-driver/mysql" // Blank import because we only need the side effects of the mysql driver
	"github.com/jmoiron/sqlx"
)

// NewDatabasePool is a factory function for creating a database
// connection pool with sqlx from a given DatabaseConfiguration instance.
func NewDatabasePool(configuration *Configuration) (pool *sqlx.DB, err error) {
	pool, err = sqlx.Open("mysql", configuration.URI())
	if err != nil {
		return nil, err
	}
	// TODO: consider tuning here.
	return pool, nil
}
