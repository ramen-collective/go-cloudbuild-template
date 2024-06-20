package database

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// Configuration provide information relative to
// MySQL database connection.
type Configuration struct {
	Address  string
	Name     string
	Password string
	Port     int
	Username string
}

// URI is a factory method for generating a database URI.
func (configuration *Configuration) URI() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configuration.Username,
		configuration.Password,
		configuration.Address,
		configuration.Port,
		configuration.Name)
}

// NewConfiguration is a factory function for creating
// a DatabaseConfiguration instance using a viper sub tree.
func NewConfiguration(v *viper.Viper) (configuration *Configuration, err error) {
	if v == nil {
		return nil, errors.New("viper is nil")
	}
	v.SetDefault("database.port", 3306)
	configuration = &Configuration{
		Address:  v.GetString("database.address"),
		Name:     v.GetString("database.name"),
		Password: v.GetString("database.password"),
		Port:     v.GetInt("database.port"),
		Username: v.GetString("database.username"),
	}
	// NOTE: add data validation here if needed.
	return configuration, nil
}
