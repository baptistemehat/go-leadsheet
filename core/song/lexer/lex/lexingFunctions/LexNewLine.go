package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexNewLine
func LexNewLine(lexer *lex.Lexer) lex.LexingFunction {
	lexer.Position += uint(len(lexertoken.NEWLINE))
	lexer.PushToken(lexertoken.TOKEN_NEWLINE)
	return LexSection
}
