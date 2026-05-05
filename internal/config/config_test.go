package config

import (
	"strings"
	"testing"
)

func TestEnvPrefix(t *testing.T) {
	if got, want := EnvPrefix("math.CT"), "MIXI2_MATH_CT"; got != want {
		t.Fatalf("EnvPrefix() = %q, want %q", got, want)
	}
}

func TestLoadCredentials(t *testing.T) {
	t.Setenv("MIXI2_MATH_CT_CLIENT_ID", "client-id")
	t.Setenv("MIXI2_MATH_CT_CLIENT_SECRET", "client-secret")
	t.Setenv("MIXI2_MATH_CT_TOKEN_URL", "https://token.example")
	t.Setenv("MIXI2_MATH_CT_API_ADDRESS", "api.example:443")

	creds, err := LoadCredentials("math.CT")
	if err != nil {
		t.Fatalf("LoadCredentials() error = %v", err)
	}
	if creds.ClientID != "client-id" || creds.ClientSecret != "client-secret" || creds.TokenURL == "" || creds.APIAddress == "" {
		t.Fatalf("unexpected credentials: %+v", creds)
	}
}

func TestLoadCredentialsMissing(t *testing.T) {
	_, err := LoadCredentials("math.AG")
	if err == nil {
		t.Fatal("LoadCredentials() error = nil, want missing env error")
	}
	if !strings.Contains(err.Error(), "MIXI2_MATH_AG_CLIENT_ID") {
		t.Fatalf("error %q does not include missing env name", err)
	}
}
