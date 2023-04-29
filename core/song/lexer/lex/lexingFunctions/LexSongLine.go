package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSongLine
func LexSongLine(lexer *lex.Lexer) lex.LexingFunction {
	for {

		// if next char is new line "\n" (ie. end of lyrics line)
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.NEWLINE) {

			// push lyrics token
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			// lex new line
			return LexNewLine

			// else if next char is left bracket "[" (ie. chord)
		} else if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.LEFT_BRACKET) {

			// push lyrics token
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			// lex left bracket
			return LexLeftBracket

			// else if next char is left parenthesis "(" (ie. whispered lyrics)
		} else if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.LEFT_PARENTHESIS) {

			// push lyrics token
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			// lex left parenthesis
			return LexLeftParenthesis

			// else if next char is left brace "{" (ie. start of new section)
		} else if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.LEFT_BRACE) {

			// throw error
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE)
		}

		// else increment position
		lexer.Inc()

		// no bracket or brace or parenthesis

		// if EOF, throw error
		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_NEWLINE_BEFORE_EOF)
		}

	}
}
