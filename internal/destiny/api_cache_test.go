package destiny_test

import (
	"os"
	"testing"
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/destiny/testutil"
	"github.com/nicoki2004/d2-internal/internal/models"
)

func TestGetProfile_UsesCache(t *testing.T) {
	charID := "char-1"
	itemHash := uint32(123456)
	instanceID := "inst-1"
	lastPlayed := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	profile := testutil.MinimalProfileWithEquipped(charID, itemHash, instanceID, lastPlayed)
	cachePath := testutil.WriteProfileCache(t, profile)

	oldCache := os.Getenv("CACHE_FILE_NAME")
	if err := os.Setenv("CACHE_FILE_NAME", cachePath); err != nil {
		t.Fatalf("set CACHE_FILE_NAME: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Setenv("CACHE_FILE_NAME", oldCache)
		_ = os.Remove(cachePath)
	})

	client := &destiny.Client{Token: &models.Token{}}
	got, err := destiny.GetProfile(client)
	if err != nil {
		t.Fatalf("GetProfile error: %v", err)
	}

	if got.Response.Characters.Data[charID].CharacterID != charID {
		t.Fatalf("expected character in cached profile")
	}
}
