package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexColumn
func LexColumn(lexer *lex.Lexer) lex.LexingFunction {
	lexer.Position += len(lexertoken.COLUMN)
	lexer.PushToken(lexertoken.TOKEN_COLUMN)
	return LexPropertyValue
}
