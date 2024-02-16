package auth

import (
	"net/http"
	"reflect"
	// "strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey abc123")
	got, err := GetAPIKey(header)
	want := "abc123"
	if err != nil {
		t.Log(err)
		// t.Log(header.Get("Authorization"))
		// h := strings.Split(header.Get("Authorization"), " ")
		// t.Logf("0 => %q", h[0])
		// t.Logf("1 => %q", h[1])
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q want %q", got, want)
	}
	// Header.Set("Authorization", "Bearer abc123"
}
