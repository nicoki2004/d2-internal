package testutil

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/nicoki2004/d2-internal/internal/destiny"
)

func WriteProfileCache(t *testing.T, profile destiny.ProfileResponse) string {
	t.Helper()

	f, err := os.CreateTemp("", "profile_cache_*.json")
	if err != nil {
		t.Fatalf("create temp cache: %v", err)
	}
	defer f.Close()

	data, err := json.Marshal(profile)
	if err != nil {
		t.Fatalf("marshal profile: %v", err)
	}

	if _, err := f.Write(data); err != nil {
		t.Fatalf("write cache: %v", err)
	}

	return f.Name()
}
