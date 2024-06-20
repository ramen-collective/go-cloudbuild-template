package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Configuration provide options for HTTP
// server exposure.
type Configuration struct {
	Hostname        string
	Playground      bool
	Port            int
	HealthPort      string
	ShutdownTimeout time.Duration
}

// NewConfiguration is a factory function for creating a
// Configuration instance using a viper sub tree.
func NewConfiguration(v *viper.Viper) (configuration *Configuration, err error) {
	if v == nil {
		return nil, errors.New("viper is nil")
	}
	v.BindEnv("server.port", "PORT")
	v.BindEnv("server.service_url", "SERVICE_URL")
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 80)
	v.SetDefault("server.playground", true)
	configuration = &Configuration{
		Hostname:   v.GetString("server.host"),
		Playground: v.GetBool("server.playground"),
		Port:       v.GetInt("server.port"),
		// TODO: set as configuration value.
		ShutdownTimeout: time.Second * 20,
	}
	// NOTE: add data validation here if needed.
	return configuration, nil
}

// Addr is a factory method for generating a server binding address.
func (configuration *Configuration) Addr() string {
	return fmt.Sprintf(
		"%s:%d",
		configuration.Hostname,
		configuration.Port)
}
