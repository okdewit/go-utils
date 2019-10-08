package httputils

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestAddParams(t *testing.T) {
	uri, _ := url.Parse("https://example.com")
	AddParams(uri, map[string]string{"key": "value"})
	assert.Equal(t, "https://example.com?key=value", uri.String())
}
