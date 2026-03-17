package destiny

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nicoki2004/d2-internal/internal/database"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/repository"
)

func SyncInventory(
	ctx context.Context,
	repo repository.DestinyRepository,
	profile *ProfileResponse,
	manifest map[string]ManifestItem,
) error {
	fmt.Println("🚀 Starting database synchronization...")

	for _, charData := range profile.Response.Characters.Data {
		err := saveCharacter(ctx, repo, charData)
		if err != nil {
			logger.GetLogger().Error("Error saving Characters: %v", err)
		}
	}

	// 1. Iterate Characters: Equipped (205)
	for charID, equipment := range profile.Response.CharacterEquipment.Data {
		location := fmt.Sprintf("Equipped:%s", charID)
		for _, item := range equipment.Items {
			if err := SaveItem(ctx, repo, item, profile, manifest, location, charID); err != nil {
				return fmt.Errorf("error saving equipped item at %s: %w", location, err)
			}
		}
	}

	// 2. Iterate Characters: Inventory/Backpack (201)
	for charID, inventory := range profile.Response.CharacterInventories.Data {
		location := fmt.Sprintf("Inventory:%s", charID)
		for _, item := range inventory.Items {
			if err := SaveItem(ctx, repo, item, profile, manifest, location, charID); err != nil {
				return fmt.Errorf("error saving inventory item at %s: %w", location, err)
			}
		}
	}

	// 3. Iterate Vault (102)
	for _, item := range profile.Response.ProfileInventory.Data.Items {
		if err := SaveItem(ctx, repo, item, profile, manifest, "Vault", "Vault"); err != nil {
			return fmt.Errorf("error saving vault item: %w", err)
		}
	}

	fmt.Println("✅ Synchronization completed.")
	return nil
}

func SaveItem(
	ctx context.Context,
	repo repository.DestinyRepository,
	item Item,
	profile *ProfileResponse,
	manifest map[string]ManifestItem,
	location string,
	charID string,
) error {
	// Validar que existe en manifest y es un arma
	hashStr := fmt.Sprintf("%d", item.ItemHash)
	def, ok := manifest[hashStr]

	if !ok || def.ItemType != 3 {
		return nil // Not a weapon, not an error
	}

	extractor := NewWeaponExtractor()
	validator := NewPerkValidator()

	// 1. Save weapon base data
	metadata := extractor.ExtractMetadata(item, profile)
	if err := saveWeaponBase(ctx, repo, item, def, metadata, location, charID); err != nil {
		return fmt.Errorf("error saving weapon base data: %w", err)
	}

	// 2. Save stats
	if err := saveWeaponStats(ctx, repo, item, profile, extractor); err != nil {
		return fmt.Errorf("error saving weapon stats: %w", err)
	}

	// 3. Save perks
	if err := saveWeaponPerks(ctx, repo, item, profile, manifest, extractor, validator); err != nil {
		return fmt.Errorf("error saving weapon perks: %w", err)
	}

	return nil
}

// saveWeaponBase inserts or updates the weapon base
func saveWeaponBase(
	ctx context.Context,
	repo repository.DestinyRepository,
	item Item,
	def ManifestItem,
	metadata WeaponMetadata,
	location string,
	charID string,
) error {
	return repo.UpsertWeapon(ctx, database.UpsertWeaponParams{
		InstanceID: item.ItemInstanceID,
		Hash:       int64(item.ItemHash),
		Name:       def.DisplayProperties.Name,
		Type:       def.ItemTypeDisplayName,
		Power:      int64(metadata.Power),
		Kills:      int64(metadata.Kills),
		Level:      int64(metadata.Level),
		Location:   location,
		// NEW
		CharacterID: sql.NullString{String: charID, Valid: true},
		Tier:        sql.NullString{String: def.Inventory.TierTypeName, Valid: true},
		IconUrl:     sql.NullString{String: BungieCDN + def.DisplayProperties.Icon, Valid: true},
		Slot:        sql.NullString{String: GetSlotName(def.Inventory.BucketTypeHash), Valid: true},
		DamageType:  sql.NullString{String: GetDamageName(uint32(def.DefaultDamageTypeHash)), Valid: true},
		AmmoType:    sql.NullInt64{Int64: int64(def.EquippingBlock.AmmoType), Valid: true},
	})
}

// saveWeaponStats saves or updates the weapon stats
func saveWeaponStats(
	ctx context.Context,
	repo repository.DestinyRepository,
	item Item,
	profile *ProfileResponse,
	extractor *WeaponExtractor,
) error {
	if err := repo.ClearWeaponStats(ctx, item.ItemInstanceID); err != nil {
		return err
	}

	stats := extractor.ExtractStats(item, profile)
	for statName, value := range stats {
		if err := repo.InsertWeaponStat(ctx, database.InsertWeaponStatParams{
			InstanceID: item.ItemInstanceID,
			StatName:   statName,
			Value:      int64(value),
		}); err != nil {
			return err
		}
	}

	return nil
}

// saveWeaponPerks saves or updates the weapon perks
func saveWeaponPerks(
	ctx context.Context,
	repo repository.DestinyRepository,
	item Item,
	profile *ProfileResponse,
	manifest map[string]ManifestItem,
	extractor *WeaponExtractor,
	validator *PerkValidator,
) error {
	if err := repo.ClearWeaponPerks(ctx, item.ItemInstanceID); err != nil {
		return err
	}

	socketsData := extractor.ExtractSockets(item, profile)
	if !socketsData.HasSockets {
		return nil
	}

	for i, socket := range socketsData.Sockets.Sockets {
		if socket.PlugHash == 0 {
			continue
		}

		socketIndexStr := fmt.Sprintf("%d", i)
		options, ok := socketsData.Plugs.Plugs[socketIndexStr]

		// Case: Socket with multiple options (perks tree)
		if socketsData.HasReusable && ok && len(options) > 0 {
			if err := savePerkOptions(ctx, repo, item, options, socket, manifest, i, validator); err != nil {
				return err
			}
		} else {
			// Case: Simple socket (no options)
			if err := saveSinglePerk(ctx, repo, item, socket, manifest, i, validator); err != nil {
				return err
			}
		}
	}

	return nil
}

// savePerkOptions saves multiple perk options in a socket
func savePerkOptions(
	ctx context.Context,
	repo repository.DestinyRepository,
	item Item,
	options []PlugEntry,
	socket SocketEntry,
	manifest map[string]ManifestItem,
	socketIndex int,
	validator *PerkValidator,
) error {
	for _, opt := range options {
		perkDef, exists := manifest[fmt.Sprintf("%d", opt.PlugItemHash)]

		if !exists || !validator.IsActualPerk(perkDef.ItemTypeDisplayName) {
			continue
		}

		if err := repo.InsertWeaponPerk(ctx, database.InsertWeaponPerkParams{
			InstanceID:  item.ItemInstanceID,
			PerkHash:    int64(opt.PlugItemHash),
			PerkName:    perkDef.DisplayProperties.Name,
			IsEquipped:  opt.PlugItemHash == socket.PlugHash,
			SocketIndex: int64(socketIndex),
		}); err != nil {
			return err
		}
	}

	return nil
}

// saveSinglePerk saves a single perk with no options
func saveSinglePerk(
	ctx context.Context,
	repo repository.DestinyRepository,
	item Item,
	socket SocketEntry,
	manifest map[string]ManifestItem,
	socketIndex int,
	validator *PerkValidator,
) error {
	perkDef, exists := manifest[fmt.Sprintf("%d", socket.PlugHash)]

	if !exists || !validator.IsActualPerk(perkDef.ItemTypeDisplayName) {
		return nil
	}

	return repo.InsertWeaponPerk(ctx, database.InsertWeaponPerkParams{
		InstanceID:  item.ItemInstanceID,
		PerkHash:    int64(socket.PlugHash),
		PerkName:    perkDef.DisplayProperties.Name,
		IsEquipped:  true,
		SocketIndex: int64(socketIndex),
	})
}

func saveCharacter(ctx context.Context, repo repository.DestinyRepository, charData CharacterData) error {
	statsParam := make([]database.UpsertCharacterStatParams, 0, len(charData.Stats))
	for key, value := range charData.Stats {
		statsParam = append(statsParam, database.UpsertCharacterStatParams{
			CharacterID: charData.CharacterID,
			StatHash:    int64(key),
			Value:       int64(value),
		})
	}

	return repo.SaveCharacterFull(ctx, database.UpsertCharacterParams{
		CharacterID:          charData.CharacterID,
		ClassType:            int64(charData.ClassType),
		LightLevel:           int64(charData.Light),
		EmblemUrl:            sql.NullString{String: charData.EmblemPath, Valid: true},
		LastPlayed:           sql.NullTime{Time: charData.DateLastPlayed, Valid: !charData.DateLastPlayed.IsZero()},
		EmblemBackgroundPath: sql.NullString{String: charData.EmblemBackground, Valid: true},
		TitleRecordHash:      int64(charData.TitleHash),
		EmblemColorA:         sql.NullInt64{Int64: charData.EmblemColor.A, Valid: true},
		EmblemColorR:         sql.NullInt64{Int64: charData.EmblemColor.R, Valid: true},
		EmblemColorG:         sql.NullInt64{Int64: charData.EmblemColor.G, Valid: true},
		EmblemColorB:         sql.NullInt64{Int64: charData.EmblemColor.B, Valid: true},
	},
		statsParam,
	)
}
