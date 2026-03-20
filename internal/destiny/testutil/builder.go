package testutil

import (
	"time"

	"github.com/nicoki2004/d2-internal/internal/destiny"
)

type ProfileBuilder struct {
	profile destiny.ProfileResponse
}

func NewProfileBuilder() *ProfileBuilder {
	b := &ProfileBuilder{}
	b.profile.Response.Characters.Data = map[string]destiny.CharacterData{}
	b.profile.Response.CharacterEquipment.Data = map[string]destiny.CharacterEquipmentData{}
	b.profile.Response.ItemComponents.Instances.Data = map[string]destiny.ItemInstanceData{}
	b.profile.Response.ProfileRecords.Data.Records = map[uint32]destiny.RecordComponent{}
	return b
}

func (b *ProfileBuilder) WithCharacter(
	charID string,
	classType int,
	light int,
	lastPlayed time.Time,
	genderType int,
) *ProfileBuilder {
	b.profile.Response.Characters.Data[charID] = destiny.CharacterData{
		CharacterID:      charID,
		ClassType:        classType,
		Light:            light,
		DateLastPlayed:   lastPlayed,
		EmblemPath:       "/emblem.png",
		EmblemBackground: "/bg.png",
		Stats:            destiny.Stats{},
		GenderType:       genderType,
	}
	return b
}

func (b *ProfileBuilder) WithStat(charID string, hash uint32, value int) *ProfileBuilder {
	c := b.profile.Response.Characters.Data[charID]
	if c.Stats == nil {
		c.Stats = destiny.Stats{}
	}
	c.Stats[hash] = value
	b.profile.Response.Characters.Data[charID] = c
	return b
}

func (b *ProfileBuilder) WithEquippedItem(charID string, itemHash uint32, instanceID string) *ProfileBuilder {
	data := b.profile.Response.CharacterEquipment.Data[charID]
	data.Items = append(data.Items, destiny.Item{
		ItemHash:       itemHash,
		ItemInstanceID: instanceID,
	})
	b.profile.Response.CharacterEquipment.Data[charID] = data
	return b
}

func (b *ProfileBuilder) WithInstance(instanceID string, damageType int, power int) *ProfileBuilder {
	b.profile.Response.ItemComponents.Instances.Data[instanceID] = destiny.ItemInstanceData{
		DamageType: damageType,
		PrimaryStat: struct {
			StatHash uint32 `json:"statHash"`
			Value    int    `json:"value"`
		}{
			Value: power,
		},
	}
	return b
}

func (b *ProfileBuilder) Build() destiny.ProfileResponse {
	return b.profile
}
