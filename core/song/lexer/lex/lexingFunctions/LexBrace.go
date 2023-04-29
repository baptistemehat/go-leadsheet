package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexLeftBrace
func LexLeftBrace(lexer *lex.Lexer) lex.LexingFunction {
	lexer.Position += len(lexertoken.LEFT_BRACE)
	lexer.PushToken(lexertoken.TOKEN_LEFT_BRACE)
	return LexSectionName
}

// LexRightBrace
func LexRightBrace(lexer *lex.Lexer) lex.LexingFunction {
	lexer.Position += len(lexertoken.RIGHT_BRACE)
	lexer.PushToken(lexertoken.TOKEN_RIGHT_BRACE)
	return LexSection
}
