package songFormatter

type LatexSongFormatter struct {
}

// func (f *LatexSongFormatter) FormatChord(chord *chord.Chord) string {
// 	// // in leadsheet package, "^{}" can replace "\chord{}"
// 	// result := fmt.Sprintf("\\chord{%s}", chord.String())

// 	// return result
// 	return ""
// }

// func (f *LatexSongFormatter) FormatLine(line *song.Line) string {

// 	result := line.Lyrics

// 	// TODO: chord formatting changes if we are in Chord only or in mixed mode
// 	if !line.IsLyricsOnly() {
// 		offset := 0
// 		for _, item := range line.Chords {
// 			chordString := item.Chord.Format(f)
// 			result = result[:item.Location+uint8(offset)] + chordString + result[item.Location+uint8(offset):]
// 			offset += len(chordString)
// 		}
// 	}
// 	result += " \\\\"
// 	return result
// }

// func (f *LatexSongFormatter) FormatSection(section *song.Section) string {

// 	result := fmt.Sprintf("\\begin{%s}\n", section.Name)
// 	for _, line := range section.Lines {
// 		result += line.Format(f) + "\n"
// 	}
// 	result += fmt.Sprintf("\\end{%s}", section.Name)

// 	return result
// }

// func (f *LatexSongFormatter) FormatSongProperties(sp *song.SongProperties) string {

// 	result := "{\n"
// 	result += fmt.Sprintf("title = {%s},\n", sp.Title)
// 	result += fmt.Sprintf("composer = {%s},\n", sp.Composer)
// 	result += fmt.Sprintf("capo = {%d},\n", sp.Capo)
// 	result += "}"

// 	return result
// }

// func (f *LatexSongFormatter) FormatSong(song *song.Song) string {

// 	result := "\\begin{song}\n"
// 	result += song.Properties.Format(f) + "\n"
// 	for _, section := range song.Sections {
// 		result += section.Format(f) + "\n"
// 	}
// 	result += "\\end{song}"

// 	return result
// }
