package lex

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

type Lexer struct {
	input              string
	tokens             chan lexertoken.Token
	nextLexingFunction LexingFunction
	currentTokenStart  int
	currentToken       lexertoken.Token
	// positionInBuffer is index of the lexer within the Input string.
	// Since it is used to index a string, positionInBuffer counts in bytes, not runes
	positionInBuffer int
}

// NewLexer creates a new lexer
func NewLexer(input string, lexingFunc LexingFunction) *Lexer {
	return &Lexer{
		input:              input,
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: lexingFunc,
		currentToken:       lexertoken.NewToken(),
		currentTokenStart:  0,
		positionInBuffer:   0,
	}
}
func (lexer *Lexer) GetPositionInBuffer() int {
	return lexer.positionInBuffer
}

// TODO : rename ConsumeRune /

// MoveAfterRune moves position just after next rune
func (lexer *Lexer) MoveAfterRune(nextRune rune) {
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
	if lexer.positionInBuffer >= len(lexer.input) {
		return lexertoken.EOF
	}

	// get next rune in input
	nextRune, _ := utf8.DecodeRuneInString(lexer.input[lexer.positionInBuffer:])
	if nextRune == utf8.RuneError {
		return lexertoken.ERROR
	}
	return nextRune
}

// PushToken pushes a token into the token channel
func (lexer *Lexer) PushToken(tokenType lexertoken.TokenType) {

	if lexer.currentTokenStart > len(lexer.input) {
		lexer.Errorf(lexererrors.LEXER_ERROR_START_OF_TOKEN_AFTER_EOF)
		return
	}

	if lexer.positionInBuffer > len(lexer.input) {
		lexer.Errorf(lexererrors.LEXER_ERROR_POSITION_AFTER_EOF)
		return
	}

	lexer.currentToken.Type = tokenType
	lexer.currentToken.Value = lexer.input[lexer.currentTokenStart:lexer.positionInBuffer]
	lexer.tokens <- lexer.currentToken

	lexer.currentToken.Type = lexertoken.TOKEN_UNKNOWN
	lexer.currentToken.Value = ""
	lexer.currentToken.Start = lexer.currentToken.End

	lexer.currentTokenStart = lexer.positionInBuffer
}

// NextToken procedes lexing until a token is produced and returns it
func (lexer *Lexer) NextToken() lexertoken.Token {
	for {
		select {
		// try to pull token from channel
		case token := <-lexer.tokens:
			return token
		// if no token to pull, resume lexing
		default:
			if lexer.nextLexingFunction == nil {

				lexer.currentToken.Type = lexertoken.TOKEN_ERROR
				lexer.currentToken.Value = fmt.Sprint(lexererrors.LEXER_ERROR_NIL_LEXING_FUNCTION)

				return lexer.currentToken
			}

			lexer.nextLexingFunction = lexer.nextLexingFunction(lexer)
		}
	}
}

// Errorf
func (lexer *Lexer) Errorf(format string, args ...interface{}) LexingFunction {

	lexer.currentToken.Type = lexertoken.TOKEN_ERROR
	lexer.currentToken.Value = fmt.Sprintf(format, args...)

	lexer.tokens <- lexer.currentToken

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

		lexer.MoveAfterRune(nextRune)
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
