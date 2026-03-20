package weapon

import (
	"context"

	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/repository"
)

func GetWeaponsWithPerks(ctx context.Context, repo repository.DestinyRepository) ([]WeaponDTO, error) {
	data, err := repo.GetAllWeaponsWithPerks(ctx)
	if err != nil {
		logger.GetLogger().Debug("Error getting All Characters: %v", err)
	}

	response := map[string]WeaponDTO

	for _, weapon := range data{
		w, ok := response[weapon.InstanceID]
		if !ok{
			w := &WeaponDTO{
				InstanceID: weapon.InstanceID,
				Hash: uint32(weapon.Hash),
				Name: weapon.Name,
				Slot: weapon.Slot,
				DamageType: weapon.Type,
				IsEquipped: weapon,
				Power: int(weapon.Power),
				
			}
				
		}
	}

}
	InstanceID string `json:"instance_id"`
	Name       string `json:"name"`
	Hash       uint32 `json:"hash"`
	Slot       string `json:"slot"`
	Tier       string `json:"tier"`
	DamageType string `json:"damage_type"`
	IsEquipped bool   `json:"is_equipped"`
	Power      int    `json:"power"`


	InstanceID  string         `json:"instance_id"`
	Hash        int64          `json:"hash"`
	Name        string         `json:"name"`
	Type        string         `json:"type"`
	Power       int64          `json:"power"`
	Kills       int64          `json:"kills"`
	Level       int64          `json:"level"`
	Location    string         `json:"location"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	Tier        sql.NullString `json:"tier"`
	IconUrl     sql.NullString `json:"icon_url"`
	Slot        sql.NullString `json:"slot"`
	DamageType  sql.NullString `json:"damage_type"`
	AmmoType    sql.NullInt64  `json:"ammo_type"`
	CharacterID sql.NullString `json:"character_id"`
	PerksList   string         `json:"perks_list"`
