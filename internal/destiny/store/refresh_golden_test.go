package store

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny/testutil"
)

func TestGetStores_GoldenJSON(t *testing.T) {
	charID := "char-1"
	itemHash := uint32(123456)
	instanceID := "inst-1"
	lastPlayed := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	testutil.InitRegistryWithSingleItem(itemHash, "Test Weapon", "/icon.png")
	profile := testutil.MinimalProfileWithEquipped(charID, itemHash, instanceID, lastPlayed)

	stores := GetStores(profile)
	got, err := json.MarshalIndent(stores, "", "  ")
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	const want = `[
  {
    "CharacterID": "char-1",
    "Name": "",
    "IsVault": false,
    "ClassType": 1,
    "ClassName": "Hunter",
    "GenderType": "0",
    "GenderName": "Male",
    "Emblem": {
      "EmblemPath": "/emblem.png",
      "EmblemBackground": "/bg.png",
      "EmblemHash": ""
    },
    "Current": true,
    "LastPlayed": "2025-01-01T00:00:00Z",
    "Level": 2000,
    "PowerLevel": 2000,
    "TitleInfo": {
      "Title": "None",
      "IsCompleted": false,
      "GildedNum": 0,
      "IsGildedForCurrentSeason": false
    },
    "Stats": [
      {
        "Hash": "2996146975",
        "Name": "Mobility",
        "Value": 30
      }
    ],
    "equiped": [
      {
        "instanceId": "inst-1",
        "hash": 123456,
        "name": "Test Weapon",
        "icon": "https://www.bungie.net/icon.png",
        "itemType": "Weapon",
        "tierName": "Legendary",
        "isExotic": false,
        "power": 2000,
        "isEquipped": true,
        "ItemCategoryHashes": [
          2,
          5
        ],
        "Slot": "Kinetic",
        "weaponStats": {
          "typeName": "Auto Rifle",
          "damageType": 3,
          "ammoType": 1
        }
      }
    ]
  },
  {
    "CharacterID": "0",
    "Name": "Vault",
    "IsVault": true,
    "ClassType": 3,
    "ClassName": "Vault",
    "GenderType": "",
    "GenderName": "",
    "Emblem": {
      "EmblemPath": "/common/destiny2_content/icons/icon_vault.png",
      "EmblemBackground": "/common/destiny2_content/icons/icon_vault_background.png",
      "EmblemHash": ""
    },
    "Current": false,
    "LastPlayed": "0001-01-01T00:00:00Z",
    "Level": 0,
    "PowerLevel": 0,
    "TitleInfo": {
      "Title": "Account Wide",
      "IsCompleted": false,
      "GildedNum": 0,
      "IsGildedForCurrentSeason": false
    },
    "Stats": [],
    "equiped": null
  }
]`

	if strings.TrimSpace(string(got)) != strings.TrimSpace(want) {
		t.Fatalf("golden mismatch\n--- got ---\n%s\n--- want ---\n%s", string(got), want)
	}
}
