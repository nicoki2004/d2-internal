package store

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
		return "Exo"
	default:
		return "Desconocido"
	}
}
