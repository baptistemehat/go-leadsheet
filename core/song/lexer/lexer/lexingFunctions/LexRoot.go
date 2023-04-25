package lexingFunctions

import (
	"strings"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexRoot(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.SkipWhitespace()

	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACE) {
		return LexLeftBrace
	} else {
		return LexPropertyKey
	}
}
