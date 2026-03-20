package store

import (
	"fmt"
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/destiny/item"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/registry"
)

func GetStores(data destiny.ProfileResponse) []StoreDTO {
	start := time.Now()
	startOne := time.Now()
	logger.GetLogger().Debug("Carga Manifest: %v", time.Since(start))

	start = time.Now()
	var lastPlayedGlobal time.Time
	for _, c := range data.Response.Characters.Data {
		if c.DateLastPlayed.After(lastPlayedGlobal) {
			lastPlayedGlobal = c.DateLastPlayed
		}
	}
	logger.GetLogger().Debug("Calculo de LastPlayed: %v", time.Since(start))

	result := make([]StoreDTO, 0, len(data.Response.Characters.Data))
	start = time.Now()
	for _, c := range data.Response.Characters.Data {
		charStats := make([]CharacterStat, 0, len(c.Stats))
		for hash, value := range c.Stats {
			if hash == 1935470627 {
				continue
			}
			charStats = append(charStats, CharacterStat{
				Hash:  fmt.Sprintf("%d", hash),
				Name:  StatHashToName[hash],
				Value: value,
			})
		}

		logger.GetLogger().Debug("Calculo de Stats: %v", time.Since(start))

		title := TitleInfo{
			Title: "None",
		}

		start = time.Now()
		if c.TitleHash > 0 {
			if record, ok := data.Response.ProfileRecords.Data.Records[c.TitleHash]; ok {
				title.Title = registry.Manifest.GetTitleName(c.TitleHash)
				title.IsCompleted = (record.State & 1) == 0
				title.GildedNum = record.CompletedCount

				if len(record.Objectives) > 0 {
					title.IsGildedForCurrentSeason = record.Objectives[0].Complete
				}
			}
		}

		logger.GetLogger().Debug("Calculo de Title: %v", time.Since(start))

		start = time.Now()
		store := StoreDTO{
			CharacterID: c.CharacterID,
			IsVault:     false,
			ClassType:   int(c.ClassType),
			ClassName:   destiny.GetClassName(int(c.ClassType)),
			Emblem: Embleminfo{
				EmblemPath:       c.EmblemPath,
				EmblemBackground: c.EmblemBackground,
			},
			Current:    c.DateLastPlayed.Equal(lastPlayedGlobal),
			LastPlayed: c.DateLastPlayed,
			Level:      c.Light,
			PowerLevel: c.Light,
			TitleInfo:  title,
			Stats:      charStats,
			GenderType: fmt.Sprintf("%d", c.GenderType),
			GenderName: GetGenderName(int(c.GenderType)),
		}

		logger.GetLogger().Debug("Calculo de Store: %v", time.Since(start))

		store.Equipped = item.MapEquippedItems(store.CharacterID, data)

		result = append(result, store)
	}
	vaultStore := StoreDTO{
		CharacterID: "0",
		Name:        "Vault",
		IsVault:     true,
		ClassType:   3,
		ClassName:   "Vault",
		Emblem: Embleminfo{
			EmblemPath:       "/common/destiny2_content/icons/icon_vault.png",
			EmblemBackground: "/common/destiny2_content/icons/icon_vault_background.png",
		},
		Current:    false,
		Level:      0,
		PowerLevel: 0,

		Stats:     []CharacterStat{},
		TitleInfo: TitleInfo{Title: "Account Wide"},
	}

	result = append(result, vaultStore)

	logger.GetLogger().Debug("Calculo Total: %v", time.Since(startOne))
	return result
}
