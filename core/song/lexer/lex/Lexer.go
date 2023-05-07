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

	// CurrentTokenStart
	currentToken      lexertoken.Token
	currentTokenStart int
	// positionInBuffer is index of the lexer within the Input string.
	// Since it is used to index a string, positionInBuffer counts in bytes, not runes
	positionInBuffer int
}

// NewLexer creates a new lexer
func NewLexer(input string, lexingFunc LexingFunction) *Lexer {
	return &Lexer{
		Input:              input,
		Tokens:             make(chan lexertoken.Token, 5),
		NextLexingFunction: lexingFunc,
		// TODO : evaluate the size needed for Tokens
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_ERROR,
			Value: "",
			Start: lexertoken.TokenPosition{0, 0},
			End:   lexertoken.TokenPosition{0, 0},
		},
		positionInBuffer: 0,
	}
}

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
		lexer.positionInBuffer += utf8.RuneLen(nextRune)
		lexer.currentToken.End.Column++
	}
}

// PeekRune returns the next rune in input.
// Returns EOF if OEF is reached.
// Returns ERROR if error occured while reading next rune.
func (lexer *Lexer) PeekRune() rune {
	// if position reached last rune of input
	if lexer.positionInBuffer >= len(lexer.Input) {
		return lexertoken.EOF
	}

	// get next rune in input
	nextRune, _ := utf8.DecodeRuneInString(lexer.Input[lexer.positionInBuffer:])
	if nextRune == utf8.RuneError {
		return lexertoken.ERROR
	}
	return nextRune
}

// PushToken pushes a token into the token channel
func (lexer *Lexer) PushToken(tokenType lexertoken.TokenType) {

	if lexer.currentTokenStart > len(lexer.Input) {
		lexer.Errorf("lexer.Start exceeds len(lexer.Input)")
		return
	}

	if lexer.positionInBuffer > len(lexer.Input) {
		lexer.Errorf("lexer.Position exceeds len(lexer.Input)")
		return
	}

	lexer.currentToken.Type = tokenType
	lexer.currentToken.Value = lexer.Input[lexer.currentTokenStart:lexer.positionInBuffer]

	lexer.Tokens <- lexer.currentToken

	lexer.currentToken.Start = lexer.currentToken.End

	lexer.currentTokenStart = lexer.positionInBuffer
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
			// TODO : handle nil return case
			lexer.NextLexingFunction = lexer.NextLexingFunction(lexer)
		}
	}
}

// Errorf
func (lexer *Lexer) Errorf(format string, args ...interface{}) LexingFunction {

	lexer.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...), /* + fmt.Sprintf(" at [%d:%d]", lexer.lineCount, lexer.columnCount)*/
	}

	// TODO : add line:column indication

	return nil
}

// SkipWhitespace
func (lexer *Lexer) SkipWhitespace() {
	for {

		nextRune := lexer.PeekRune()

		if nextRune == lexertoken.NEWLINE {
			lexer.Newline()
		}

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
	lexer.currentTokenStart = lexer.positionInBuffer
}

// Newline
func (lexer *Lexer) Newline() {
	lexer.currentToken.Start.Line++
	lexer.currentToken.Start.Column = 0
	lexer.currentToken.End.Line++
	lexer.currentToken.End.Column = 0
}
