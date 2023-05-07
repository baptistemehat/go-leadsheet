package lex

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
	"github.com/stretchr/testify/assert"
)

func TestNewLexer(t *testing.T) {
	// TODO
	NewLexer("", nil)
}

type moveAfterRuneTestCase struct {
	name          string
	inputRune     rune
	expectedLexer Lexer
}

var moveAfterRunTestCases = []moveAfterRuneTestCase{
	{
		name:      "one byte rune",
		inputRune: 'a',
		expectedLexer: Lexer{
			Input:  "abcde",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				// TODO : try to use a real LexingFunction
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 1},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  1,
			},
		}},
	{
		name:      "two byte rune",
		inputRune: 'ä',
		expectedLexer: Lexer{
			Input:  "äbcde",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 1},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  2,
			},
		}},
	{
		name:      "EOF rune",
		inputRune: lexertoken.EOF,
		expectedLexer: Lexer{
			Input:  "",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 0},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  0,
			},
		}},
	{
		name:      "ERROR rune",
		inputRune: lexertoken.ERROR,
		expectedLexer: Lexer{
			Input:  "",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 0},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  0,
			},
		}},
}

func TestMoveAfterRune(t *testing.T) {
	for _, testCase := range moveAfterRunTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := NewLexer(testCase.expectedLexer.Input, testCase.expectedLexer.status.NextLexingFunction)

			lexer.MoveAfterRune(testCase.inputRune)

			assert.Equal(t, testCase.expectedLexer.status, lexer.status)
		})
	}
}

type peekRunTestCase struct {
	name          string
	expectedRune  rune
	expectedLexer Lexer
}

var peekRuneTestCases = []peekRunTestCase{
	{
		name:         "one byte rune",
		expectedRune: 'c',
		expectedLexer: Lexer{
			Input:  "abcde",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				// TODO : try to use a real LexingFunction
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "ab",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 2},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  2,
			},
		}},
	{
		name:         "two byte rune",
		expectedRune: 'ä',
		expectedLexer: Lexer{
			Input:  "äbcde",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				// TODO : try to use a real LexingFunction
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 0},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  0,
			},
		}},
	{
		name:         "empty input",
		expectedRune: lexertoken.EOF,
		expectedLexer: Lexer{
			Input:  "",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				// TODO : try to use a real LexingFunction
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 0},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  0,
			},
		}},
	{
		name:         "EOF",
		expectedRune: lexertoken.EOF,
		expectedLexer: Lexer{
			Input:  "abcde",
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				// TODO : try to use a real LexingFunction
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "ab",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 5},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  5,
			},
		}},
	{
		name:         "RuneError",
		expectedRune: lexertoken.ERROR,
		expectedLexer: Lexer{
			Input:  string(utf8.RuneError),
			Tokens: make(chan lexertoken.Token, 5),
			status: LexerStatus{
				// TODO : try to use a real LexingFunction
				NextLexingFunction: nil,
				CurrentToken: lexertoken.Token{
					Type:  lexertoken.TOKEN_UNKNOWN,
					Value: "",
					Start: lexertoken.TokenPosition{Line: 0, Column: 0},
					End:   lexertoken.TokenPosition{Line: 0, Column: 0},
				},
				CurrentTokenStart: 0,
				PositionInBuffer:  0,
			},
		}},
}

func TestPeekRune(t *testing.T) {
	for _, testCase := range peekRuneTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := testCase.expectedLexer

			actualRune := lexer.PeekRune()

			assert.Equal(t, testCase.expectedRune, actualRune, "should be equal")
		})
	}
}

func TestNextToken(t *testing.T) {
	// TODO
}

func TestErrorf(t *testing.T) {
	// TODO
}

func TestSkipWitheSpace(t *testing.T) {
	// TODO
}

func TestNewLine(t *testing.T) {
	// TODO
}
