package character

import (
	"github.com/nicoki2004/d2-internal/internal/destiny"
)

type CharacterDTO struct {
	ID               string        `json:"id"`
	Class            string        `json:"class"`
	Light            int           `json:"light"`
	Stats            Stats         `json:"stats"`
	EmblemPath       string        `json:"emblem_path"`
	EmblemBackground string        `json:"emblem_background"`
	EmblemColor      destiny.Color `json:"emblem_color"`
	TitleHash        string        `json:"title"`
	// MaxPower         float64       `json:"max_power"`
	// Equipped []item.ItemDTO `json:"equiped"`
	// Inventorey []item.Weapon `json:"inventory"`
}

type Stats map[string]int
