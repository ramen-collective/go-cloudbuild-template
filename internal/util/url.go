package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const (
	signatureQueryKey  string = "signature"
	expirationQueryKey string = "expires"
)

// URLGeneratorInterface should be implemented by url generate service
type URLGeneratorInterface interface {
	SignedURL(URL string, ttl time.Duration) string
	ValidateURISignature(URI string) bool
}

// URLGenerator handle URL generation service
type URLGenerator struct {
	baseURL    string
	signingKey string
}

// NewURLGenerator instanciate a new URLGenerator
func NewURLGenerator(baseURL string, signingKey string) *URLGenerator {
	return &URLGenerator{
		baseURL:    baseURL,
		signingKey: signingKey,
	}
}

// SignedURL returns a signed url for a given path.
func (u URLGenerator) SignedURL(path string, ttl time.Duration) string {
	url := fmt.Sprintf("%s/%s?%s=%d", u.baseURL, path, expirationQueryKey, time.Now().Add(ttl).Unix())
	signature := u.GenerateSignatureForURL(url)
	return fmt.Sprintf("%s&%s=%s", url, signatureQueryKey, signature)
}

// ValidateURISignature validates a URL signature.
func (u URLGenerator) ValidateURISignature(URI string) bool {
	URL := fmt.Sprintf("%s%s", u.baseURL, URI)
	parsedURL, err := url.Parse(URL)
	if err != nil {
		return false
	}
	query := parsedURL.Query()
	if !query.Has(signatureQueryKey) || !query.Has(expirationQueryKey) {
		return false
	}
	signature := query.Get(signatureQueryKey)
	query.Del(signatureQueryKey)
	parsedURL.RawQuery = query.Encode()
	if signature != u.GenerateSignatureForURL(parsedURL.String()) {
		return false
	}
	expirationTime, _ := strconv.ParseInt(query.Get(expirationQueryKey), 10, 64)
	if expirationTime < time.Now().Unix() {
		return false
	}
	return true
}

// GenerateSignatureForURL generates a hmac sha256 signature
// for a given URL.
func (u URLGenerator) GenerateSignatureForURL(URL string) string {
	hmac := hmac.New(sha256.New, []byte(u.signingKey))
	hmac.Write([]byte(URL))
	return base64.URLEncoding.EncodeToString(hmac.Sum(nil))
}
