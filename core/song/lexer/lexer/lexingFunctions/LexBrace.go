package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexer"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

func LexLeftBrace(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.LEFT_BRACE)
	lexer.PushToken(lexertoken.TOKEN_LEFT_BRACE)
	return LexSectionName
}

func LexRightBrace(lexer *lexer.Lexer) lexer.LexingFunction {
	lexer.Position += len(lexertoken.RIGHT_BRACE)
	lexer.PushToken(lexertoken.TOKEN_RIGHT_BRACE)
	return LexSection
}
