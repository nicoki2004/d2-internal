package destiny

type ItemSlot int

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

var StatHashToName = map[uint32]string{
	2996146975: "Mobility",
	392767087:  "Resilience",
	1943323491: "Recovery",
	1735777505: "Discipline",
	144602215:  "Intellect",
	4244567218: "Strength",
}
