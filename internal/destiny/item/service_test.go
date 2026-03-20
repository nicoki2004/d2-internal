package item

import (
	"testing"

	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/destiny/testutil"
	"github.com/nicoki2004/d2-internal/internal/registry"
)

func TestMapEquippedItems_ManifestMiss(t *testing.T) {
	registry.InitItemRegistry(map[uint32]destiny.ManifestItem{})

	var profile destiny.ProfileResponse
	profile.Response.CharacterEquipment.Data = map[string]destiny.CharacterEquipmentData{
		"char-1": {
			Items: []destiny.Item{{
				ItemHash:       123,
				ItemInstanceID: "inst-1",
			}},
		},
	}

	items := MapEquippedItems("char-1", profile)
	if len(items) != 0 {
		t.Fatalf("expected 0 items when manifest is missing, got %d", len(items))
	}
}

func TestMapEquippedItems_Minimal(t *testing.T) {
	itemHash := uint32(123456)
	instanceID := "inst-1"

	testutil.InitRegistryWithSingleItem(itemHash, "Test Weapon", "/icon.png")

	var profile destiny.ProfileResponse
	profile.Response.CharacterEquipment.Data = map[string]destiny.CharacterEquipmentData{
		"char-1": {
			Items: []destiny.Item{{
				ItemHash:       itemHash,
				ItemInstanceID: instanceID,
			}},
		},
	}

	profile.Response.ItemComponents.Instances.Data = map[string]destiny.ItemInstanceData{
		instanceID: {
			DamageType: 3,
			PrimaryStat: struct {
				StatHash uint32 `json:"statHash"`
				Value    int    `json:"value"`
			}{
				Value: 2000,
			},
		},
	}

	items := MapEquippedItems("char-1", profile)
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	itm := items[0]
	if itm.ItemType != ItemWeapon {
		t.Fatalf("expected ItemWeapon, got %s", itm.ItemType)
	}
	if itm.Slot != "Kinetic" {
		t.Fatalf("expected slot Kinetic, got %s", itm.Slot)
	}
	if itm.WeaponStats == nil || itm.WeaponStats.TypeName != "Auto Rifle" {
		t.Fatalf("expected weapon type Auto Rifle")
	}
}
