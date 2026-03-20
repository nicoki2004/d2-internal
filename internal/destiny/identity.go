package destiny

import "fmt"

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

var classNames = map[uint32]string{
	0: "Titan",
	1: "Hunter",
	2: "Warlock",
}

// GetClassName returns the guardian name based on its type (Titan, Hunter, Warlock)
func GetClassName(classType int) string {
	return classNames[uint32(classType)]
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
