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

	Start         int
	NextRuneWidth int
	// Position is index of the lexer within the Input string.
	// Since it is used to index a string, Position counts in bytes, not runes
	Position int
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

// Inc increments lexer position.
// If EOF is reached, pushes EOF token.
// func (lexer *Lexer) Inc() {

// 	lexer.Position++

// 	// if position reached last rune of input
// 	if lexer.Position >= len(lexer.Input) {
// 		lexer.PushToken(lexertoken.TOKEN_EOF)
// 	}
// }

// // Dec decrements lexer position.
// // Position is capped at 0.
// func (lexer *Lexer) Dec() {
// 	if lexer.Position != 0 {
// 		lexer.Position--
// 	}
// }

// TODO : rename ConsumeRune /

// GoToNextRune moves position just after next rune
func (lexer *Lexer) GoToNextRune(nextRune rune) {
	switch nextRune {
	// TODO : rename to be rune
	case lexertoken.EOF:
		return
	case lexertoken.ERROR:
		return
	default:
		lexer.Position += utf8.RuneLen(nextRune)
	}
}

// PeekRune returns the next rune in input, and updates NextRuneWidth.
// Returns EOF if OEF is reached.
// Returns ERROR if error occured while reading next rune.
func (lexer *Lexer) PeekRune() rune {
	// if position reached last rune of input
	if lexer.Position >= len(lexer.Input) {
		return lexertoken.EOF
	}

	// get next rune in input
	nextRune, width := utf8.DecodeRuneInString(lexer.Input[lexer.Position:])
	if nextRune == utf8.RuneError {
		return lexertoken.ERROR
	}
	lexer.NextRuneWidth = width

	return nextRune
}

// PushToken pushes a token into the token channel
func (lexer *Lexer) PushToken(tokenType lexertoken.TokenType) {

	if lexer.Start > len(lexer.Input) {
		lexer.Errorf("lexer.Start exceeds len(lexer.Input)")
		return
	}

	if lexer.Position > len(lexer.Input) {
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

	// TODO : add line:column indication

	return nil
}

// IsEOF
func (lexer *Lexer) IsEOF() bool {
	return lexer.Position >= len(lexer.Input)
}

// SkipWhitespace
func (lexer *Lexer) SkipWhitespace() {
	for {

		nextRune := lexer.PeekRune()

		if !unicode.IsSpace(nextRune) {
			break
		}

		// and here we only check lexer.IsEOF ?
		if nextRune == lexertoken.EOF {
			lexer.PushToken(lexertoken.TOKEN_EOF)
			break
		}

		lexer.GoToNextRune(nextRune)
	}
}
