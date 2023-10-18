package parsing

import "github.com/baptistemehat/go-leadsheet/core/datamodel/song"

// Parser
type LeadsheetParser interface {
	Parse(string) (song.Song, error)
}
