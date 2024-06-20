package gqlutil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewIDFromValidString(t *testing.T) {
	require := require.New(t)

	ID := NewIDFromString("1234")

	require.Equal(ID.Int(), uint64(1234))
	require.Equal(ID.String(), "1234")
}

func TestNewIDFromInvalidString(t *testing.T) {
	require := require.New(t)

	ID := NewIDFromString("azerty")

	require.Equal(ID.Int(), uint64(0))
}
