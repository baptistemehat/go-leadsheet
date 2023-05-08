package lexererrors

const (
	LEXER_ERROR_UNEXPECTED_EOF                    string = "Unexpected end of file"
	LEXER_ERROR_MISSING_RIGHT_BRACE               string = "Missing closing section brace"
	LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE string = "Missing newline before start of section"
	LEXER_ERROR_MISSING_NEWLINE_BEFORE_EOF        string = "Missing newline before end of file"
	LEXER_ERROR_MISSING_RIGHT_PARENTHESIS         string = "Missing closing parenthesis"
	LEXER_ERROR_START_OF_TOKEN_AFTER_EOF          string = "Start of current token is after EOF"
	LEXER_ERROR_POSITION_AFTER_EOF                string = "Lexer position is after EOF"
	LEXER_ERROR_NIL_LEXING_FUNCTION               string = "Cannot call nil lexing function"
)
