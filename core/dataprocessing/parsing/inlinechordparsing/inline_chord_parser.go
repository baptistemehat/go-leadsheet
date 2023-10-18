package inlinechordparsing

import (
	"fmt"
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/common/logger"
	"github.com/baptistemehat/go-leadsheet/core/datamodel/music"
	song_pkg "github.com/baptistemehat/go-leadsheet/core/datamodel/song"
	"github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing"
	lexingFunctions "github.com/baptistemehat/go-leadsheet/core/dataprocessing/parsing/inlinechordparsing/lexing/lexingfunctions"
)

type DefaultScheme struct {
}

// InlineChordParser
type InlineChordParser struct {
	Scheme DefaultScheme
}

// Parse
func (p InlineChordParser) Parse(input string) (song_pkg.Song, error) {

	// variables we need to keep while looping
	var tokenValue, propertyKey string

	line := song_pkg.NewLine()
	section := song_pkg.NewSection()
	song := song_pkg.NewSong()
	lexer := lexing.NewLexer(input, lexingFunctions.LexRoot)

	for terminate := false; !terminate; {
		// lex next token
		token := lexer.NextToken()

		logger.Logger.Trace().Msgf("received new token: %s", token.String())

		// trim all non lyrics tokens
		// we need to keep spaces in lyrics
		if token.Type != lexing.TOKEN_LYRICS {
			tokenValue = strings.TrimSpace(token.Value)
		} else {
			tokenValue = token.Value
		}

		switch token.Type {

		case lexing.TOKEN_ERROR:
			logger.Logger.Error().Msgf("error while parsing input: %s", tokenValue)
			terminate = true

		case lexing.TOKEN_EOF:
			// if a section is being parsed
			if len(section.Name) > 0 {

				// stop section parsing, add section to song
				section = section.TrimEmptyLines()
				song.AddSection(section)
			}

			// EOF, terminate parsing
			terminate = true

		case lexing.TOKEN_PROPERTY_KEY:
			propertyKey = tokenValue

		case lexing.TOKEN_PROPERTY_VALUE:
			if err := song.Properties.SetProperty(propertyKey, tokenValue); err != nil {
				return song_pkg.NewSong(), err
			}
			propertyKey = ""

		case lexing.TOKEN_SECTION_NAME:

			// if a section is being parsed
			if len(section.Name) > 0 {

				// stop section parsing, add section to song
				section = section.TrimEmptyLines()
				song.AddSection(section)
			}

			// start parsing new section
			section.Clear()
			section.SetName(tokenValue)

		case lexing.TOKEN_LYRICS:

			// accumulate lyrics to produce lyrics line
			line.AppendLyrics(tokenValue)

		case lexing.TOKEN_NEWLINE:

			// if a line is being parsed
			//if !line.IsEmpty() {

			// stop lyrics parsing, ie add lyrics to section
			section.AddLine(line)
			//}

			// clear line for future parsing
			line.Clear()

		case lexing.TOKEN_CHORD:

			// parse chord
			chord, err := music.ParseChord(tokenValue)
			if err != nil {
				return song_pkg.NewSong(), fmt.Errorf("illegal chord format : %s", tokenValue)
			}

			line.AddChord(chord, uint8(len(line.Lyrics)))

		case lexing.TOKEN_LEFT_PARENTHESIS:
			// TODO : enter "whispered" context

		case lexing.TOKEN_RIGHT_PARENTHESIS:
			// TODO : exit "whispered" context
		}
	}

	return song, nil
}
