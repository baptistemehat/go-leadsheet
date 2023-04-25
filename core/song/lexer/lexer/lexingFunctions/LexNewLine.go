package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexNewLine(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.NEWLINE)
	lexer.PushToken(lexertoken.TOKEN_NEWLINE)
	return LexSection
}
