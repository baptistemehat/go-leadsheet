package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexChord
func LexChord(lexer *lex.Lexer) lex.LexingFunction {
	for {
		// if next char is right bracket "]" (ie. end of chord)
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_BRACKET) {

			// push chord token
			lexer.PushToken(lexertoken.TOKEN_CHORD)

			// lex right bracket
			return LexRightBracket
		}

		// retur error if "non-chord" char

		// else increment position
		lexer.Inc()

		// if EOF, throw error
		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_BRACE)
		}
	}
}
