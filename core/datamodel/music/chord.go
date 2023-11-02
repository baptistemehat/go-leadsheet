package music

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// https://en.wikipedia.org/wiki/Chord_notation
// ROOT NOTE
// Chord type (ie type of triad: major, minor, aug, dim)
// Extensions/Alterations : 5(power), b5, #5, 6, 9, #11, 6/9
// Suspensions : sus2, sus4
// Additions : add2, 6/9, add 9
// Bass note : /B
// This should work : C7b5#11sus4add9/Gb

// Chord
type Chord struct {
	Root        Note
	TriadType   Triad
	Seventh     Seventh
	Extensions  []uint8
	Suspensions []uint8
	Additions   []uint8
	BassNote    Note
}

// Formatter
type ChordFormatter interface {
	FormatChord(Chord) (string, error)
}

// NewChord creates a new chord
func NewChord(root Note, triadType Triad, seventh Seventh, extensions []uint8, bassNote Note) Chord {

	return Chord{
		Root:       root,
		TriadType:  triadType,
		Seventh:    seventh,
		Extensions: extensions,
		BassNote:   bassNote,
	}
}

// ParseChord
// temporary implementation
func ParseChord(chordString string) (Chord, error) {

	// TODO :
	// 1. make this parsing really complete
	// 2. create a "real" parser with tokens
	re := regexp.MustCompile(`([A-G][b#]?)(m|)(7|M7|)(/[A-G][b#]?)?`)

	tokens := re.FindStringSubmatch(chordString)
	tokens = tokens[1:]

	if len(tokens) < 3 || len(tokens) > 4 {
		return Chord{}, fmt.Errorf("illegal chord name: %s", chordString)
	}

	rootToken := tokens[0]
	var root Note
	var err error
	if root, err = StringToNote(rootToken); err != nil {
		return Chord{}, fmt.Errorf("")
	}

	triadTypeToken := tokens[1]
	var triadType Triad
	if triadType, err = StringToTriad(triadTypeToken); err != nil {
		return Chord{}, fmt.Errorf("")
	}

	seventhTypeToken := tokens[2]
	var seventhType Seventh
	if seventhType, err = StringToSeventh(seventhTypeToken); err != nil {
		return Chord{}, fmt.Errorf("")
	}

	if len(tokens) == 4 {
		if len(tokens[3]) > 0 {
			bassToken := tokens[3]
			var bass Note
			if bass, err = StringToNote(bassToken[1:]); err != nil {
				return Chord{}, fmt.Errorf("")
			}
			return NewChord(root, triadType, seventhType, []uint8{}, bass), nil

		}
	}

	return NewChord(root, triadType, seventhType, []uint8{}, root), nil
}

// Equal
// temporary implementation
func (chord Chord) Equal(otherChord Chord) bool {
	result := (chord.Root == otherChord.Root) && (chord.TriadType == otherChord.TriadType) && (chord.BassNote == otherChord.BassNote)

	for i := range chord.Extensions {
		if chord.Extensions[i] != otherChord.Extensions[i] {
			return false
		}
	}

	return result
}

// MarshalJSON
func (chord *Chord) MarshalJSON() ([]byte, error) {
	if result, err := chord.String(); err != nil {
		return []byte{}, err
	} else {

		return json.Marshal(result)
	}
}

// String
func (chord *Chord) String() (string, error) {
	var result, root, triadType, seventhType, bassNote string
	var err error

	if root, err = chord.Root.String(); err != nil {
		return "", err
	} else {
		result = root
	}

	if triadType, err = chord.TriadType.String(); err != nil {
		return "", err
	} else {
		result += triadType
	}

	if seventhType, err = chord.Seventh.String(); err != nil {
		return "", err
	} else {
		result += seventhType
	}

	if chord.BassNote != chord.Root {
		if bassNote, err = chord.BassNote.String(); err != nil {
			return "", err
		} else {
			result += "/" + bassNote
		}
	}

	return result, nil
}

// Format
func (c *Chord) Format(f ChordFormatter) (string, error) {
	return f.FormatChord(*c)
}
