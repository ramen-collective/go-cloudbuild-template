package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

const (
	_deleteQuery = "DELETE FROM %s WHERE %s;"
)

// Container contains an instance of all model repositories
type Container struct {
	User UserRepositoryInterface
}

// NewContainer returns a *RepositoryContainer
func NewContainer(db *sqlx.DB) *Container {
	return &Container{
		User: NewUserRepository(db),
	}
}

// Delete execute and returns rows affected by a built
// DELETE query using provided WHERE clause.
func Delete(
	database *sqlx.DB,
	table string,
	where string,
	arguments ...interface{},
) (bool, error) {
	query := fmt.Sprintf(_deleteQuery, table, where)
	cursor, err := database.Exec(query, arguments...)
	if err != nil {
		return false, err
	}
	count, err := cursor.RowsAffected()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// DeleteByID remove a row from the given (database,table pair
// denoted by the given identifier.
func DeleteByID(database *sqlx.DB, table string, id uint64) (bool, error) {
	return Delete(database, table, "id = ?", id)
}

// DeleteByUUID remove a row from the given (database,table pair
// denoted by the given identifier.
func DeleteByUUID(database *sqlx.DB, table string, uuid string) (bool, error) {
	return Delete(database, table, "uuid = ?", uuid)
}

// Update execute and returns rows affected by a built
// UPDATE query using provided WHERE clause.
func Update(
	database *sqlx.DB,
	table string,
	where string,
	fields map[string]interface{},
	whereArguments ...interface{},
) (bool, error) {
	var builder strings.Builder
	arguments := make([]interface{}, 0, len(fields)+1)
	builder.WriteString("UPDATE ")
	builder.WriteString(table)
	builder.WriteString(" SET ")
	i := 0
	for field, value := range fields {
		builder.WriteString(field)
		builder.WriteString(" = ?")
		if i < len(fields)-1 {
			builder.WriteString(",")
		}
		i++
		arguments = append(arguments, value)
	}
	arguments = append(arguments, whereArguments...)
	builder.WriteString(" WHERE ")
	builder.WriteString(where)
	builder.WriteString(";")
	cursor, err := database.Exec(builder.String(), arguments...)
	if err != nil {
		return false, err
	}
	count, err := cursor.RowsAffected()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// UpdateByID updates a row from the given (database,table pair
// denoted by the given identifier.
func UpdateByID(
	database *sqlx.DB,
	table string,
	id uint64,
	fields map[string]interface{},
) (bool, error) {
	return Update(database, table, "id = ?", fields, id)
}

// UpdateByUUID updates a row from the given (database,table pair
// denoted by the given identifier).
func UpdateByUUID(
	database *sqlx.DB,
	table string,
	uuid string,
	fields map[string]interface{},
) (bool, error) {
	return Update(database, table, "uuid = ?", fields, uuid)
}

// UpdateByKey updates a row from the given (database,table pair
// denoted by the given identifier).
func UpdateByKey(
	database *sqlx.DB,
	table string,
	keyValue string,
	keyName string,
	fields map[string]interface{},
) (bool, error) {
	return Update(database, table, keyName+" = ?", fields, keyValue)
}

// KeyCountByID makes a map[uint64]int from a *sqlx.Rows
func KeyCountByID(rows *sqlx.Rows, IDColumn string, CountColumn string) (map[uint64]int, error) {
	counts := map[uint64]int{}
	for rows.Next() {
		cols := make(map[string]interface{})
		err := rows.MapScan(cols)
		if err != nil {
			return nil, err
		}
		var dbID uint64
		var dbCount int
		if value, ok := cols[IDColumn].(int64); ok {
			dbID = uint64(value)
		} else {
			return counts, fmt.Errorf("ID Column is not an int64")
		}
		if value, ok := cols[CountColumn].(int64); ok {
			dbCount = int(value)
		} else {
			return counts, fmt.Errorf("Count Column is not an int64")
		}
		counts[dbID] = dbCount
	}
	return counts, nil
}
