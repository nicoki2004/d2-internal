package destiny

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
