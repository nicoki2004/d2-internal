package store

import (
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny/item"
)

type StoreDTO struct {
	CharacterID string
	Name        string
	IsVault     bool
	ClassType   int
	ClassName   string
	GenderType  string
	GenderName  string

	Emblem     Embleminfo
	Current    bool
	LastPlayed time.Time

	Level      int
	PowerLevel int
	TitleInfo  TitleInfo
	Stats      []CharacterStat

	Equiped []item.ItemDTO `json:"equiped"`
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
