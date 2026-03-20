package character

import (
	"github.com/nicoki2004/d2-internal/internal/destiny"
)

type CharacterDTO struct {
	ID               string        `json:"id"`
	Class            string        `json:"class"`
	Light            int           `json:"light"`             // Power actual equipado
	Stats            Stats         `json:"stats"`             // Mobility, Resilience, etc.
	EmblemPath       string        `json:"emblem_path"`       // iconPath
	EmblemBackground string        `json:"emblem_background"` // secondarySpecial (el banner ancho)
	EmblemColor      destiny.Color `json:"emblem_color"`      // color (RGBA para el fondo de texto)
	TitleHash        string        `json:"title"`             // titleRecordHash
	// MaxPower         float64       `json:"max_power"`         // Calculado por ti (tu "base" sin artefacto)
	// Equipped []item.ItemDTO `json:"equiped"`
	// Inventorey []item.Weapon `json:"inventory"`
}

type Stats map[string]int
