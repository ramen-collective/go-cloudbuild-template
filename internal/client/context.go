package client

import "context"

// ContextKey represents a context key for authentication
// context values
type ContextKey string

const (
	clientRegionContextKey      ContextKey = "client.region"
	clientStoreRegionContextKey ContextKey = "client.store.region"
	clientCityContextKey        ContextKey = "client.city"
	clientIPContextKey          ContextKey = "client.ip"
)

// WithRegion build a context with a given region
func WithRegion(ctx context.Context, region string) context.Context {
	return context.WithValue(ctx, clientRegionContextKey, region)
}

// WithStoreRegion build a context with a given store region
func WithStoreRegion(ctx context.Context, region string) context.Context {
	return context.WithValue(ctx, clientStoreRegionContextKey, region)
}

// WithCity build a context with a given city
func WithCity(ctx context.Context, city string) context.Context {
	return context.WithValue(ctx, clientCityContextKey, city)
}

// WithIP build a context with a given ip
func WithIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, clientIPContextKey, ip)
}

// RegionForContext extract region with priority over
func RegionForContext(ctx context.Context) string {
	region := ctx.Value(clientRegionContextKey)
	if region == nil {
		return ""
	}
	return region.(string)
}

// StoreRegionForContext extract store region
func StoreRegionForContext(ctx context.Context) string {
	region := ctx.Value(clientStoreRegionContextKey)
	if region == nil {
		return ""
	}
	return region.(string)
}

// AnyRegionForContext extract store region or client if not present
func AnyRegionForContext(ctx context.Context) string {
	region := StoreRegionForContext(ctx)
	if region == unknownRegion || region == "" {
		region = RegionForContext(ctx)
	}
	if region == "" {
		region = unknownRegion
	}
	return region
}

// CityForContext extract city
func CityForContext(ctx context.Context) string {
	city := ctx.Value(clientCityContextKey)
	if city == nil {
		return ""
	}
	return city.(string)
}

// IPForContext extract IP
func IPForContext(ctx context.Context) string {
	ip := ctx.Value(clientIPContextKey)
	if ip == nil {
		return ""
	}
	return ip.(string)
}
