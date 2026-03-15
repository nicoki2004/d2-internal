package character

import (
	"context"

	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/repository"
)

func GetCharacterProfile(ctx context.Context, repo repository.DestinyRepository) ([]CharacterDTO, error) {
	data, err := repo.GetCharactersWithStats(ctx)
	if err != nil {
		logger.GetLogger().Debug("Error getting All Characters: %v", err)
	}

	manifestRecord, err := destiny.LoadManifestDefinition("./definition_manifest.json")
	if err != nil {
		logger.GetLogger().Debug("Cannot get record Manifest - %v", err)
	}

	response := make(map[string]*CharacterDTO)

	// Range of all characters, to get a new slice for response
	for _, c := range data {
		// Review if the character exists in the MAP
		char, ok := response[c.CharacterID]
		// If not exists, create a new character.
		if !ok {
			var titleHash int64
			if h, ok := c.TitleRecordHash.(float64); ok {
				titleHash = int64(h)
			} else if h, ok := c.TitleRecordHash.(int64); ok {
				titleHash = h
			}
			char = &CharacterDTO{
				ID:               c.CharacterID,
				Class:            destiny.GetClassName(int(c.ClassType)),
				Light:            int(c.LightLevel),
				EmblemPath:       destiny.BungieCDN + c.EmblemUrl.String,
				EmblemBackground: destiny.BungieCDN + c.EmblemBackgroundPath.String,
				EmblemColor: destiny.Color{
					R: int64(c.EmblemColorR.Int64),
					G: int64(c.EmblemColorG.Int64),
					B: int64(c.EmblemColorB.Int64),
					A: int64(c.EmblemColorA.Int64),
				},
				TitleHash: manifestRecord.GetTitleName(uint32(titleHash)),
				Stats:     make(map[string]int),
			}

			// Asigno el Character to the map
			response[c.CharacterID] = char
		}
		// Agrego los stats
		var stahsHash int64
		if h, ok := c.StatHash.(int64); ok {
			stahsHash = int64(h)
		} else {
			stahsHash = 0
		}
		if name, exists := destiny.StatHashToName[uint32(stahsHash)]; exists {
			char.Stats[name] = int(c.StatValue.Int64)
		}

	}

	result := make([]CharacterDTO, 0, len(response))

	for _, char := range response {
		result = append(result, *char)
	}

	return result, nil
}
