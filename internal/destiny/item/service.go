package item

import (
	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/registry"
)

func MapEquippedItems(charID string, data destiny.ProfileResponse) []ItemDTO {
	items := []ItemDTO{}

	// 1. Buscamos qué tiene puesto el personaje (Componente 205)
	equippedData, ok := data.Response.CharacterEquipment.Data[charID]
	if !ok {
		return items
	}

	for _, item := range equippedData.Items {
		// 2. Datos Estáticos (Manifest)
		def, found := registry.GetItem(item.ItemHash)

		if !found {
			// Logueamos pero con Debug o una sola vez para no spamear
			logger.GetLogger().Warn("Registry miss: %d", item.ItemHash)
			continue
		}
		// def := manifest.GetItemDefinition(item.ItemHash)

		// 3. Datos Dinámicos (Instancia - Componente 300)
		instance, ok := data.Response.ItemComponents.Instances.Data[item.ItemInstanceID]
		if !ok {
			continue
		}

		typeName, slotName := destiny.GetItemIdentity(def.ItemCategoryHashes)

		dto := ItemDTO{
			InstanceID:         item.ItemInstanceID,
			Hash:               item.ItemHash,
			Name:               def.DisplayProperties.Name,
			Icon:               "https://www.bungie.net" + def.DisplayProperties.Icon,
			Power:              instance.PrimaryStat.Value,
			IsEquipped:         true,
			TierName:           def.Inventory.TierTypeName,
			IsExotic:           def.Inventory.TierType == 6,
			ItemCategoryHashes: def.ItemCategoryHashes,
			Slot:               slotName,
		}

		switch def.ItemType {
		case 3:
			dto.ItemType = ItemWeapon
			dto.WeaponStats = &WeaponSpec{
				TypeName:   typeName,
				DamageType: instance.DamageType,
				AmmoType:   int(def.EquippingBlock.AmmoType),
			}
		case 2:
			dto.ItemType = ItemArmor
			dto.ArmorStats = &ArmorSpec{
				SlotName: slotName,
			}
		}

		items = append(items, dto)
	}

	return items
}
