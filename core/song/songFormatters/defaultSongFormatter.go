package songFormatter

import (
	"github.com/baptistemehat/go-leadsheet/core/song/model/song"
)

type DefaultSongFormatter struct {
}

func (f *DefaultSongFormatter) FormatChord(chord *song.Chord) string {
	// result := chord.Root.String()

	// result += chord.ChordType.String()

	// // TODO : handle extensions, 9, 11, add9, 6/9

	// if chord.BassNote != chord.Root {
	// 	result += "/" + chord.BassNote.String()
	// }

	// return result
	return ""
}

func (f *DefaultSongFormatter) FormatLine(line *song.Line) string {
	return ""
}

func (f *DefaultSongFormatter) FormatSection(section *song.Section) string {
	return ""
}

func (f *DefaultSongFormatter) FormatSongProperties(songProperties *song.SongProperties) string {
	return ""
}

func (f *DefaultSongFormatter) FormatSong(song *song.Song) string {
	return ""
}
