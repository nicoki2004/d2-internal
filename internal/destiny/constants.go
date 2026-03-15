package destiny

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
