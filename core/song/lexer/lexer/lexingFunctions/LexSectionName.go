package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexSectionName(lexer *lexer.Lexer) lexer.LexingFunction {
	for {
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_BRACE) {
			lexer.PushToken(lexertoken.TOKEN_SECTION_NAME)
			return LexRightBrace
		}

		lexer.Inc()

		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_BRACE)
		}
	}
}
