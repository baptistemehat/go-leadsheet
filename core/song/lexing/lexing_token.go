package lexing

import (
	"fmt"
	"unicode/utf8"
)

const (
	RUNE_EOF               rune = 0
	RUNE_ERROR             rune = utf8.RuneError
	RUNE_LEFT_PARENTHESIS  rune = '('
	RUNE_RIGHT_PARENTHESIS rune = ')'
	RUNE_LEFT_BRACKET      rune = '['
	RUNE_RIGHT_BRACKET     rune = ']'
	RUNE_LEFT_BRACE        rune = '{'
	RUNE_RIGHT_BRACE       rune = '}'
	RUNE_COLUMN            rune = ':'
	RUNE_NEWLINE           rune = '\n'
)

type TokenType int

const (
	TOKEN_UNKNOWN TokenType = iota
	TOKEN_ERROR
	TOKEN_EOF
	TOKEN_LEFT_PARENTHESIS
	TOKEN_RIGHT_PARENTHESIS
	TOKEN_LEFT_BRACKET
	TOKEN_RIGHT_BRACKET
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_COLUMN
	TOKEN_NEWLINE
	TOKEN_PROPERTY_KEY
	TOKEN_PROPERTY_VALUE
	TOKEN_SECTION_NAME
	TOKEN_LYRICS
	TOKEN_CHORD
)

var tokenTypeToString = map[TokenType]string{
	TOKEN_UNKNOWN:           "UNKNOWN",
	TOKEN_ERROR:             "ERROR",
	TOKEN_EOF:               "EOF",
	TOKEN_LEFT_PARENTHESIS:  "LEFT_PARENTHESIS",
	TOKEN_RIGHT_PARENTHESIS: "RIGHT_PARENTHESIS",
	TOKEN_LEFT_BRACKET:      "LEFT_BRACKET",
	TOKEN_RIGHT_BRACKET:     "RIGHT_BRACKET",
	TOKEN_LEFT_BRACE:        "LEFT_BRACE",
	TOKEN_RIGHT_BRACE:       "RIGHT_BRACE",
	TOKEN_COLUMN:            "COLUMN",
	TOKEN_NEWLINE:           "NEWLINE",
	TOKEN_PROPERTY_KEY:      "PROPERTY_KEY",
	TOKEN_PROPERTY_VALUE:    "PROPERTY_VALUE",
	TOKEN_SECTION_NAME:      "SECTION_NAME",
	TOKEN_LYRICS:            "LYRICS",
	TOKEN_CHORD:             "CHORD",
}

func (tokenType *TokenType) String() string {
	if result, ok := tokenTypeToString[*tokenType]; !ok {
		return "UNKNOWN"
	} else {
		return result
	}
}

type TokenPosition struct {
	Line   int
	Column int
}

func (tokenPosition *TokenPosition) String() string {
	return fmt.Sprintf("[%d:%d]", tokenPosition.Line, tokenPosition.Column)
}

// Token
type Token struct {
	Type  TokenType
	Value string
	Start TokenPosition
	End   TokenPosition
}

// TODO : add line and column to lexer

func NewToken() Token {
	return Token{
		Type:  TOKEN_UNKNOWN,
		Value: "",
		Start: TokenPosition{Line: 0, Column: 0},
		End:   TokenPosition{Line: 0, Column: 0},
	}
}

// String
func (token Token) String() string {
	tokenValue := token.Value

	if token.Type == TOKEN_EOF {
		tokenValue = "EOF"
	}
	return fmt.Sprintf("start%s end%s %s %q ", token.Start.String(), token.End.String(), token.Type.String(), tokenValue)
}
