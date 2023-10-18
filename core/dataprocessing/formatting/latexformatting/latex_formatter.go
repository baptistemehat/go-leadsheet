package latexformatting

import (
	"fmt"
	"os"

	"github.com/baptistemehat/go-leadsheet/core/datamodel/music"
	"github.com/baptistemehat/go-leadsheet/core/datamodel/song"
	"golang.org/x/exp/slices"
)

type LatexSongFormatter struct {
}

var LATEX_ESCAPE_RUNES = map[rune]string{
	'&':  "\\&",
	'%':  "\\%",
	'$':  "\\$",
	'#':  "\\#",
	'_':  "\\_",
	'{':  "\\{",
	'}':  "\\}",
	'~':  "\\textasciitilde",
	'^':  "\\textasciicircum",
	'\\': "\\textbackslash",
}

// TODO: test this function
func LatexEscape(latexString string) string {
	result := ""
	for _, r := range latexString {
		if escapeSequence, found := LATEX_ESCAPE_RUNES[r]; found {
			result += escapeSequence + " "
		} else {
			result += string(r)
		}
	}
	return result
}

// FormatChord
func (f *LatexSongFormatter) FormatChord(chord *music.Chord) (string, error) {

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
func (f *LatexSongFormatter) FormatLine(line *song.Line) (string, error) {

	// Latex formatting : "\chord{Am} On a dark desert highway,\chord{E7} cool wind in my hair \\"

	// TODO: create a sort of proxy that would automatically apply escaping when getting lyrics, title, ...
	result := LatexEscape(line.Lyrics)

	// TODO: if only chords, special formatting
	// TODO: add test for this case
	if line.IsEmpty() {
		result = "~"
	} else if line.IsChordsOnly() {

		// TODO: only add \n if line is after a title
		result = "\\doublebar"

		// foreach chord in line
		for i, chordIndex := range line.Chords {

			if i > 0 {
				result += " \\normalbar"
			}

			// TODO : handle all errors
			// TODO : Create a dedicated formatter or formatting function here ?

			chordString, _ := chordIndex.Chord.String()
			result += fmt.Sprintf(" ~\\writechord{%s}", chordString)
		}
		result += " \\doublebar"

	} else if !line.IsLyricsOnly() { // if line contains chords

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

// cf. core/latex/static/config.tex
var DEFAULT_SECTION_NAMES = []string{
	"Intro",
	"Verse",
	"Chorus",
	"Bridge",
	"Solo",
	"Outro",
}

var userDefinedSectionNames []string

func AppendToFile(buffer, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	//TODO handle errors here
	file.WriteString(buffer)

	return nil
}

// FormatSection
func (f *LatexSongFormatter) FormatSection(section *song.Section) (string, error) {

	// Latex formatting : \begin{Verse} ... \end{Verse}

	if !slices.Contains(DEFAULT_SECTION_NAMES, section.Name) && !slices.Contains(userDefinedSectionNames, section.Name) {
		// TODO : create a section type
		// TODO : write a custom file, imported by config.tex which woulddefine new section types
		// Write input to file
		buffer := fmt.Sprintf("\\newversetype{%s}[template=SideName, after-label={}, name={%s}]", LatexEscape(section.Name), LatexEscape(section.Name))
		// TODO : make WriteStringToFile a util function (and make it append to file)
		if err := AppendToFile(buffer, "../tmp/latex/tmp/user_config.tex"); err != nil {
			//TODO error handling
		}
		userDefinedSectionNames = append(userDefinedSectionNames, section.Name)
	}

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
func (f *LatexSongFormatter) FormatSongProperties(sp *song.SongProperties) (string, error) {

	// Temporary implementation

	result := "{\n"
	result += fmt.Sprintf("title = {%s},\n", LatexEscape(sp.Title))
	result += fmt.Sprintf("composer = {%s},\n", LatexEscape(sp.Composer))
	result += fmt.Sprintf("capo = {%d},\n", sp.Capo)
	result += fmt.Sprintf("key = {%s},\n", LatexEscape(sp.Key))
	result += "}"

	return result, nil
}

// FormatSong
func (f *LatexSongFormatter) FormatSong(song *song.Song) (string, error) {

	err := os.MkdirAll("latex/tmp", os.ModePerm)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll("latex/tmp/out", os.ModePerm)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll("latex/tmp/songs", os.ModePerm)
	if err != nil {
		return "", err
	}

	// TODO centralize filepath of custom config file
	file, err := os.Create("../tmp/latex/tmp/user_config.tex")
	if err != nil {
		return "", err
	}
	file.Close()

	userDefinedSectionNames = []string{}

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
