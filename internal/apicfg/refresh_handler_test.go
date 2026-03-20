package apicfg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/nicoki2004/d2-internal/internal/config"
	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/destiny/testutil"
	"github.com/nicoki2004/d2-internal/internal/models"
)

func TestHandlerGetProfile_GoldenJSON(t *testing.T) {
	charID := "char-1"
	itemHash := uint32(123456)
	instanceID := "inst-1"
	lastPlayed := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	testutil.InitRegistryWithSingleItem(itemHash, "Test Weapon", "/icon.png")
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

	cfg := &APIConfig{
		Client: destiny.Client{
			Cfg:   &config.Config{},
			Token: &models.Token{},
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/refresh", nil)
	w := httptest.NewRecorder()
	cfg.HandlerGetProfile(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got any
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	gotBytes, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		t.Fatalf("marshal got: %v", err)
	}

	const want = `[
  {
    "CharacterID": "char-1",
    "ClassName": "Hunter",
    "ClassType": 1,
    "Current": true,
    "Emblem": {
      "EmblemBackground": "/bg.png",
      "EmblemHash": "",
      "EmblemPath": "/emblem.png"
    },
    "GenderName": "Male",
    "GenderType": "0",
    "IsVault": false,
    "LastPlayed": "2025-01-01T00:00:00Z",
    "Level": 2000,
    "Name": "",
    "PowerLevel": 2000,
    "Stats": [
      {
        "Hash": "2996146975",
        "Name": "Mobility",
        "Value": 30
      }
    ],
    "TitleInfo": {
      "GildedNum": 0,
      "IsCompleted": false,
      "IsGildedForCurrentSeason": false,
      "Title": "None"
    },
    "equiped": [
      {
        "ItemCategoryHashes": [
          2,
          5
        ],
        "Slot": "Kinetic",
        "hash": 123456,
        "icon": "https://www.bungie.net/icon.png",
        "instanceId": "inst-1",
        "isEquipped": true,
        "isExotic": false,
        "itemType": "Weapon",
        "name": "Test Weapon",
        "power": 2000,
        "tierName": "Legendary",
        "weaponStats": {
          "ammoType": 1,
          "damageType": 3,
          "typeName": "Auto Rifle"
        }
      }
    ]
  },
  {
    "CharacterID": "0",
    "ClassName": "Vault",
    "ClassType": 3,
    "Current": false,
    "Emblem": {
      "EmblemBackground": "/common/destiny2_content/icons/icon_vault_background.png",
      "EmblemHash": "",
      "EmblemPath": "/common/destiny2_content/icons/icon_vault.png"
    },
    "GenderName": "",
    "GenderType": "",
    "IsVault": true,
    "LastPlayed": "0001-01-01T00:00:00Z",
    "Level": 0,
    "Name": "Vault",
    "PowerLevel": 0,
    "Stats": [],
    "TitleInfo": {
      "GildedNum": 0,
      "IsCompleted": false,
      "IsGildedForCurrentSeason": false,
      "Title": "Account Wide"
    },
    "equiped": null
  }
]`

	if strings.TrimSpace(string(gotBytes)) != strings.TrimSpace(want) {
		t.Fatalf("golden mismatch\n--- got ---\n%s\n--- want ---\n%s", string(gotBytes), want)
	}
}
