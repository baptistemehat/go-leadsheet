package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSectionName
func LexSectionName(lexer *lex.Lexer) lex.LexingFunction {
	for {
		// if newt char is right brace "}" (ie. end of section name)
		if strings.HasPrefix(lexer.Input[lexer.Position:], lexertoken.RIGHT_BRACE) {

			// push section name token
			lexer.PushToken(lexertoken.TOKEN_SECTION_NAME)
			// lex right brace
			return LexRightBrace
		}

		// TODO : exclude non-alphanum char

		// else increment position
		lexer.Inc()

		// if EOF, throw error
		if lexer.IsEOF() {
			return lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_RIGHT_BRACE)
		}
	}
}
