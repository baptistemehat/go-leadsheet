package model

import (
	"testing"
)

type testCase struct {
	Input    string
	Expected Chord
}

var testCases = []testCase{
	{
		Input:    "Am",
		Expected: NewChord(A, Minor, SeventhNone, []uint8{}, A),
	},
	{
		Input:    "G7/B",
		Expected: NewChord(G, Major, SeventhNone, []uint8{}, B),
	},
}

func TestParseChord(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(testCase.Input, func(t *testing.T) {
			var actualChord Chord
			var error error
			if actualChord, error = ParseChord(testCase.Input); error != nil {
				t.Errorf("Error while parsing chord, %s", error)
			}
			if !actualChord.Equal(testCase.Expected) {
				expectedString, _ := testCase.Expected.String()
				actualString, _ := actualChord.String()
				t.Errorf("Expected: %s, Got: %s", expectedString, actualString)
			}
		})
	}
}
