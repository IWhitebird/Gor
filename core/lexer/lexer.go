/*

---------------------------GOR LEXER-----------------------------
=================================================================
Package lexer implements a lexer for the Gor programming language.
The lexer takes a string of input and splits it into a slice of tokens.
=================================================================

*/

package lexer

import (
	"fmt"
	"strings"
)

// TokenType represents the type of a token.
type TokenType int

const (

	// Litreals
	Number TokenType = iota
	String
	Identifier

	// Operators
	Equals
	OpenParenthesis
	CloseParenthesis
	OpenBrace
	CloseBrace
	Colon
	SemiColon
	Comma
	OpenBracket
	CloseBracket
	Dot
	BinaryOperator

	Quote
	Greater
	Lesser
	EqualsOperator
	NotEqualsOperator
	OrOperator
	AndOperator
	GreaterOrEqual
	LesserOrEqual

	Ampersand
	Exclamation
	Bar

	// Keywords
	Let
	Const
	Function
	If
	Else
	Return
	For

	// End of File
	EOF
)

var KEYWORDS = map[string]TokenType{
	"let":    Let,
	"const":  Const,
	"fn":     Function,
	"if":     If,
	"else":   Else,
	"return": Return,
	"for":    For,
}

// Token represents a lexical token with a type and a value.
type Token struct {
	Type  TokenType
	Value string
}

// token creates a new token with the given type and value.
func token(value string, tokenType TokenType) Token {
	return Token{Value: value, Type: tokenType}
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isSkippable(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

// Tokenize generates a slice of tokens from the given input string.
func Tokenize(inputToken string) []Token {
	src := []byte(inputToken)
	srcLen := len(src)
	tokens := make([]Token, 0, srcLen/2)

	for i := 0; i < srcLen; i++ {
		c := src[i]

		switch c {
		case '#':
			for i+1 < srcLen && src[i+1] != '\n' && src[i+1] != '\r' {
				i++
			}
		case '=':
			if i+1 < srcLen && src[i+1] == '=' {
				tokens = append(tokens, token("==", EqualsOperator))
				i++
			} else {
				tokens = append(tokens, token("=", Equals))
			}
		case '+', '-', '*', '/', '%':
			tokens = append(tokens, token(string(c), BinaryOperator))
		case '(':
			tokens = append(tokens, token("(", OpenParenthesis))
		case ')':
			tokens = append(tokens, token(")", CloseParenthesis))
		case ':':
			tokens = append(tokens, token(":", Colon))
		case ';':
			tokens = append(tokens, token(";", SemiColon))
		case '{':
			tokens = append(tokens, token("{", OpenBrace))
		case '}':
			tokens = append(tokens, token("}", CloseBrace))
		case '[':
			tokens = append(tokens, token("[", OpenBracket))
		case ']':
			tokens = append(tokens, token("]", CloseBracket))
		case '.':
			tokens = append(tokens, token(".", Dot))
		case ',':
			tokens = append(tokens, token(",", Comma))
		case '!':
			if i+1 < srcLen && src[i+1] == '=' {
				tokens = append(tokens, token("!=", NotEqualsOperator))
				i++
			} else {
				tokens = append(tokens, token("!", Exclamation))
			}
		case '>':
			if i+1 < srcLen && src[i+1] == '=' {
				tokens = append(tokens, token(">=", GreaterOrEqual))
				i++
			} else {
				tokens = append(tokens, token(">", Greater))
			}
		case '<':
			if i+1 < srcLen && src[i+1] == '=' {
				tokens = append(tokens, token("<=", LesserOrEqual))
				i++
			} else {
				tokens = append(tokens, token("<", Lesser))
			}
		case '&':
			if i+1 < srcLen && src[i+1] == '&' {
				tokens = append(tokens, token("&&", AndOperator))
				i++
			} else {
				tokens = append(tokens, token("&", Ampersand))
			}
		case '|':
			if i+1 < srcLen && src[i+1] == '|' {
				tokens = append(tokens, token("||", OrOperator))
				i++
			} else {
				tokens = append(tokens, token("|", Bar))
			}
		case '"':
			var sb strings.Builder
			i++
			for i < srcLen && src[i] != '"' {
				sb.WriteByte(src[i])
				i++
			}
			tokens = append(tokens, token(sb.String(), String))
		default:
			if isNumber(c) {
				start := i
				for i+1 < srcLen && isNumber(src[i+1]) {
					i++
				}
				tokens = append(tokens, token(string(src[start:i+1]), Number))
			} else if isAlpha(c) {
				start := i
				for i+1 < srcLen && (isAlpha(src[i+1]) || isNumber(src[i+1])) {
					i++
				}
				word := string(src[start : i+1])
				if tokenType, ok := KEYWORDS[word]; ok {
					tokens = append(tokens, token(word, tokenType))
				} else {
					tokens = append(tokens, token(word, Identifier))
				}
			} else if isSkippable(c) {
				continue
			} else {
				fmt.Println(">> Lexer Error >>")
				fmt.Println("Unknown Token: ", string(c))
			}
		}
	}
	tokens = append(tokens, token("EndOfFile", EOF))

	return tokens
}
