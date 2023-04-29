package formatter

import (
	"fmt"

	"github.com/baptistemehat/go-leadsheet/core/song/model"
)

type LatexSongFormatter struct {
}

// FormatChord
func (f *LatexSongFormatter) FormatChord(chord *model.Chord) (string, error) {

	// Latex formatting : "\chord{E7}" OR "^{E7}"

	if chordString, err := chord.String(); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("\\chord{%s}", chordString), nil
	}
}

// InsertString inserts an insert string into a source string, at an index
func InsertString(source string, insert string, index uint8) (string, error) {
	// TODO : handle case where index > len(source)
	return source[:index] + insert + source[index:], nil
}

// FormatLine
func (f *LatexSongFormatter) FormatLine(line *model.Line) (string, error) {

	// Latex formatting : "\chord{Am} On a dark desert highway,\chord{E7} cool wind in my hair \\"

	result := line.Lyrics

	// if line contains chords
	if !line.IsLyricsOnly() {

		// tracks the cumulated length of chord strings
		chordsLength := uint8(0)

		// foreach chord in line
		for _, chordIndex := range line.Chords {

			// TODO : handle all errors
			chordString, _ := chordIndex.Chord.Format(f)

			// insert chord in lyrics
			result, _ = InsertString(result, chordString, chordIndex.Index+chordsLength)

			// increase length of chord strings
			chordsLength += uint8(len(chordString))
		}
	}

	result += " \\\\"
	return result, nil
}

// FormatSection
func (f *LatexSongFormatter) FormatSection(section *model.Section) (string, error) {

	// Latex formatting : \begin{Verse} ... \end{Verse}

	// \begin{}
	result := fmt.Sprintf("\\begin{%s}\n", section.Name)

	// foreach line in section
	for _, line := range section.Lines {
		// format line and add to section string
		lineString, _ := line.Format(f)
		result += lineString + "\n"
	}

	// \end{}
	result += fmt.Sprintf("\\end{%s}", section.Name)

	return result, nil
}

// FormatSongProperties
func (f *LatexSongFormatter) FormatSongProperties(sp *model.SongProperties) (string, error) {

	// Temporary implementation

	result := "{\n"
	result += fmt.Sprintf("title = {%s},\n", sp.Title)
	result += fmt.Sprintf("composer = {%s},\n", sp.Composer)
	result += fmt.Sprintf("capo = {%d},\n", sp.Capo)
	result += fmt.Sprintf("key = {%s},\n", sp.Key)
	result += "}"

	return result, nil
}

// FormatSong
func (f *LatexSongFormatter) FormatSong(song *model.Song) (string, error) {

	// Latex formatting : \begin{song} ... \end{song}

	// \begin{song}
	result := "\\begin{song}\n"

	// format properties
	songString, _ := song.Properties.Format(f)
	result += songString + "\n"

	// foreach section in song
	for _, section := range song.Sections {
		// format section and add to song string
		sectionString, _ := section.Format(f)
		result += sectionString + "\n"
	}

	// \end{song}
	result += "\\end{song}"

	return result, nil
}
