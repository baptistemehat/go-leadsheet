package model

// ChordIndex
type ChordIndex struct {
	Index uint8 `json:"index"`
	Chord Chord `json:"chord"`
}

// Line
type Line struct {
	Chords []ChordIndex `json:"chords"`
	Lyrics string       `json:"lyrics"`
}

// NewLine
func NewLine() Line {
	return Line{
		Lyrics: "",
	}
}

// AppendLyrics
func (line *Line) AppendLyrics(lyrics string) {
	line.Lyrics += lyrics
}

// AddChord
// temporary implementation
func (line *Line) AddChord(chord Chord, index uint8) {
	line.Chords = append(line.Chords, ChordIndex{Index: index, Chord: chord})

	/*
		index := -1
		for i := 0; i < len(line.Chords) && index < 0; i++ {
			if location <= line.Chords[i].Location {
				index = i
			}
		}

		if index < 0 {
			index = 0
		}

		line.Chords = append(line.Chords, struct {
			Location uint8
			Chord    Chord
		}{})

		copy(line.Chords[index+1:], line.Chords[index:])
		line.Chords[index] = struct {
			Location uint8
			Chord    Chord
		}{
			Location: location,
			Chord:    *chord,
		}
	*/
}

// Clear
func (line *Line) Clear() {
	line.Chords = []ChordIndex{}
	line.Lyrics = ""
}

// IsLyricsOnly
func (line *Line) IsLyricsOnly() bool {
	return len(line.Lyrics) > 0 && len(line.Chords) == 0
}

// IsChordsOnly
func (line *Line) IsChordsOnly() bool {
	return len(line.Chords) > 0 && len(line.Lyrics) == 0
}

func (line *Line) IsEmpty() bool {
	return len(line.Chords) == 0 && len(line.Lyrics) == 0
}

// Format
func (line *Line) Format(f Formatter) (string, error) {
	return f.FormatLine(line)
}
