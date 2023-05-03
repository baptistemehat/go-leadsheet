package lex

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

type Lexer struct {
	Input              string
	Tokens             chan lexertoken.Token
	NextLexingFunction LexingFunction

	Start            uint
	Position         uint
	CurrentRuneWidth uint
}

// NewLexer creates a new lexer
func NewLexer(input string, lexingFunc LexingFunction) *Lexer {
	l := &Lexer{
		Input:              input,
		NextLexingFunction: lexingFunc,
		Tokens:             make(chan lexertoken.Token, 5),
		// TODO : evaluate the size needed for Tokens
	}

	return l
}

// Inc increments lexer position
// if EOF is reached, pushes EOF token
func (lexer *Lexer) Inc() {

	lexer.Position++

	// if position reached last rune of input
	if lexer.Position >= (uint)(utf8.RuneCountInString(lexer.Input)) {
		lexer.PushToken(lexertoken.TOKEN_EOF)
	}
}

// Dec decrements lexer position
func (lexer *Lexer) Dec() {
	// decrement position
	lexer.Position--
}

// NextRune moves position to next rune in input and returns it
func (lexer *Lexer) NextRune() rune {

	// if position reached last rune of input
	if lexer.Position >= (uint)(utf8.RuneCountInString(lexer.Input)) {
		lexer.CurrentRuneWidth = 0
		return lexertoken.EOF
	}

	// get next rune in input
	nextRune, width := utf8.DecodeRuneInString(lexer.Input[lexer.Position:])
	lexer.CurrentRuneWidth = uint(width)
	lexer.Position += uint(width)

	return nextRune
}

// PushToken pushes a token into the token channel
func (lexer *Lexer) PushToken(tokenType lexertoken.TokenType) {

	if lexer.Start > uint(len(lexer.Input)) {
		lexer.Errorf("lexer.Start exceeds len(lexer.Input)")
		return
	}

	if lexer.Position > uint(len(lexer.Input)) {
		lexer.Errorf("lexer.Position exceeds len(lexer.Input)")
		return
	}

	lexer.Tokens <- lexertoken.Token{
		Type:  tokenType,
		Value: lexer.Input[lexer.Start:lexer.Position],
	}
	lexer.Start = lexer.Position
}

// NextToken procedes lexing until a token is produced and returns it
func (lexer *Lexer) NextToken() lexertoken.Token {
	for {
		select {
		// try to pull token from channel
		case token := <-lexer.Tokens:
			return token
		// if no token to pull, resume lexing
		default:
			lexer.NextLexingFunction = lexer.NextLexingFunction(lexer)
		}
	}
}

// Errorf
func (lexer *Lexer) Errorf(format string, args ...interface{}) LexingFunction {
	lexer.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

// IsEOF
func (lexer *Lexer) IsEOF() bool {
	return lexer.Position >= uint(len(lexer.Input))
}

// SkipWhitespace
func (lexer *Lexer) SkipWhitespace() {
	for {
		r := lexer.NextRune()

		if !unicode.IsSpace(r) {
			lexer.Dec()
			break
		}

		if r == lexertoken.EOF {
			lexer.PushToken(lexertoken.TOKEN_EOF)
			return
		}
	}
}
