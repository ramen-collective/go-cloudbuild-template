package client

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestRegionForContext(t *testing.T) {
	require := require.New(t)
	ctx := context.WithValue(context.Background(), clientRegionContextKey, "FR")
	require.Equal("FR", RegionForContext(ctx))
}

func TestStoreRegionForContext(t *testing.T) {
	require := require.New(t)
	ctx := context.WithValue(context.Background(), clientStoreRegionContextKey, "FR")
	require.Equal("FR", StoreRegionForContext(ctx))
}

func TestNoStoreRegionForContext(t *testing.T) {
	require := require.New(t)
	require.Equal("", StoreRegionForContext(context.Background()))
}

func TestAnyRegionForContext(t *testing.T) {
	require := require.New(t)
	ctx := context.WithValue(context.Background(), clientStoreRegionContextKey, "FR")
	ctx = context.WithValue(ctx, clientRegionContextKey, "BE")
	require.Equal("FR", AnyRegionForContext(ctx))
	ctx = context.WithValue(context.Background(), clientRegionContextKey, "BE")
	require.Equal("BE", AnyRegionForContext(ctx))
	require.Equal("UNKNOWN", AnyRegionForContext(context.Background()))
}

func TestCityForContext(t *testing.T) {
	require := require.New(t)
	ctx := context.WithValue(context.Background(), clientCityContextKey, "Paris")
	require.Equal("Paris", CityForContext(ctx))
}

func TestIPForContext(t *testing.T) {
	require := require.New(t)
	ctx := context.WithValue(context.Background(), clientIPContextKey, "192.168.0.1")
	require.Equal("192.168.0.1", IPForContext(ctx))
}
