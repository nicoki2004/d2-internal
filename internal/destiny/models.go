package destiny

import "time"

// MembershipResponse contains Bungie subscription information for the user
// Returns all Destiny accounts connected to the Bungie account
type MembershipResponse struct {
	Response struct {
		DestinyMemberships  []DestinyMembership `json:"destinyMemberships"`
		PrimaryMembershipId string              `json:"primaryMembershipId"`
	} `json:"Response"`
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
}

// DestinyMembership represents an individual gaming platform (PS5, Xbox, Steam, etc)
type DestinyMembership struct {
	MembershipType          int    `json:"membershipType"`
	MembershipId            string `json:"membershipId"`
	DisplayName             string `json:"displayName"`
	BungieGlobalDisplayName string `json:"bungieGlobalDisplayName"`
}

// ProfileResponse is the complete response from Bungie's Profile endpoint
// Contains characters, inventory, equipment, stats and all technical details
type ProfileResponse struct {
	Response struct {
		// 100 - PROFILE INFORMATION
		Characters struct {
			Data map[string]CharacterData `json:"data"`
		} `json:"characters"`

		// 102 - THE VAULT (STORAGE)
		ProfileInventory struct {
			Data struct {
				Items []Item `json:"items"`
			} `json:"data"`
		} `json:"profileInventory"`

		// 201 - THE BACKPACKS (INVENTORIES)
		CharacterInventories struct {
			Data map[string]CharacterEquipmentData `json:"data"`
		} `json:"characterInventories"`
		// 205 - EQUIPPED ITEMS
		CharacterEquipment struct {
			Data map[string]CharacterEquipmentData `json:"data"`
		} `json:"characterEquipment"`

		// 300-309 - TECHNICAL DETAILS
		ItemComponents struct {
			Instances struct {
				Data map[string]ItemInstanceData `json:"data"`
			} `json:"instances"`

			Objectives struct {
				Data map[string]ItemObjectivesComponent `json:"data"`
			} `json:"objectives"`

			Stats struct {
				Data map[string]ItemStatsComponent `json:"data"`
			} `json:"stats"`

			// COMPONENT 309: The one you need for the kills tracker
			ItemPlugObjectives struct {
				Data map[string]ItemPlugObjectivesComponent `json:"data"`
			} `json:"plugObjectives"`
			// Other optional but recommended
			Sockets struct {
				Data map[string]ItemSocketsComponent `json:"data"`
			} `json:"sockets"`

			ReusablePlugs struct {
				Data map[string]ItemReusablePlugsComponent `json:"data"`
			} `json:"reusablePlugs"`
		} `json:"itemComponents"`
	} `json:"Response"`
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
}

// CharacterData represents a user character (Component 200)
type CharacterData struct {
	CharacterId      string    `json:"characterId"`
	ClassType        int       `json:"classType"`
	Light            int       `json:"light"`
	DateLastPlayed   time.Time `json:"dateLastPlayed"`
	EmblemPath       string    `json:"emblemPath"`
	Stats            Stats     `json:"stats"`
	EmblemBackground string    `json:"emblemBackgroundPath"`
	EmblemColor      Color     `json:"emblemColor"`
	TitleHash        uint32    `json:"titleRecordHash"`
}

type Stats map[uint32]int

type Color struct {
	R int64 `json:"red"`
	G int64 `json:"green"`
	B int64 `json:"blue"`
	A int64 `json:"alpha"`
}

// Item representa un elemento en el inventario (arma, armadura, consumible, etc)
type Item struct {
	ItemHash       uint32 `json:"itemHash"`       // Definition hash for the item (reference to manifest)
	ItemInstanceId string `json:"itemInstanceId"` // Unique ID for this specific instance
}

type CharacterEquipmentData struct {
	Items []Item `json:"items"`
}

// For Power (Component 300)
type ItemInstanceData struct {
	PrimaryStat struct {
		Value int `json:"value"` // Power lives here (460)
	} `json:"primaryStat"`
}

// For Kills and Level (Component 301)
type ItemObjectivesComponent struct {
	Objectives []ObjectiveData `json:"objectives"`
}

type ObjectiveData struct {
	ObjectiveHash   uint32 `json:"objectiveHash"`   // MUST be camelCase
	Progress        int    `json:"progress"`        // MUST be lowercase
	CompletionValue int    `json:"completionValue"` // MUST be camelCase
	Complete        bool   `json:"complete"`
	Visible         bool   `json:"visible"`
}

// ItemStatsComponent represents the stats of an instance (Component 304)
type ItemStatsComponent struct {
	Stats map[uint32]StatData `json:"stats"`
}

// StatData contains the individual stat value
type StatData struct {
	StatHash uint32 `json:"statHash"`
	Value    int    `json:"value"`
}

// Auxiliary struct for 309
type ItemPlugObjectivesComponent struct {
	// CHANGE: Bungie calls it "objectivesPerPlug" internally for each socket
	ObjectivesPerPlug map[string][]ObjectiveData `json:"objectivesPerPlug"`
}

// This one stays clean, only with the sockets
type ItemSocketsComponent struct {
	Sockets []SocketEntry `json:"sockets"`
}

type SocketEntry struct {
	PlugHash  uint32 `json:"plugHash"` // The Hash of the Perk or Modifier
	IsEnabled bool   `json:"isEnabled"`
	IsVisible bool   `json:"isVisible"`
}

// These remain the same
type ItemReusablePlugsComponent struct {
	Plugs map[string][]PlugEntry `json:"plugs"`
}

type PlugEntry struct {
	PlugItemHash uint32 `json:"plugItemHash"`
	CanInsert    bool   `json:"canInsert"`
}
