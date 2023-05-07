package lex

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

type LexerStatus struct {
	NextLexingFunction LexingFunction
	CurrentToken       lexertoken.Token
	CurrentTokenStart  int
	// positionInBuffer is index of the lexer within the Input string.
	// Since it is used to index a string, positionInBuffer counts in bytes, not runes
	PositionInBuffer int
}

type Lexer struct {
	Input  string
	Tokens chan lexertoken.Token
	status LexerStatus
}

// NewLexer creates a new lexer
func NewLexer(input string, lexingFunc LexingFunction) *Lexer {
	return &Lexer{
		Input:  input,
		Tokens: make(chan lexertoken.Token, 5),
		status: LexerStatus{
			NextLexingFunction: lexingFunc,
			CurrentToken:       lexertoken.NewToken(),
			CurrentTokenStart:  0,
			PositionInBuffer:   0,
		},
	}
}
func (lexer *Lexer) GetPositionInBuffer() int {
	return lexer.status.PositionInBuffer
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
		lexer.status.PositionInBuffer += utf8.RuneLen(nextRune)
		lexer.status.CurrentToken.End.Column++
	}
}

// PeekRune returns the next rune in input.
// Returns EOF if OEF is reached.
// Returns ERROR if error occured while reading next rune.
func (lexer *Lexer) PeekRune() rune {
	// if position reached last rune of input
	if lexer.status.PositionInBuffer >= len(lexer.Input) {
		return lexertoken.EOF
	}

	// get next rune in input
	nextRune, _ := utf8.DecodeRuneInString(lexer.Input[lexer.status.PositionInBuffer:])
	if nextRune == utf8.RuneError {
		return lexertoken.ERROR
	}
	return nextRune
}

// PushToken pushes a token into the token channel
func (lexer *Lexer) PushToken(tokenType lexertoken.TokenType) {

	if lexer.status.CurrentTokenStart > len(lexer.Input) {
		lexer.Errorf("lexer.Start exceeds len(lexer.Input)")
		return
	}

	if lexer.status.PositionInBuffer > len(lexer.Input) {
		lexer.Errorf("lexer.Position exceeds len(lexer.Input)")
		return
	}

	lexer.status.CurrentToken.Type = tokenType
	lexer.status.CurrentToken.Value = lexer.Input[lexer.status.CurrentTokenStart:lexer.status.PositionInBuffer]

	lexer.Tokens <- lexer.status.CurrentToken

	lexer.status.CurrentToken.Start = lexer.status.CurrentToken.End

	lexer.status.CurrentTokenStart = lexer.status.PositionInBuffer
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
			lexer.status.NextLexingFunction = lexer.status.NextLexingFunction(lexer)
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

		lexer.MoveAfterRune(nextRune)
	}
	lexer.status.CurrentTokenStart = lexer.status.PositionInBuffer
}

// Newline
func (lexer *Lexer) Newline() {
	lexer.status.CurrentToken.Start.Line++
	lexer.status.CurrentToken.Start.Column = 0
	lexer.status.CurrentToken.End.Line++
	lexer.status.CurrentToken.End.Column = 0
}
