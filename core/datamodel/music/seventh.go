package music

import "fmt"

// Seventh
type Seventh uint8

const (
	SeventhNone  Seventh = iota // none
	MinorSeventh                // 7
	MajorSeventh                // M7
)

var seventhToString = map[Seventh]string{
	SeventhNone:  "",
	MinorSeventh: "7",
	MajorSeventh: "M7",
}

// String
func (seventh Seventh) String() (string, error) {
	if result, ok := seventhToString[seventh]; !ok {
		return "", fmt.Errorf("illegal seventh: %d", seventh)
	} else {
		return result, nil
	}
}

var stringToSeventh = map[string]Seventh{
	"":   SeventhNone,
	"7":  MinorSeventh,
	"M7": MajorSeventh,
}

// StringToSeventh
func StringToSeventh(seventh string) (Seventh, error) {
	if result, ok := stringToSeventh[seventh]; !ok {
		return SeventhNone, fmt.Errorf("illegal seventh: %s", seventh)
	} else {
		return result, nil
	}
}
