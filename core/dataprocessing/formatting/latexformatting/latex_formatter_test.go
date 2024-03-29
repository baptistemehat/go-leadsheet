package latexformatting

import (
	"testing"

	"github.com/baptistemehat/go-leadsheet/core/datamodel/music"
	"github.com/baptistemehat/go-leadsheet/core/datamodel/song"
)

type latexTestCase struct {
	chord    music.Chord
	expected string
}

var latexTestCases []latexTestCase = []latexTestCase{
	{music.NewChord(music.D, music.Minor, music.SeventhNone, []uint8{}, music.D), "\\chord{Dm}"},
	{music.NewChord(music.E, music.Major, music.MinorSeventh, []uint8{}, music.Gs), "\\chord{E7/G#}"},
}

func TestLatexEscape(t *testing.T) {
	input := "abc&efg%hij$klm#nop_qrs{tuv}wxy~zAB^CDE\\FGH"
	expected := "abc\\& efg\\% hij\\$ klm\\# nop\\_ qrs\\{ tuv\\} wxy\\textasciitilde zAB\\textasciicircum CDE\\textbackslash FGH"

	if actual := LatexEscape(input); actual != expected {
		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expected)
	}
}

// TODO : add error cases

func TestFormatChord(t *testing.T) {

	f := LatexSongFormatter{}

	for _, testCase := range latexTestCases {
		if actual, _ := f.FormatChord(&testCase.chord); actual != testCase.expected {
			t.Errorf("Actual: '%s' - Expected: '%s'", actual, testCase.expected)
		}
	}
}

func TestFormatLine_LyricsOnly(t *testing.T) {
	testCases := []struct {
		line     song.Line
		expected string
	}{
		{song.NewLine(), ""},
		{song.NewLine(), "Somewhere over the rainbow"},
	}

	f := LatexSongFormatter{}

	for _, testCase := range testCases {
		if actual, _ := f.FormatLine(&testCase.line); actual != testCase.expected {
			t.Errorf("Actual: '%s' - Expected: '%s'", actual, testCase.expected)
		}
	}
}

func TestFormatLine_ChordsOnly(t *testing.T) {
	// TODO
}

func TestFormatLine_LyricsAndChords(t *testing.T) {

	line := song.NewLine()
	line.AppendLyrics("Somewhere over the rainbow")
	line.AddChord(music.NewChord(music.G, music.Major, music.SeventhNone, []uint8{}, music.G), 0)
	line.AddChord(music.NewChord(music.B, music.Minor, music.SeventhNone, []uint8{}, music.B), 10)
	line.AddChord(music.NewChord(music.C, music.Major, music.SeventhNone, []uint8{}, music.C), 25)

	expectedString := "\\chord{G}Somewhere \\chord{Bm}over the rainbo\\chord{C}w \\\\"

	f := LatexSongFormatter{}

	if actual, _ := f.FormatLine(&line); actual != expectedString {
		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
	}
}

func TestFormatSection(t *testing.T) {

	section := song.NewSection()
	section.SetName("Verse")

	line := song.NewLine()
	line.AppendLyrics("Somewhere over the rainbow")
	line.AddChord(music.NewChord(music.G, music.Major, music.SeventhNone, []uint8{}, music.G), 0)
	line.AddChord(music.NewChord(music.B, music.Minor, music.SeventhNone, []uint8{}, music.B), 10)
	line.AddChord(music.NewChord(music.C, music.Major, music.SeventhNone, []uint8{}, music.C), 25)

	section.AddLine(line)

	expectedString := "\\begin{Verse}\n" +
		"\\chord{G}Somewhere \\chord{Bm}over the rainbo\\chord{C}w \\\\" + "\n" +
		"\\end{Verse}"

	f := LatexSongFormatter{}

	if actual, _ := f.FormatSection(&section); actual != expectedString {
		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
	}
}

func TestFormatSongProperties(t *testing.T) {
	songProperties := song.SongProperties{
		Title:    "Somewhere over the rainbow",
		Composer: "Israel Kamakawiwo'ole",
		Capo:     5,
		Key:      "G",
	}

	expectedString := "{\n" +
		"title = {Somewhere over the rainbow},\n" +
		"composer = {Israel Kamakawiwo'ole},\n" +
		"capo = {5},\n" +
		"key = {G},\n" +
		"}"

	f := LatexSongFormatter{}

	if actual, _ := f.FormatSongProperties(&songProperties); actual != expectedString {
		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
	}

}

func TestFormatSong(t *testing.T) {
	s := song.NewSong()

	songProperties := song.SongProperties{
		Title:    "Somewhere over the rainbow",
		Composer: "Israel Kamakawiwo'ole",
		Capo:     0,
	}

	s.SetProperties(songProperties)

	section := song.NewSection()
	section.SetName("Verse")

	line := song.NewLine()
	line.AppendLyrics("Somewhere over the rainbow")
	line.AddChord(music.NewChord(music.G, music.Major, music.SeventhNone, []uint8{}, music.G), 0)
	line.AddChord(music.NewChord(music.B, music.Minor, music.SeventhNone, []uint8{}, music.B), 10)
	line.AddChord(music.NewChord(music.C, music.Major, music.SeventhNone, []uint8{}, music.C), 25)

	section.AddLine(line)
	s.AddSection(section)

	expectedString := "\\begin{song}\n" +
		"{\n" +
		"title = {Somewhere over the rainbow},\n" +
		"composer = {Israel Kamakawiwo'ole},\n" +
		"capo = {0},\n" +
		"}\n" +
		"\\begin{Verse}\n" +
		"\\chord{G}Somewhere \\chord{Bm}over the rainbo\\chord{C}w \\\\" + "\n" +
		"\\end{Verse}\n" +
		"\\end{song}"

	f := LatexSongFormatter{}

	if actual, _ := f.FormatSong(&s); actual != expectedString {
		t.Errorf("Actual: '%s' - Expected: '%s'", actual, expectedString)
	}
}

func TestFull(t *testing.T) {
	// 	p := parser.Parser{}

	// 	song, _ := p.Parse(`

	// Title: Hotel California
	// Composer: Eagles
	// Capo: 0
	// Key: Bm

	// {Verse}

	// [Am] On a dark desert highway,[E7] cool wind in my hair
	// [G] Warm smell of colitas [D] rising up through the air
	// [F] Up ahead in the distance,[C] I saw a shimmering light
	// [Dm] My head grew heavy and my sight grew dim,[E7] I had to stop for the night

	// {Chorus}

	// [F] Welcome to the Hotel Califo[C]rnia.
	// Such a [E7]lovely place, (such a lovely place), such a [Am]lovely face
	// [F]Plenty of room at the Hotel Cali[C]fornia
	// Any [Dm]time of year, (any time of year) You can [E7]find it here

	// 	`)

	// 	f := LatexSongFormatter{}
	// 	actualString, _ := f.FormatSong(&song)
	// 	fmt.Println(actualString)

	// 	if err := os.WriteFile("../../latex/tmp/songs/leadsheet.tex", []byte(actualString), 0666); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	pg := pdfGenerator.PdfGenerator{}

	// 	fmt.Println(pg.Tex2pdf())

}
