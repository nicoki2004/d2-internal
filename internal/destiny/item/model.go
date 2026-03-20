package item

type ItemType string

const (
	ItemWeapon ItemType = "Weapon"
	ItemArmor  ItemType = "Armor"
	ItemOther  ItemType = "Other"
)

type ItemDTO struct {
	InstanceID         string   `json:"instanceId"`
	Hash               uint32   `json:"hash"`
	Name               string   `json:"name"`
	Icon               string   `json:"icon"`
	ItemType           ItemType `json:"itemType"`
	TierName           string   `json:"tierName"` // "Exotic", "Legendary"
	IsExotic           bool     `json:"isExotic"`
	Power              int      `json:"power"`
	IsEquipped         bool     `json:"isEquipped"`
	ItemCategoryHashes []uint32
	Slot               string

	// Specs condicionales
	WeaponStats *WeaponSpec `json:"weaponStats,omitempty"`
	ArmorStats  *ArmorSpec  `json:"armorStats,omitempty"`
}

type WeaponSpec struct {
	TypeName   string `json:"typeName"`   // "Hand Cannon"
	DamageType int    `json:"damageType"` // 1: Kinetic, 2: Arc, etc.
	AmmoType   int    `json:"ammoType"`   // 1: Primary, 2: Special...
}

type ArmorSpec struct {
	SlotName string `json:"slotName"` // "Helmet", "Chest"
}
