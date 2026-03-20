package config

import "testing"

func TestValidate_MissingAll(t *testing.T) {
	cfg := &Config{}
	if err := cfg.Validate(); err == nil {
		t.Fatalf("expected validation error")
	}
}

func TestValidate_OK(t *testing.T) {
	cfg := &Config{
		APIKey:      "key",
		ClientID:    "id",
		Secret:      "secret",
		RedirectURL: "https://localhost:4200/",
	}
	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
