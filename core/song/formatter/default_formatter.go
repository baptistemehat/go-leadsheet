package formatter

import (
	"github.com/baptistemehat/go-leadsheet/core/song/model"
)

type DefaultSongFormatter struct {
}

func (f *DefaultSongFormatter) FormatChord(chord *model.Chord) (string, error) {

	return chord.String()
}

func (f *DefaultSongFormatter) FormatLine(line *model.Line) (string, error) {
	return "", nil
}

func (f *DefaultSongFormatter) FormatSection(section *model.Section) (string, error) {
	return "", nil
}

func (f *DefaultSongFormatter) FormatSongProperties(songProperties *model.SongProperties) (string, error) {
	return "", nil
}

func (f *DefaultSongFormatter) FormatSong(song *model.Song) (string, error) {
	return "", nil
}
