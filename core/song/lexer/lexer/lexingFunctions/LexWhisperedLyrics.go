package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexWhisperedLyrics(lexer *lexer.Lexer) lexer.LexingFunction {
	for {
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_PARENTHESIS) {
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexRightParenthesis
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_PARENTHESIS)
		}
	}
}
