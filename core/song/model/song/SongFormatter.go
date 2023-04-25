package song

type SongFormatter interface {
	FormatChord(*Chord) string
	FormatLine(*Line) string
	FormatSection(*Section) string
	FormatSongProperties(*SongProperties) string
	FormatSong(*Song) string
}
