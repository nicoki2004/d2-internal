package store

import (
	"testing"
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny/testutil"
)

func TestGetStores_MinimalProfile(t *testing.T) {
	charID := "char-1"
	itemHash := uint32(123456)
	instanceID := "inst-1"

	testutil.InitRegistryWithSingleItem(itemHash, "Test Weapon", "/icon.png")

	lastPlayed := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	profile := testutil.MinimalProfileWithEquipped(charID, itemHash, instanceID, lastPlayed)

	stores := GetStores(profile)
	if len(stores) != 2 {
		t.Fatalf("expected 2 stores (character + vault), got %d", len(stores))
	}

	store := stores[0]
	if store.CharacterID != charID {
		t.Fatalf("expected characterID %s, got %s", charID, store.CharacterID)
	}
	if store.ClassName != "Hunter" {
		t.Fatalf("expected ClassName Hunter, got %s", store.ClassName)
	}
	if !store.Current {
		t.Fatalf("expected Current=true for last played character")
	}
	if len(store.Equiped) != 1 {
		t.Fatalf("expected 1 equipped item, got %d", len(store.Equiped))
	}

	itm := store.Equiped[0]
	if itm.Name != "Test Weapon" {
		t.Fatalf("expected item name Test Weapon, got %s", itm.Name)
	}
	if itm.Slot != "Kinetic" {
		t.Fatalf("expected slot Kinetic, got %s", itm.Slot)
	}
	if itm.WeaponStats == nil || itm.WeaponStats.TypeName != "Auto Rifle" {
		t.Fatalf("expected weapon type Auto Rifle")
	}
}
