package songFormatter

import "github.com/baptistemehat/go-leadsheet/core/song/model/song"

type latexTestCase struct {
	chord    song.Chord
	expected string
}

// var latexTestCases []latexTestCase = []latexTestCase{
// 	{*song.NewChord(song.C, song.Minor7, []uint8{}, song.B), "\\chord{Cm7/B}"},
// 	{*song.NewChord(song.G, song.HalfDiminished7, []uint8{}, song.G), "\\chord{Gm7b5}"},
// }

// func TestFormatChord(t *testing.T) {

// 	f := LatexSongFormatter{}

// 	for _, testCase := range latexTestCases {
// 		if actual := f.FormatChord(&testCase.chord); actual != testCase.expected {
// 			t.Errorf("Actual: '%s' - Expected: '%s'", actual, testCase.expected)
// 		}
// 	}
// }

// func TestFormatLine_LyricsOnly(t *testing.T) {
// 	testCases := []struct {
// 		line     song.Line
// 		expected string
// 	}{
// 		{*song.NewLine(""), ""},
// 		{*song.NewLine("Somewhere over the rainbow"), "Somewhere over the rainbow"},
// 	}

// 	f := LatexSongFormatter{}

// 	for _, testCase := range testCases {
// 		if actual := f.FormatLine(&testCase.line); actual != testCase.expected {
// 			t.Errorf("Actual: '%s' - Expected: '%s'", actual, testCase.expected)
// 		}
// 	}
// }

// func TestFormatLine_ChordsOnly(t *testing.T) {
// 	// TODO
// }

// func TestFormatLine_LyricsAndChords(t *testing.T) {

// 	line := song.NewLine("Somewhere over the rainbow")
// 	line.AddChord(song.NewChord(song.G, song.Major, []uint8{}, song.G), 0)
// 	line.AddChord(song.NewChord(song.B, song.Minor, []uint8{}, song.B), 10)
// 	line.AddChord(song.NewChord(song.C, song.Major, []uint8{}, song.C), 25)

// 	expectedString := "\\chord{G}Somewhere \\chord{Bm}over the rainbo\\chord{C}w \\\\"

// 	f := LatexSongFormatter{}

// 	if actual := f.FormatLine(line); actual != expectedString {
// 		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
// 	}
// }

// func TestFormatSection(t *testing.T) {

// 	section := song.NewSection("Verse")

// 	line := song.NewLine("Somewhere over the rainbow")
// 	line.AddChord(song.NewChord(song.G, song.Major, []uint8{}, song.G), 0)
// 	line.AddChord(song.NewChord(song.B, song.Minor, []uint8{}, song.B), 10)
// 	line.AddChord(song.NewChord(song.C, song.Major, []uint8{}, song.C), 25)

// 	section.AddLine(line)

// 	expectedString := "\\begin{Verse}\n" +
// 		"\\chord{G}Somewhere \\chord{Bm}over the rainbo\\chord{C}w \\\\" + "\n" +
// 		"\\end{Verse}"

// 	f := LatexSongFormatter{}

// 	if actual := f.FormatSection(section); actual != expectedString {
// 		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
// 	}
// }

// func TestFormatSongProperties(t *testing.T) {
// 	// TODO
// }

// func TestFormatSong(t *testing.T) {
// 	s := song.NewSong()

// 	songProperties := &song.SongProperties{
// 		Title:    "Somewhere over the rainbow",
// 		Composer: "Israel Kamakawiwo'ole",
// 		Capo:     0,
// 	}

// 	s.SetProperties(songProperties)

// 	section := song.NewSection("Verse")

// 	line := song.NewLine("Somewhere over the rainbow")
// 	line.AddChord(song.NewChord(song.G, song.Major, []uint8{}, song.G), 0)
// 	line.AddChord(song.NewChord(song.B, song.Minor, []uint8{}, song.B), 10)
// 	line.AddChord(song.NewChord(song.C, song.Major, []uint8{}, song.C), 25)

// 	section.AddLine(line)
// 	s.AddSection(section)

// 	expectedString := "\\begin{song}\n" +
// 		"{\n" +
// 		"title = {Somewhere over the rainbow},\n" +
// 		"composer = {Israel Kamakawiwo'ole},\n" +
// 		"capo = {0},\n" +
// 		"}\n" +
// 		"\\begin{Verse}\n" +
// 		"\\chord{G}Somewhere \\chord{Bm}over the rainbo\\chord{C}w \\\\" + "\n" +
// 		"\\end{Verse}\n" +
// 		"\\end{song}"

// 	f := LatexSongFormatter{}

// 	if actual := f.FormatSong(s); actual != expectedString {
// 		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
// 	}
// }
