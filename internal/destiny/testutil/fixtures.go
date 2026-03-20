package testutil

import (
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/registry"
)

func MinimalWeaponManifest(name, icon string) destiny.ManifestItem {
	return destiny.ManifestItem{
		DisplayProperties: struct {
			Name string `json:"name"`
			Icon string `json:"icon"`
		}{
			Name: name,
			Icon: icon,
		},
		ItemCategoryHashes:  []uint32{2, 5}, // Kinetic + Auto Rifle
		ItemTypeDisplayName: "Auto Rifle",
		ItemType:            3,
		Inventory: struct {
			TierTypeName   string `json:"tierTypeName"`
			TierType       int    `json:"tierType"`
			BucketTypeHash uint32 `json:"bucketTypeHash"`
		}{
			TierTypeName: "Legendary",
			TierType:     5,
		},
		EquippingBlock: struct {
			AmmoType int32 `json:"ammoType"`
		}{
			AmmoType: 1,
		},
		DefaultDamageTypeHash: 0,
	}
}

func InitRegistryWithSingleItem(itemHash uint32, name, icon string) {
	registry.InitItemRegistry(map[uint32]destiny.ManifestItem{
		itemHash: MinimalWeaponManifest(name, icon),
	})
}

func MinimalProfileWithEquipped(
	charID string,
	itemHash uint32,
	instanceID string,
	lastPlayed time.Time,
) destiny.ProfileResponse {
	builder := NewProfileBuilder().
		WithCharacter(charID, 1, 2000, lastPlayed, 0).
		WithStat(charID, 2996146975, 30).
		WithEquippedItem(charID, itemHash, instanceID).
		WithInstance(instanceID, 3, 2000)

	return builder.Build()
}
