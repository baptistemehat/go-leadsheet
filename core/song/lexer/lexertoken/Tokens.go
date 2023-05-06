package lexertoken

import "unicode/utf8"

const EOF rune = 0
const ERROR rune = utf8.RuneError

const LEFT_PARENTHESIS rune = '('
const RIGHT_PARENTHESIS rune = ')'
const LEFT_BRACKET rune = '['
const RIGHT_BRACKET rune = ']'
const LEFT_BRACE rune = '{'
const RIGHT_BRACE rune = '}'
const COLUMN rune = ':'
const NEWLINE rune = '\n'
