package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexColumn(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.COLUMN)
	lexer.PushToken(lexertoken.TOKEN_COLUMN)
	return LexPropertyValue
}
