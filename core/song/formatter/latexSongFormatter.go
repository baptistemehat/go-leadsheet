package formatter

import (
	"fmt"

	"github.com/baptistemehat/go-leadsheet/core/song/model"
)

type LatexSongFormatter struct {
}

func (f *LatexSongFormatter) FormatChord(chord *model.Chord) (string, error) {
	// in leadsheet package, "^{}" can replace "\chord{}"
	if chordString, err := chord.String(); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("\\chord{%s}", chordString), nil
	}
}

func (f *LatexSongFormatter) FormatLine(line *model.Line) (string, error) {

	result := line.Lyrics

	// TODO: chord formatting changes if we are in Chord only or in mixed mode
	if !line.IsLyricsOnly() {
		offset := 0
		for _, item := range line.Chords {
			chordString, _ := item.Chord.Format(f)
			result = result[:item.Index+uint8(offset)] + chordString + result[item.Index+uint8(offset):]
			offset += len(chordString)
		}
	}
	result += " \\\\"
	return result, nil
}

func (f *LatexSongFormatter) FormatSection(section *model.Section) (string, error) {

	result := fmt.Sprintf("\\begin{%s}\n", section.Name)
	for _, line := range section.Lines {
		lineString, _ := line.Format(f)
		result += lineString + "\n"
	}
	result += fmt.Sprintf("\\end{%s}", section.Name)

	return result, nil
}

func (f *LatexSongFormatter) FormatSongProperties(sp *model.SongProperties) (string, error) {

	result := "{\n"
	result += fmt.Sprintf("title = {%s},\n", sp.Title)
	result += fmt.Sprintf("composer = {%s},\n", sp.Composer)
	result += fmt.Sprintf("capo = {%d},\n", sp.Capo)
	result += fmt.Sprintf("key = {%s},\n", sp.Key)
	result += "}"

	return result, nil
}

func (f *LatexSongFormatter) FormatSong(song *model.Song) (string, error) {

	result := "\\begin{song}\n"
	songString, _ := song.Properties.Format(f)
	result += songString + "\n"
	for _, section := range song.Sections {
		sectionString, _ := section.Format(f)
		result += sectionString + "\n"
	}
	result += "\\end{song}"

	return result, nil
}
