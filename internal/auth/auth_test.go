package auth

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	authHeader := make(map[string][]string)
	authHeader["Authorization"] = append(authHeader["Authorization"], "abc123")
	got, _ := GetAPIKey(authHeader)
	want := "abc123"
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q want %q", got, want)
	}
	// Header.Set("Authorization", "Bearer abc123"
}
