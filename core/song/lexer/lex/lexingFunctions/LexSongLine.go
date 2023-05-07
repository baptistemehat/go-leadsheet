package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lex"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
)

// LexSongLine
func LexSongLine(lexer *lex.Lexer) lex.LexingFunction {
	for {

		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexertoken.EOF:
			lexer.PushToken(lexertoken.TOKEN_EOF)
			return nil

		case lexertoken.ERROR:
			lexer.Errorf("unexpected character found")
			return nil

		case lexertoken.NEWLINE:
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexNewLine

		case lexertoken.LEFT_BRACKET:
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexLeftBracket

		case lexertoken.LEFT_PARENTHESIS:
			lexer.PushToken(lexertoken.TOKEN_LYRICS)
			return LexLeftParenthesis

		case lexertoken.LEFT_BRACE:
			lexer.Errorf(lexererrors.LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE)
			return nil
		}

		lexer.GoToNextRune(nextRune)
	}
}
