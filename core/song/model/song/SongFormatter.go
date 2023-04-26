package song

type SongFormatter interface {
	FormatChord(*Chord) (string, error)
	FormatLine(*Line) (string, error)
	FormatSection(*Section) (string, error)
	FormatSongProperties(*SongProperties) (string, error)
	FormatSong(*Song) (string, error)
}
