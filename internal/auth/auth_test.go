package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_empty_header(t *testing.T) {
	header := http.Header{}
	got1, got2 := GetAPIKey(header)
	want1, want2 := "", ErrNoAuthHeaderIncluded
	if want1 != got1 {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}

	if want2 != got2 {
		t.Fatalf("expected: %v, got: %v", want2, got2)
	}
}

func TestGetAPIKey_malformed(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"ApiKeymy-secret-key"},
	}
	got1, got2 := GetAPIKey(header)
	want1, want2msg := "", "malformed authorization header"
	if want1 != got1 {
		t.Fatalf("expected: %v, got: %v", want1, got1)
	}

	if want2msg != got2.Error() {
		t.Fatalf("expected: %v, got: %v", want2msg, got2)
	}
}
