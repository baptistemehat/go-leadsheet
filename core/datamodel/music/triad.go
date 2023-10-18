package music

import "fmt"

// Triad
type Triad uint8

const (
	Major Triad = iota
	Minor
	Augmented
	Diminished
)

var triadToString = map[Triad]string{
	Major:      "",
	Minor:      "m",
	Augmented:  "aug",
	Diminished: "dim",
}

// String
func (triad Triad) String() (string, error) {
	if result, ok := triadToString[triad]; !ok {
		return "", fmt.Errorf("illegal triad type: %d", triad)
	} else {
		return result, nil
	}
}

var stringToTriad = map[string]Triad{
	"":    Major,
	"m":   Minor,
	"aug": Augmented,
	"dim": Diminished,
}

// StringToTriad
func StringToTriad(triad string) (Triad, error) {
	if result, ok := stringToTriad[triad]; !ok {
		return 0, fmt.Errorf("illegal triad type: %s", triad)
	} else {
		return result, nil
	}
}
