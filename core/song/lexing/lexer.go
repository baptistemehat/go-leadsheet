package lexing

import (
	"fmt"
	"unicode"
	"unicode/utf8"

)

type Lexer struct {
	input              string
	tokens             chan Token
	currentToken       Token
	nextLexingFunction LexingFunction
	position           int
	currentTokenStart  int
}

// NewLexer creates a new lexer
func NewLexer(input string, lexingFunc LexingFunction) *Lexer {
	return &Lexer{
		input:              input,
		tokens:             make(chan Token, 5),
		currentToken:       NewToken(),
		nextLexingFunction: lexingFunc,
		position:           0,
		currentTokenStart:  0,
	}
}

// TODO : rename ConsumeRune /

// MoveAfterRune moves position just after next rune
func (lexer *Lexer) MoveAfterRune(nextRune rune) {
	switch nextRune {
	// TODO : rename to be rune
	case RUNE_EOF:
		return
	case RUNE_ERROR:
		return
	default:
		lexer.position += utf8.RuneLen(nextRune)
		lexer.currentToken.End.Column++
		if nextRune == RUNE_NEWLINE {
			lexer.countNewline()
		}
	}
}

// PeekRune returns the next rune in input.
// Returns EOF if OEF is reached.
// Returns ERROR if error occured while reading next rune.
func (lexer *Lexer) PeekRune() rune {

	// if position reached last rune of input
	if lexer.position >= len(lexer.input) {
		return RUNE_EOF
	}

	// get next rune in input
	nextRune, _ := utf8.DecodeRuneInString(lexer.input[lexer.position:])
	if nextRune == utf8.RuneError {
		return RUNE_ERROR
	}
	return nextRune
}

// PushToken pushes a token into the token channel
func (lexer *Lexer) PushToken(tokenType TokenType) {

	if lexer.currentTokenStart > len(lexer.input) {
		lexer.Errorf(LEXER_ERROR_START_OF_TOKEN_AFTER_EOF)
		return
	}

	if lexer.position > len(lexer.input) {
		lexer.Errorf(LEXER_ERROR_POSITION_AFTER_EOF)
		return
	}

	// prepare token and push it
	lexer.currentToken.Type = tokenType
	lexer.currentToken.Value = lexer.input[lexer.currentTokenStart:lexer.position]
	lexer.tokens <- lexer.currentToken

	// reset token
	lexer.currentToken.Type = TOKEN_UNKNOWN
	lexer.currentToken.Value = ""
	lexer.currentToken.Start = lexer.currentToken.End

	// update positions
	lexer.currentTokenStart = lexer.position
	if tokenType == TOKEN_NEWLINE {
		lexer.countNewline()
	}
}

// NextToken procedes lexing until a token is produced and returns it
func (lexer *Lexer) NextToken() Token {
	for {
		select {
		case token := <-lexer.tokens:
			return token
		default:
			if lexer.nextLexingFunction == nil {

				lexer.currentToken.Type = TOKEN_ERROR
				lexer.currentToken.Value = fmt.Sprint(LEXER_ERROR_NIL_LEXING_FUNCTION)

				return lexer.currentToken
			}

			lexer.nextLexingFunction = lexer.nextLexingFunction(lexer)
		}
	}
}

// Errorf pushes an error token in the token channel.
// Errorf is a lexing function.
func (lexer *Lexer) Errorf(format string, args ...interface{}) LexingFunction {

	lexer.currentToken.Type = TOKEN_ERROR
	lexer.currentToken.Value = fmt.Sprintf(format, args...)

	lexer.tokens <- lexer.currentToken

	return nil
}

// SkipWhitespace moves position to the next non-whitespaces rune.
// Non-whitespace runes are all runes for which utf8.IsSpace returns false.
// If EOF is reached, pushes an EOF 
func (lexer *Lexer) SkipWhitespace() {
	for {

		nextRune := lexer.PeekRune()

		if nextRune == RUNE_EOF {
			lexer.PushToken(TOKEN_EOF)
			break
		}

		if !unicode.IsSpace(nextRune) {
			break
		}

		lexer.MoveAfterRune(nextRune)
	}

	lexer.currentTokenStart = lexer.position
}

// countNewline
func (lexer *Lexer) countNewline() {
	lexer.currentToken.Start.Line++
	lexer.currentToken.Start.Column = 0
	lexer.currentToken.End.Line++
	lexer.currentToken.End.Column = 0
}
