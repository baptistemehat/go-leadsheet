package model

import "fmt"

type Note uint8

const (
	C  Note = iota // C
	Cs             // C#
	Db             // Db
	D              // D
	Ds             // D#
	Eb             // Eb
	E              // E
	F              // F
	Fs             // F#
	Gb             // Gb
	G              // G
	Gs             // G#
	Ab             // Ab
	A              // A
	As             // A#
	Bb             // Bb
	B              // B
)

var noteToString = map[Note]string{
	C:  "C",
	Cs: "C#",
	Db: "Db",
	D:  "D",
	Ds: "D#",
	Eb: "Eb",
	E:  "E",
	F:  "F",
	Fs: "F#",
	Gb: "Gb",
	G:  "G",
	Gs: "G#",
	Ab: "Ab",
	A:  "A",
	As: "A#",
	Bb: "Bb",
	B:  "B",
}

func (note Note) String() (string, error) {
	if result, ok := noteToString[note]; !ok {
		return "", fmt.Errorf("illegal note: %d", note)
	} else {
		return result, nil
	}
}

var stringToNote = map[string]Note{
	"C":  C,
	"C#": Cs,
	"Db": Db,
	"D":  D,
	"D#": Ds,
	"Eb": Eb,
	"E":  E,
	"F":  F,
	"F#": Fs,
	"Gb": Gb,
	"G":  G,
	"G#": Gs,
	"Ab": Ab,
	"A":  A,
	"A#": As,
	"Bb": Bb,
	"B":  B,
}

func StringToNote(note string) (Note, error) {

	if result, ok := stringToNote[note]; !ok {
		return C, fmt.Errorf("illegal note name: %s", note)
	} else {
		return result, nil
	}
}

func (note *Note) Transponse(n int8) Note {
	return C
}
