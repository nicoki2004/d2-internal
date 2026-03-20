package store

import (
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny/item"
)

type StoreDTO struct {
	CharacterID string
	Name        string
	// Make a Store Vault or Character
	IsVault bool
	// Class
	ClassType int
	ClassName string
	// Gender
	GenderType string
	GenderName string

	Emblem     Embleminfo
	Current    bool
	LastPlayed time.Time

	Level      int
	PowerLevel int
	TitleInfo  TitleInfo
	Stats      []CharacterStat

	Equipped []item.ItemDTO `json:"equiped"`
}

type Embleminfo struct {
	EmblemPath       string
	EmblemBackground string
	EmblemHash       string
}

type EmblemColor struct {
	Red   int
	Green int
	Blue  int
	Alpha int
}

type TitleInfo struct {
	Title                    string
	IsCompleted              bool
	GildedNum                int
	IsGildedForCurrentSeason bool
}

type CharacterStat struct {
	Hash  string
	Name  string
	Value int
}

var StatHashToName = map[uint32]string{
	2996146975: "Mobility",
	392767087:  "Resilience",
	1943323491: "Recovery",
	1735777505: "Discipline",
	144602215:  "Intellect",
	4244567218: "Strength",
}

func GetGenderName(genderType int) string {
	switch genderType {
	case 0:
		return "Male"
	case 1:
		return "Female"
	case 2:
		return "Non-binary"
	default:
		return "Unknown"
	}
}

func GetRaceName(raceType int) string {
	switch raceType {
	case 0:
		return "Humano"
	case 1:
		return "Awoken"
	case 2:
		// En la API 2 suele ser Exo
		return "Exo"
	default:
		return "Desconocido"
	}
}
