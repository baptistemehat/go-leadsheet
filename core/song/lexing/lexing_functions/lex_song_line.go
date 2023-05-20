package lexingFunctions

import (
	"github.com/baptistemehat/go-leadsheet/core/song/lexing"
)

// LexSongLine
func LexSongLine(lexer *lexing.Lexer) lexing.LexingFunction {
	for {

		nextRune := lexer.PeekRune()

		switch nextRune {

		case lexing.RUNE_EOF:
			lexer.PushToken(lexing.TOKEN_LYRICS)
			lexer.PushToken(lexing.TOKEN_EOF)
			return nil

		case lexing.RUNE_ERROR:
			lexer.Errorf(lexing.LEXER_ERROR_UNEXPECTED_RUNE)
			return nil

		case lexing.RUNE_NEWLINE:
			lexer.PushToken(lexing.TOKEN_LYRICS)
			return LexNewLine

		case lexing.RUNE_LEFT_BRACKET:
			lexer.PushToken(lexing.TOKEN_LYRICS)
			return LexLeftBracket

		case lexing.RUNE_LEFT_PARENTHESIS:
			lexer.PushToken(lexing.TOKEN_LYRICS)
			return LexLeftParenthesis

		case lexing.RUNE_LEFT_BRACE:
			lexer.Errorf(lexing.LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE)
			return nil
		}

		lexer.MoveAfterRune(nextRune)
	}
}
