package lex

import (
	"testing"
	"unicode/utf8"

	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexererrors"
	"github.com/baptistemehat/go-leadsheet/core/song/lexer/lexertoken"
	"github.com/stretchr/testify/assert"
)

// testTokenType is a test token type
const testTokenType lexertoken.TokenType = -1

// testToken is a test token
var testToken = lexertoken.Token{
	Type:  testTokenType,
	Value: "test",
	Start: lexertoken.TokenPosition{Line: 0, Column: 0},
	End:   lexertoken.TokenPosition{Line: 0, Column: 12},
}

// MockLexingFuncCallCounter
type MockLexingFuncCallCounter struct {
	callCount int
}

// ResetCount resets the call counter
func (mockLexingFuncCallCount *MockLexingFuncCallCounter) ResetCount() {
	mockLexingFuncCallCount.callCount = 0
}

// GetCount returns the call counter
func (mockLexingFuncCallCount *MockLexingFuncCallCounter) GetCount() int {
	return mockLexingFuncCallCount.callCount
}

// Increment increments the call counter
func (mockLexingFuncCallCounter *MockLexingFuncCallCounter) Increment() {
	mockLexingFuncCallCounter.callCount++
}

// mockLexingFuncCallCounter is the function call counter associated to
var mockLexingFuncCallCounter MockLexingFuncCallCounter

// mockLexingFunction mocks a lexing function.
// Increments the MockLexingFuncCallCounter and pushes a testToken in the Tokens channel.
func mockLexingFunction(lexer *Lexer) LexingFunction {
	lexer.tokens <- testToken
	mockLexingFuncCallCounter.Increment()
	return mockLexingFunction
}

// assertEqualLexer assert if two Lexers are equal.
// This function only asserts equality for fields implementing '=='.
// NextLexingFunction function and Tokens channel are not considred in this function.
func assertEqualLexer(t *testing.T, lexer, otherLexer *Lexer) {
	assert.Equal(t, lexer.input, otherLexer.input)
	assert.Equal(t, lexer.currentToken, otherLexer.currentToken)
	assert.Equal(t, lexer.currentTokenStart, otherLexer.currentTokenStart)
	assert.Equal(t, lexer.positionInBuffer, otherLexer.positionInBuffer)
}

func TestNewLexer(t *testing.T) {
	// TODO
	NewLexer("", nil)
}

type moveAfterRuneTestCase struct {
	name          string
	inputRune     rune
	inputLexer    *Lexer
	expectedLexer *Lexer
}

var moveAfterRunTestCases = []moveAfterRuneTestCase{
	{
		name:      "one byte rune",
		inputRune: 'a',
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 2},
				End:   lexertoken.TokenPosition{Line: 0, Column: 4},
			},
			currentTokenStart: 2,
			positionInBuffer:  4,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 2},
				End:   lexertoken.TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 2,
			positionInBuffer:  5,
		}},
	{
		name:      "two byte rune",
		inputRune: 'ä',
		inputLexer: &Lexer{
			input:              "äbcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		},
		expectedLexer: &Lexer{
			input:              "äbcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 1},
			},
			currentTokenStart: 0,
			positionInBuffer:  2,
		}},
	{
		name:      "EOF rune",
		inputRune: lexertoken.EOF,
		inputLexer: &Lexer{
			input:              "",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		},
		expectedLexer: &Lexer{
			input:              "",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		}},
	{
		name:      "ERROR rune",
		inputRune: lexertoken.ERROR,
		inputLexer: &Lexer{
			input:              "",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		},
		expectedLexer: &Lexer{
			input:              "",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		}},
}

func TestMoveAfterRune(t *testing.T) {
	for _, testCase := range moveAfterRunTestCases {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.inputLexer.MoveAfterRune(testCase.inputRune)

			assertEqualLexer(t, testCase.expectedLexer, testCase.inputLexer)

		})
	}
}

type peekRunTestCase struct {
	name          string
	expectedRune  rune
	expectedLexer *Lexer
}

var peekRuneTestCases = []peekRunTestCase{
	{
		name:         "one byte rune",
		expectedRune: 'c',
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "ab",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 2},
			},
			currentTokenStart: 0,
			positionInBuffer:  2,
		}},
	{
		name:         "two byte rune",
		expectedRune: 'ä',
		expectedLexer: &Lexer{
			input:              "äbcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		}},
	{
		name:         "empty input",
		expectedRune: lexertoken.EOF,
		expectedLexer: &Lexer{
			input:              "",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		}},
	{
		name:         "EOF",
		expectedRune: lexertoken.EOF,
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "ab",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 0,
			positionInBuffer:  5,
		}},
	{
		name:         "RuneError",
		expectedRune: lexertoken.ERROR,
		expectedLexer: &Lexer{
			input:              string(utf8.RuneError),
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 0},
			},
			currentTokenStart: 0,
			positionInBuffer:  0,
		}},
}

func TestPeekRune(t *testing.T) {
	for _, testCase := range peekRuneTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			lexer := *testCase.expectedLexer

			actualRune := lexer.PeekRune()

			assert.Equal(t, testCase.expectedRune, actualRune, "should be equal")
			assertEqualLexer(t, testCase.expectedLexer, &lexer)
		})
	}
}

type pushTokenTestCase struct {
	name           string
	inputTokenType lexertoken.TokenType
	expectedToken  lexertoken.Token
	inputLexer     *Lexer
	expectedLexer  *Lexer
}

var pushTokenTestCases = []pushTokenTestCase{
	{
		name:           "basic case",
		inputTokenType: lexertoken.TOKEN_LYRICS,
		expectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_LYRICS,
			Value: "bcd",
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "bcd",
				Start: lexertoken.TokenPosition{Line: 0, Column: 1},
				End:   lexertoken.TokenPosition{Line: 0, Column: 4},
			},
			currentTokenStart: 1,
			positionInBuffer:  4,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "",
				Start: lexertoken.TokenPosition{Line: 0, Column: 4},
				End:   lexertoken.TokenPosition{Line: 0, Column: 4},
			},
			currentTokenStart: 4,
			positionInBuffer:  4,
		},
	},
	{
		name:           "current token start position out of file",
		inputTokenType: lexertoken.TOKEN_LYRICS,
		expectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_ERROR,
			Value: lexererrors.LEXER_ERROR_START_OF_TOKEN_AFTER_EOF,
			Start: lexertoken.TokenPosition{Line: 0, Column: 0},
			End:   lexertoken.TokenPosition{Line: 0, Column: 5},
		},
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "abcde",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 6,
			positionInBuffer:  5,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_ERROR,
				Value: lexererrors.LEXER_ERROR_START_OF_TOKEN_AFTER_EOF,
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 6,
			positionInBuffer:  5,
		},
	},
	{
		name:           "lexer position out of file",
		inputTokenType: lexertoken.TOKEN_LYRICS,
		expectedToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_ERROR,
			Value: lexererrors.LEXER_ERROR_POSITION_AFTER_EOF,
			Start: lexertoken.TokenPosition{Line: 0, Column: 0},
			End:   lexertoken.TokenPosition{Line: 0, Column: 5},
		},
		inputLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_UNKNOWN,
				Value: "abcde",
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 0,
			positionInBuffer:  6,
		},
		expectedLexer: &Lexer{
			input:              "abcde",
			tokens:             make(chan lexertoken.Token, 5),
			nextLexingFunction: nil,
			currentToken: lexertoken.Token{
				Type:  lexertoken.TOKEN_ERROR,
				Value: lexererrors.LEXER_ERROR_POSITION_AFTER_EOF,
				Start: lexertoken.TokenPosition{Line: 0, Column: 0},
				End:   lexertoken.TokenPosition{Line: 0, Column: 5},
			},
			currentTokenStart: 0,
			positionInBuffer:  6,
		},
	},
}

func TestPushToken(t *testing.T) {
	for _, testCase := range pushTokenTestCases {
		t.Run(testCase.name, func(t *testing.T) {

			testCase.inputLexer.PushToken(testCase.inputTokenType)

			assertEqualLexer(t, testCase.expectedLexer, testCase.inputLexer)

			select {
			case actualToken := <-testCase.inputLexer.tokens:
				assert.Equal(t, testCase.expectedToken, actualToken)
			default:
				t.Fatal("token channel should contain a token")
			}

			select {
			case <-testCase.inputLexer.tokens:
				t.Fatal("no more token should be present on channel")
			default:
			}
		})
	}
}

type nextTokenTestCase struct {
	name              string
	expectedToken     lexertoken.Token
	expectedCallCount int
	inputLexer        *Lexer
	expectedLexer     *Lexer
}

var nextTokenTestCase_tokenInChannel = nextTokenTestCase{
	name:              "token in token channel",
	expectedToken:     testToken,
	expectedCallCount: 0,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_UNKNOWN,
			Value: "bcd",
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		positionInBuffer:  4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_UNKNOWN,
			Value: "bcd",
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		positionInBuffer:  4,
	},
}
var nextTokenTestCase_noTokenInChannel = nextTokenTestCase{
	name:              "no token in token channel",
	expectedToken:     testToken,
	expectedCallCount: 1,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_UNKNOWN,
			Value: "bcd",
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		positionInBuffer:  4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_UNKNOWN,
			Value: "bcd",
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		positionInBuffer:  4,
	},
}
var nextTokenTestCase_nilLexingFunc = nextTokenTestCase{
	name: "nil lexing function",
	expectedToken: lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: lexererrors.LEXER_ERROR_NIL_LEXING_FUNCTION,
		Start: lexertoken.TokenPosition{Line: 0, Column: 1},
		End:   lexertoken.TokenPosition{Line: 0, Column: 4},
	},

	expectedCallCount: 0,
	inputLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: nil,
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_UNKNOWN,
			Value: "bcd",
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},
		currentTokenStart: 1,
		positionInBuffer:  4,
	},
	expectedLexer: &Lexer{
		input:              "abcde",
		tokens:             make(chan lexertoken.Token, 5),
		nextLexingFunction: mockLexingFunction,
		currentToken: lexertoken.Token{
			Type:  lexertoken.TOKEN_ERROR,
			Value: lexererrors.LEXER_ERROR_NIL_LEXING_FUNCTION,
			Start: lexertoken.TokenPosition{Line: 0, Column: 1},
			End:   lexertoken.TokenPosition{Line: 0, Column: 4},
		},

		currentTokenStart: 1,
		positionInBuffer:  4,
	},
}

func TestNextToken(t *testing.T) {
	t.Run(nextTokenTestCase_tokenInChannel.name, func(t *testing.T) {

		mockLexingFuncCallCounter.ResetCount()

		nextTokenTestCase_tokenInChannel.inputLexer.tokens <- testToken

		actualToken := nextTokenTestCase_tokenInChannel.inputLexer.NextToken()

		assert.Equal(t, nextTokenTestCase_tokenInChannel.expectedToken, actualToken)
		assert.Equal(t, nextTokenTestCase_tokenInChannel.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, nextTokenTestCase_tokenInChannel.expectedLexer, nextTokenTestCase_tokenInChannel.inputLexer)
	})

	t.Run(nextTokenTestCase_noTokenInChannel.name, func(t *testing.T) {

		mockLexingFuncCallCounter.ResetCount()

		actualToken := nextTokenTestCase_noTokenInChannel.inputLexer.NextToken()

		assert.Equal(t, nextTokenTestCase_noTokenInChannel.expectedToken, actualToken)
		assert.Equal(t, nextTokenTestCase_noTokenInChannel.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, nextTokenTestCase_noTokenInChannel.expectedLexer, nextTokenTestCase_noTokenInChannel.inputLexer)
	})

	t.Run(nextTokenTestCase_nilLexingFunc.name, func(t *testing.T) {

		mockLexingFuncCallCounter.ResetCount()

		actualToken := nextTokenTestCase_nilLexingFunc.inputLexer.NextToken()

		assert.Equal(t, nextTokenTestCase_nilLexingFunc.expectedToken, actualToken)
		assert.Equal(t, nextTokenTestCase_nilLexingFunc.expectedCallCount, mockLexingFuncCallCounter.GetCount())
		assertEqualLexer(t, nextTokenTestCase_nilLexingFunc.expectedLexer, nextTokenTestCase_nilLexingFunc.inputLexer)
	})
}

func TestErrorf(t *testing.T) {
	// TODO
}

func TestSkipWitheSpace(t *testing.T) {
	// TODO
}

func TestNewLine(t *testing.T) {
	// TODO
}
