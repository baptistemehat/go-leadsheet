package songFormatter

import (
	"github.com/baptistemehat/go-leadsheet/core/song/model/song"
)

type DefaultSongFormatter struct {
}

func (f *DefaultSongFormatter) FormatChord(chord *song.Chord) (string, error) {

	return chord.String()
}

func (f *DefaultSongFormatter) FormatLine(line *song.Line) (string, error) {
	return "", nil
}

func (f *DefaultSongFormatter) FormatSection(section *song.Section) (string, error) {
	return "", nil
}

func (f *DefaultSongFormatter) FormatSongProperties(songProperties *song.SongProperties) (string, error) {
	return "", nil
}

func (f *DefaultSongFormatter) FormatSong(song *song.Song) (string, error) {
	return "", nil
}
