package client

import (
	"net/http"
)

const (
	appleStoreRegionHeader string = "X-AppleStore-Region"
	playStoreRegionHeader  string = "X-PlayStore-Region"

	clientRegionHeader string = "X-Client-Region"
	clientCityHeader   string = "X-Client-City"
	clientIPHeader     string = "X-Client-IP"

	unknownRegion string = "UNKNOWN"
	unknownCity   string = "UNKNOWN"
)

func appleStoreRegionForHeaders(headers http.Header) string {
	region := headers.Get(appleStoreRegionHeader)
	region, exist := countryISOMapping[region]
	if !exist {
		return ""
	}
	return region
}

func playStoreRegionForHeaders(headers http.Header) string {
	region := headers.Get(playStoreRegionHeader)
	// TODO: handle any specific ISO conversion
	return region
}

func clientStoreRegionForHeaders(headers http.Header) string {
	appleStoreRegion := appleStoreRegionForHeaders(headers)
	playStoreRegion := playStoreRegionForHeaders(headers)
	var storeRegion string
	if appleStoreRegion != "" {
		storeRegion = appleStoreRegion
	} else if playStoreRegion != "" {
		storeRegion = playStoreRegion
	} else {
		storeRegion = unknownRegion
	}
	return storeRegion
}

func clientRegionForHeaders(headers http.Header) string {
	region := headers.Get(clientRegionHeader)
	if region == "" {
		region = unknownRegion
	}
	return region
}

func clientCityForHeaders(headers http.Header) string {
	city := headers.Get(clientCityHeader)
	if city == "" {
		city = unknownCity
	}
	return city
}

func clientIPAddressForHeaders(headers http.Header) string {
	ip := headers.Get(clientIPHeader)
	// TODO: X-Forwarded-For fallback
	// TODO: other standard address header fallback
	return ip
}

// ClientMiddleware is a custom type for our middleware
type ClientMiddleware func(next http.Handler) http.Handler

// NewClientMiddleware create a ClientMiddleware
func NewClientMiddleware() ClientMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			ctx := request.Context()
			ctx = WithRegion(ctx, clientRegionForHeaders(request.Header))
			ctx = WithStoreRegion(ctx, clientStoreRegionForHeaders(request.Header))
			ctx = WithCity(ctx, clientCityForHeaders(request.Header))
			ctx = WithIP(ctx, clientIPAddressForHeaders(request.Header))
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}
