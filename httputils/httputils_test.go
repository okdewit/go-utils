package httputils

import (
	"net/url"
	"testing"
)

func TestAddParams(t *testing.T) {
	uri, _ := url.Parse("https://example.com")
	AddParams(uri, map[string]string{"key": "value"})
	if uri.String() != "https://example.com?key=value" {
		t.Errorf("AddParams result is %s; want https://example.com?key=value", uri.String())
	}
}
