package songFormatter

import "github.com/baptistemehat/go-leadsheet/core/song/model/song"

type testCase struct {
	chord    song.Chord
	expected string
}

// var testCases = []testCase{
// 	{*chord.NewChord(song.C, song.Minor, []uint8{}, song.C), "Cm"},
// 	{*song.NewChord(song.D, song.HalfDiminished7, []uint8{}, song.D), "Dm7b5"},
// }

// func TestChordToString(t *testing.T) {

// 	f := &DefaultSongFormatter{}

// 	for _, testCase := range testCases {
// 		if actual := testCase.chord.Format(f); actual != testCase.expected {
// 			t.Errorf("Output %s not equal to expected %s", actual, testCase.expected)
// 		}
// 	}

// }
