package parser

import (
	"fmt"
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer/lexingFunctions"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
	"github.com/baptistemehat/go-leadsheet/core/song/model/song"
)

type DefaultScheme struct {
}

type Parser struct {
	Scheme DefaultScheme
}

func (p *Parser) Parse(input string) (song.Song, error) {
	s := song.NewSong()

	var token lexertoken.Token
	var tokenValue string

	l := lexer.NewLexer(input, lexingFunctions.LexRoot)

	section := song.NewSection()
	line := song.NewLine()
	propertyKey := ""
	lyrics := ""

	var ch song.Chord
	var error error

	for {
		token = l.NextToken()

		if token.Type != lexertoken.TOKEN_LYRICS {
			tokenValue = strings.TrimSpace(token.Value)
		} else {
			tokenValue = token.Value
		}

		if token.IsEOF() {
			// if a section is being parsed
			if len(section.Name) > 0 {
				s.AddSection(section)
			}
			break
		}

		switch token.Type {
		case lexertoken.TOKEN_PROPERTY_KEY:

			propertyKey = tokenValue

		case lexertoken.TOKEN_PROPERTY_VALUE:
			s.Properties.SetProperty(propertyKey, tokenValue)
			propertyKey = ""

		case lexertoken.TOKEN_SECTION_NAME:

			// end last section and start new

			if len(section.Name) > 0 {
				s.AddSection(section)
				section.Clear()
			}

			section.SetName(tokenValue)

		case lexertoken.TOKEN_LYRICS:
			lyrics += tokenValue

		case lexertoken.TOKEN_NEWLINE:

			// end last line and start new line
			if len(lyrics) > 0 || len(line.Chords) > 0 {
				line.SetLyrics(lyrics)
				section.AddLine(line)
			}

			lyrics = ""
			line.Clear()

		case lexertoken.TOKEN_CHORD:

			if ch, error = song.ParseChord(tokenValue); error != nil {
				return song.Song{}, fmt.Errorf("illegal chord format : %s", tokenValue)
			}

			line.AddChord(ch, uint8(len(lyrics)))

		case lexertoken.TOKEN_LEFT_PARENTHESIS:
			// enter "whispered" context
		case lexertoken.TOKEN_RIGHT_PARENTHESIS:
			// exit "whispered" context
		}

	}

	return s, nil
}
