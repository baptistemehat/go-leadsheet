package song

import (
	"github.com/baptistemehat/go-leadsheet/core/datamodel/music"
)

// Song
type Song struct {
	Properties SongProperties `json:"properties"`
	Sections   []Section      `json:"sections"`
}

// Formatter
type SongFormatter interface {
	FormatChord(*music.Chord) (string, error)
	FormatLine(*Line) (string, error)
	FormatSection(*Section) (string, error)
	FormatSongProperties(*SongProperties) (string, error)
	FormatSong(*Song) (string, error)
}

// NewSong
func NewSong() Song {
	return Song{
		Properties: SongProperties{},
		Sections:   []Section{},
	}
}

// AddSection
func (song *Song) AddSection(section Section) {
	song.Sections = append(song.Sections, section)
}

// SetProperies
func (song *Song) SetProperties(sp SongProperties) {
	song.Properties = sp
}

// Format
func (song *Song) Format(f SongFormatter) (string, error) {
	return f.FormatSong(song)
}
