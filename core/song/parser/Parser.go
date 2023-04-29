package parser

import (
	"fmt"
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex/lexingFunctions"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
	"github.com/baptistemehat/go-leadsheet/core/song/model"
)

type DefaultScheme struct {
}

// Parser
type Parser interface {
	Parse(string) (model.Song, error)
}

// InlineChordParser
type InlineChordParser struct {
	Scheme DefaultScheme
}

// Parse
func (p InlineChordParser) Parse(input string) (model.Song, error) {

	// variables we need to keep while looping
	var tokenValue, propertyKey string

	line := model.NewLine()
	section := model.NewSection()
	song := model.NewSong()
	lexer := lex.NewLexer(input, lexingFunctions.LexRoot)

	for terminate := false; !terminate; {
		// lex next token
		token := lexer.NextToken()

		// trim all non lyrics tokens
		// we need to keep spaces in lyrics
		if token.Type != lexertoken.TOKEN_LYRICS {
			tokenValue = strings.TrimSpace(token.Value)
		} else {
			tokenValue = token.Value
		}

		switch token.Type {

		case lexertoken.TOKEN_EOF:

			// if a section is being parsed
			if len(section.Name) > 0 {

				// stop section parsing, add section to song
				song.AddSection(section)
			}

			// EOF, terminate parsing
			terminate = true

		case lexertoken.TOKEN_PROPERTY_KEY:
			propertyKey = tokenValue

		case lexertoken.TOKEN_PROPERTY_VALUE:
			song.Properties.SetProperty(propertyKey, tokenValue)
			propertyKey = ""

		case lexertoken.TOKEN_SECTION_NAME:

			// if a section is being parsed
			if len(section.Name) > 0 {

				// stop section parsing, add section to song
				song.AddSection(section)
			}

			// start parsing new section
			section.Clear()
			section.SetName(tokenValue)

		case lexertoken.TOKEN_LYRICS:

			// accumulate lyrics to produce lyrics line
			line.AppendLyrics(tokenValue)

		case lexertoken.TOKEN_NEWLINE:

			// if a line is being parsed
			if !line.IsEmpty() {

				// stop lyrics parsing, ie add lyrics to section
				section.AddLine(line)
			}

			// clear line for future parsing
			line.Clear()

		case lexertoken.TOKEN_CHORD:

			// parse chord
			chord, err := model.ParseChord(tokenValue)
			if err != nil {
				return model.Song{}, fmt.Errorf("illegal chord format : %s", tokenValue)
			}

			line.AddChord(chord, uint8(len(line.Lyrics)))

		case lexertoken.TOKEN_LEFT_PARENTHESIS:
			// TODO : enter "whispered" context

		case lexertoken.TOKEN_RIGHT_PARENTHESIS:
			// TODO : exit "whispered" context
		}

	}

	return song, nil
}
