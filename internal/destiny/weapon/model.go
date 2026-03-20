package weapon

type WeaponDTO struct {
	InstanceID string `json:"instance_id"`
	Name       string `json:"name"`
	Hash       uint32 `json:"hash"`
	Slot       string `json:"slot"`
	Tier       string `json:"tier"`
	DamageType string `json:"damage_type"`
	IsEquipped bool   `json:"is_equipped"`
	Power      int    `json:"power"`
}
