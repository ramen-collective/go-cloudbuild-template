package database

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestNilViper(t *testing.T) {
	require := require.New(t)
	_, err := NewConfiguration(nil)
	require.EqualError(err, "viper is nil")
}

func TestNewConfiguration(t *testing.T) {
	require := require.New(t)
	os.Setenv("DATABASE_ADDRESS", "1.1.1.1")
	os.Setenv("DATABASE_NAME", "api-template")
	os.Setenv("DATABASE_USERNAME", "jean")
	os.Setenv("DATABASE_PASSWORD", "meskin")
	v := viper.New()
	configuration, err := NewConfiguration(v)
	require.NoError(err)
	require.Equal(
		"jean:meksin@tcp(1.1.1.1:3306)/api-template?parseTime=true",
		configuration.URI(),
	)
}
