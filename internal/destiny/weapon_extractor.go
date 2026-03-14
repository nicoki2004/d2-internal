package destiny

// WeaponMetadata contains information about a weapon
type WeaponMetadata struct {
	Level    int
	Kills    int
	Progress float64
	Power    int
}

// WeaponExtractor extracts weapon information from the profile
type WeaponExtractor struct{}

// NewWeaponExtractor creates a new extractor
func NewWeaponExtractor() *WeaponExtractor {
	return &WeaponExtractor{}
}

// ExtractMetadata extracts all information about a weapon
func (we *WeaponExtractor) ExtractMetadata(item Item, profile *ProfileResponse) WeaponMetadata {
	metadata := WeaponMetadata{}

	// Get level, kills and progress
	level, kills, progress := GetWeaponMetadata(item.ItemInstanceId, profile)
	metadata.Level = level
	metadata.Kills = kills
	metadata.Progress = progress

	// Get power from itemComponents.instances
	if inst, ok := profile.Response.ItemComponents.Instances.Data[item.ItemInstanceId]; ok {
		metadata.Power = inst.PrimaryStat.Value
	}

	return metadata
}

// ExtractStats extracts weapon statistical stats
func (we *WeaponExtractor) ExtractStats(item Item, profile *ProfileResponse) map[string]int {
	stats := make(map[string]int)

	if statsData, ok := profile.Response.ItemComponents.Stats.Data[item.ItemInstanceId]; ok {
		for hash, stat := range statsData.Stats {
			if statName, exists := StatsDictionary[hash]; exists {
				stats[statName] = stat.Value
			}
		}
	}

	return stats
}

// ExtractSockets extracts weapon socket information
func (we *WeaponExtractor) ExtractSockets(item Item, profile *ProfileResponse) SocketsData {
	data := SocketsData{}

	socketsData, hasSockets := profile.Response.ItemComponents.Sockets.Data[item.ItemInstanceId]
	reusableData, hasReusable := profile.Response.ItemComponents.ReusablePlugs.Data[item.ItemInstanceId]

	data.HasSockets = hasSockets
	data.HasReusable = hasReusable
	data.Sockets = socketsData
	data.Plugs = reusableData

	return data
}

// SocketsData contains socket information
type SocketsData struct {
	HasSockets  bool
	HasReusable bool
	Sockets     ItemSocketsComponent
	Plugs       ItemReusablePlugsComponent
}

func GetWeaponMetadata(instanceId string, profile *ProfileResponse) (level int, kills int, progress float64) {
	// 1. Try component 301 (Standard weapons)
	if objData, ok := profile.Response.ItemComponents.Objectives.Data[instanceId]; ok {
		for _, obj := range objData.Objectives {
			processObjective(obj, &kills, &progress)
		}
	}

	// 2. Try component 309 (Crafted/enhanced weapons like your Commemoration)
	if plugData, ok := profile.Response.ItemComponents.ItemPlugObjectives.Data[instanceId]; ok {
		for _, objectivesList := range plugData.ObjectivesPerPlug {
			for _, obj := range objectivesList {
				processObjective(obj, &kills, &progress)
			}
		}
	}
	return level, kills, progress
}

func processObjective(obj ObjectiveData, kills *int, progress *float64) {
	switch obj.ObjectiveHash {
	// Kills tracker (Generic PvE/PvP counter)
	case 73837075:
		if obj.Progress > 0 {
			*kills = obj.Progress
		}
	// Other progress trackers (can be level or killcounts)
	case 562334711, 867865505, 1970111194:
		if obj.CompletionValue > 0 && obj.CompletionValue > 1 {
			// If completionValue > 1, it's a tracker with multiple values
			*progress = (float64(obj.Progress) / float64(obj.CompletionValue)) * 100
		} else if obj.Progress > 100 {
			// If progress > 100, it's probably kills
			*kills = obj.Progress
		}
	}
}
