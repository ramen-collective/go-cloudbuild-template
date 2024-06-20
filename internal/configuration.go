package internal

import (
	"os"
	"strings"

	"github.com/ramen-collective/go-cloudbuild-template/internal/database"
	"github.com/ramen-collective/go-cloudbuild-template/internal/server"
	"github.com/spf13/viper"
)

// Configuration is the root level configuration holder.
type Configuration struct {
	Database *database.Configuration
	Server   *server.Configuration
}

// A static factory function that create a viper.Viper
// instance using the YAML file denoted by the given
// path and environment variable precedence.
func newViper(path string) (v *viper.Viper, err error) {
	var instance *viper.Viper = viper.New()
	_, fileError := os.Stat(path + "/api-template.yaml")
	replacer := strings.NewReplacer(".", "_")
	instance.SetEnvKeyReplacer(replacer)
	instance.AutomaticEnv()
	if !os.IsNotExist(fileError) {
		instance.AddConfigPath(path)
		instance.SetConfigName("api-template")
		instance.SetConfigType("yaml")
		err = instance.ReadInConfig()
	}
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// NewConfiguration is a factory function for creating a root level
// application Configuration. Evaluating from both YAML file and
// environment variables.
func NewConfiguration(path string) (configuration *Configuration, err error) {
	v, err := newViper(path)
	if err != nil {
		return nil, err
	}
	databaseConfiguration, err := database.NewConfiguration(v)
	if err != nil {
		return nil, err
	}
	serverConfiguration, err := server.NewConfiguration(v)
	if err != nil {
		return nil, err
	}
	return &Configuration{
		Database: databaseConfiguration,
		Server:   serverConfiguration,
	}, nil
}
