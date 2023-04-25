package lexererrors

const (
	LEXER_ERROR_UNEXPECTED_EOF                    string = "Unexpected end of file"
	LEXER_ERROR_MISSING_RIGHT_BRACE               string = "Missing closing section brace"
	LEXER_ERROR_MISSING_NEWLINE_BEFORE_LEFT_BRACE string = "Missing newline before start of section"
	LEXER_ERROR_MISSING_NEWLINE_BEFORE_EOF        string = "Missing newline before end of file"
	LEXER_ERROR_MISSING_RIGHT_PARENTHESIS         string = "Missing closing parenthesis"
)
