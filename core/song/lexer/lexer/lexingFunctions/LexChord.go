package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexChord(lexer *lexer.Lexer) lexer.LexingFunction {
	for {
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_BRACKET) {
			lexer.PushToken(lexertoken.TOKEN_CHORD)
			return LexRightBracket
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_BRACE)
		}
	}
}
