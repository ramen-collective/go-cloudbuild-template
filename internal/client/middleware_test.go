package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockHandler struct {
	http.Handler
	assert *require.Assertions
}

func (handler MockHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	handler.assert.Equal("BE", ctx.Value(clientStoreRegionContextKey))
	handler.assert.Equal("FR", ctx.Value(clientRegionContextKey))
	handler.assert.Equal("Paris", ctx.Value(clientCityContextKey))
	handler.assert.Equal("192.168.0.1", ctx.Value(clientIPContextKey))
}

func TestClientMiddleware(t *testing.T) {
	assert := require.New(t)
	request, err := http.NewRequest("GET", "/down", nil)
	assert.NoError(err)
	request.Header.Add("X-AppleStore-Region", "BEL")
	request.Header.Add("X-Client-Region", "FR")
	request.Header.Add("X-Client-City", "Paris")
	request.Header.Add("X-Client-IP", "192.168.0.1")
	handler := &MockHandler{assert: assert}
	factory := NewClientMiddleware()
	middleware := factory(handler)
	middleware.ServeHTTP(nil, request)
}
