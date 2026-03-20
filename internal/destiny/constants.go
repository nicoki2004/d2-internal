package destiny

import "fmt"

const (
	PlayStation = 2
	Steam       = 3
)

const (
	// Profile (100-199)
	ProfilesComponent           = "100"
	VendorReceiptsComponent     = "101"
	ProfileInventoriesComponent = "102"
	ProfileCurrenciesComponent  = "103"
	ProfileProgressionComponent = "104"
	PlatformSilverComponent     = "105"

	// Characters (200-299)
	CharactersComponent            = "200"
	CharacterInventoriesComponent  = "201"
	CharacterProgressionsComponent = "202"
	CharacterRenderDataComponent   = "203"
	CharacterActivitiesComponent   = "204"
	CharacterEquipmentComponent    = "205"
	CharacterLoadoutsComponent     = "206"

	// Items (300-399)
	ItemInstancesComponent            = "300"
	ItemObjectivesComponentNumber     = "301"
	ItemPerksComponent                = "302"
	ItemRenderDataComponent           = "303"
	ItemStatsComponentNumber          = "304"
	ItemSocketsComponentNumber        = "305"
	ItemTalentGridsComponent          = "306"
	ItemCommonDataComponent           = "307"
	ItemPlugStatesComponent           = "308"
	ItemPlugObjectivesComponentNumber = "309"
	ItemReusablePlugsComponentNumber  = "310"

	// Vendors & Social (400-599)
	VendorsComponent          = "400"
	VendorCategoriesComponent = "401"
	VendorSalesComponent      = "402"
	KiosksComponent           = "500"
	CurrencyLookupsComponent  = "600"

	// Collections & Triumphs (800-1100)
	PresentationNodesComponent = "700"
	CollectiblesComponent      = "800"
	RecordsComponent           = "900"
	TransitoryComponent        = "1000"
	MetricsComponent           = "1100"
	StringVariablesComponent   = "1200"
	CraftablesComponent        = "1300"
)

var StatsDictionary = map[uint32]string{
	// --- Armas de Fuego ---
	4232813984: "RPM",          // Velocidad de fuego
	4043523819: "Impact",       // Impacto
	1240592695: "Range",        // Alcance
	155624089:  "Stability",    // Estabilidad
	943549884:  "Handling",     // Manejo
	4284893193: "Reload Speed", // Velocidad de recarga
	3871231066: "Magazine",     // Capacidad de cargador

	// --- Stats Técnicos / Extra ---
	1345609583: "Aim Assistance",   // Asistencia de puntería
	2714457168: "Airborne",         // Precisión en aire
	3555269338: "Zoom",             // Zoom
	2715839340: "Recoil Direction", // Dirección del retroceso

	// --- Espadas ---
	2837207746: "Swing Speed",      // Velocidad de oscilación
	925767036:  "Ammo Capacity",    // Capacidad de munición
	419712076:  "Guard Resistance", // Resistencia de guardia
	3022301683: "Charge Rate",      // Velocidad de carga
	3736848092: "Guard Endurance",  // Resistencia de guardia
}

const BungieCDN = "https://www.bungie.net"

func GetSlotName(hash uint32) string {
	switch hash {
	case 1498876634:
		return "Kinetic" // Slot 1
	case 2465295065:
		return "Energy" // Slot 2
	case 953998645:
		return "Power" // Slot 3
	default:
		return "Other"
	}
}

// GetDamageName Translate hash to Damage Name
func GetDamageName(hash uint32) string {
	switch hash {
	case 1847026933, 3:
		return "Solar"
	case 2301139358, 2303181850, 2:
		return "Arc"
	case 3454344763, 3454344768, 4:
		return "Void"
	case 1513472331, 151347233, 6:
		return "Stasis"
	case 3946443463, 3949783978, 7:
		return "Strand"
	case 3373582059, 1, 0:
		return "Kinetic"
	default:
		return "Kinetic" // If it's 0 or unknown, the vast majority are kinetic
	}
}

// GetClassName returns the guardian name based on its type (Titan, Hunter, Warlock)
func GetClassName(classType int) string {
	return classNames[uint32(classType)]
}

var classNames = map[uint32]string{
	0: "Titan",
	1: "Hunter",
	2: "Warlock",
}

var StatHashToName = map[uint32]string{
	2996146975: "Mobility",
	392767087:  "Resilience",
	1943323491: "Recovery",
	1735777505: "Discipline",
	144602215:  "Intellect",
	4244567218: "Strength",
}

type ItemSlot int

// ItemCategoryHashes: Para saber QUÉ es el objeto (Slot y Tipo)
const (
	// Categorías Generales
	CategoryWeapon ItemSlot = iota
	SlotUnknown
	CategoryArmor

	// Slots de Armadura (Siguen funcionando igual para todas las clases)
	SlotHelmet
	SlotGauntlets
	SlotChest
	SlotLegs
	SlotClassItem
	SlotSubclass
	SlotKinetic
	SlotEnergy
	SlotPower

	// Tipos de Arma (Los más comunes para tu UI)
	TypeAutoRifle
	TypeHandCannon
	TypePulseRifle
	TypeScoutRifle
	TypeFusionRifle
	TypeSniperRifle
	TypeShotgun
	TypeMachineGun
	TypeRocketLauncher
	TypeSidearm
	TypeSword
	TypeGrenadeLauncher
	TypeBow
	TypeGlaive
	TypeLinearFusion
	TypeSubmachineGun
	TypeTraceRifle

	SlotGhost
	SlotVehicle
	SlotShip
	SlotEmblem
	SlotClanBanner
	SlotFinisher
	SlotEmote
	SlotArtifact
)

var HashToSlot = map[uint32]ItemSlot{
	2:  SlotKinetic,
	3:  SlotEnergy,
	4:  SlotPower,
	45: SlotHelmet,
	46: SlotGauntlets,
	47: SlotChest,
	48: SlotLegs,
	49: SlotClassItem,
	50: SlotSubclass,

	39:         SlotGhost,
	43:         SlotVehicle,
	42:         SlotShip,
	19:         SlotEmblem,
	58:         SlotClanBanner,
	1112488720: SlotFinisher,
	44:         SlotEmote,
	1378222069: SlotArtifact,
}

var HashToType = map[uint32]ItemSlot{
	5:          TypeAutoRifle,
	6:          TypeHandCannon,
	7:          TypePulseRifle,
	8:          TypeScoutRifle,
	9:          TypeFusionRifle,
	10:         TypeSniperRifle,
	11:         TypeShotgun,
	12:         TypeMachineGun,
	13:         TypeRocketLauncher,
	14:         TypeSidearm,
	54:         TypeSword,
	3317538576: TypeBow,
	3954685534: TypeSubmachineGun,
	153950757:  TypeGrenadeLauncher,
	3871742104: TypeGlaive,
	1504945536: TypeLinearFusion,
	2489664120: TypeTraceRifle,
}

var SlotNames = map[ItemSlot]string{
	CategoryWeapon: "Weapon",
	CategoryArmor:  "Armor",

	SlotHelmet:    "Helmet",
	SlotGauntlets: "Gauntlets",
	SlotChest:     "Chest",
	SlotLegs:      "Legs",
	SlotClassItem: "ClassItem",
	SlotSubclass:  "SubClass",
	SlotKinetic:   "Kinetic",
	SlotEnergy:    "Energy",
	SlotPower:     "Power",

	TypeAutoRifle:       "Auto Rifle",
	TypeHandCannon:      "Hand Cannon",
	TypePulseRifle:      "Pulse Rifle",
	TypeScoutRifle:      "Scout Rifle",
	TypeFusionRifle:     "Fusion Rifle",
	TypeSniperRifle:     "Sniper Rifle",
	TypeShotgun:         "Shotgun",
	TypeMachineGun:      "Machine Gun",
	TypeRocketLauncher:  "Rocket Launcher",
	TypeSidearm:         "Sidearms",
	TypeSword:           "Sword",
	TypeSubmachineGun:   "Submachine Gun",
	SlotUnknown:         "Unknown",
	TypeGrenadeLauncher: "Grenade Launcher",
	TypeBow:             "Bow",
	TypeGlaive:          "Glaive",
	TypeLinearFusion:    "Linear Fusion Rifle",
	TypeTraceRifle:      "Trace Rifle",

	SlotGhost:      "Ghost Shell",
	SlotVehicle:    "Vehicle",
	SlotShip:       "Ship",
	SlotEmblem:     "Emblem",
	SlotClanBanner: "Clan Banner",
	SlotFinisher:   "Finisher",
	SlotEmote:      "Emote",
	SlotArtifact:   "Seasonal Artifact",
}

func GetItemIdentity(hashes []uint32) (string, string) {
	detectedType := SlotUnknown
	detectedSlot := SlotUnknown

	for _, h := range hashes {

		if slotName, ok := HashToSlot[h]; ok {
			detectedSlot = slotName
		}

		if typeName, ok := HashToType[h]; ok {
			detectedType = typeName
		}
	}

	if detectedType == SlotUnknown {
		detectedType = detectedSlot
	}

	return SlotNames[detectedType], SlotNames[detectedSlot]
}

func (s ItemSlot) MarshalJSON() ([]byte, error) {
	name, ok := SlotNames[s]
	if !ok {
		name = "Unknown"
	}
	return fmt.Appendf(nil, "%q", name), nil
}

const (
	PlugMasterworkArmor  = 2457930460
	PlugMasterworkWeapon = 782502718

	PlugSubclassAspect   = 3032847657
	PlugSubclassFragment = 1920373979

	PlugTitanPrismAspect   = 912150793
	PlugHunterPrismAspect  = 1164816619
	PlugWarlockPrismAspect = 3154627255
)

const (
	TierExotic    = 6
	TierLegendary = 5
	TierRare      = 4
)
