package destiny

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ManifestItem struct {
	DisplayProperties struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"displayProperties"`
	ItemCategoryHashes  []uint32 `json:"itemCategoryHashes"`
	ItemTypeDisplayName string   `json:"itemTypeDisplayName"`
	ItemType            int      `json:"itemType"`
	Inventory           struct {
		TierTypeName   string `json:"tierTypeName"` // "Exotic"
		TierType       int    `json:"tierType"`
		BucketTypeHash uint32 `json:"bucketTypeHash"` // Para el Slot
	} `json:"inventory"`
	EquippingBlock struct {
		AmmoType int32 `json:"ammoType"` // 1: Primary, 2: Special, 3: Heavy
	} `json:"equippingBlock"`
	DefaultDamageTypeHash uint32 `json:"defaultDamageTypeHash"`
}

type DestinyRecordDefinition struct {
	DisplayProperties struct {
		Name string `json:"name"`
	} `json:"displayProperties"`
	TitleInfo struct {
		HasTitle       bool `json:"hasTitle"`
		TitlesByGender struct {
			Male   string `json:"Male"`
			Female string `json:"Female"`
		} `json:"titlesByGender"`
	} `json:"titleInfo"`
}

type ManifestRecord struct {
	titles map[uint32]string
}

func LoadManifestItem(filePath string) (map[uint32]ManifestItem, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	manifestMap := make(map[uint32]ManifestItem)

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&manifestMap)
	if err != nil {
		return nil, fmt.Errorf("error decoding manifest: %w", err)
	}

	return manifestMap, nil
}

func LoadManifestDefinition(path string) (*ManifestRecord, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Cargamos el JSON en un mapa de mensajes crudos (sin procesar todavía)
	var rawManifest map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawManifest); err != nil {
		return nil, err
	}

	processedTitles := make(map[uint32]string)

	for key, value := range rawManifest {
		var rec DestinyRecordDefinition
		// Solo parseamos lo necesario para ver si tiene título
		if err := json.Unmarshal(value, &rec); err != nil {
			continue
		}

		if rec.TitleInfo.HasTitle && rec.TitleInfo.TitlesByGender.Male != "" {
			// Convertimos la llave string a uint32
			hash, _ := strconv.ParseUint(key, 10, 32)
			processedTitles[uint32(hash)] = rec.TitleInfo.TitlesByGender.Male
		}
	}

	return &ManifestRecord{titles: processedTitles}, nil
}

// GetTitleName devuelve el nombre o un string vacío si no existe
func (m *ManifestRecord) GetTitleName(hash uint32) string {
	return m.titles[hash]
}
