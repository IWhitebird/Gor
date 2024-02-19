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
	"strconv"
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

// KEYWORDS maps keyword strings to their corresponding token types.
const (
	//Variable Declares
	letStr   = "let"
	constStr = "const"

	//Function Declaration
	fnStr     = "fn"
	ifStr     = "if"
	elseStr   = "else"
	returnStr = "return"
	forStr    = "for"
)

var KEYWORDS = map[string]TokenType{
	letStr:    Let,
	constStr:  Const,
	fnStr:     Function,
	ifStr:     If,
	elseStr:   Else,
	returnStr: Return,
	forStr:    For,
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

// isNumber checks if a given string represents a number.
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// isAlpha checks if a given string represents an alphabetic character.
func isAlpha(s string) bool {
	return (s >= "a" && s <= "z") || (s >= "A" && s <= "Z") || s == "_"
}

// isSkippable check if the given character is useless
func isSkippable(s string) bool {
	return s == " " || s == "\t" || s == "\n" || s == "\r"
}

// Tokenize generates a slice of tokens from the given input string.
func Tokenize(inputToken string) []Token {
	var tokens []Token
	for i := 0; i < len(inputToken); i++ {

		var t string = inputToken[i : i+1]

		if t != "" {

			switch t {
            case "#":
                for j := i + 1; j < len(inputToken); j++ {
                    if inputToken[j:j+1] == "\n" {
                        i = j
                        break
                    }
                }
			case "=":
				if inputToken[i+1:i+2] == "=" {
					tokens = append(tokens, token("==", EqualsOperator))
					i++
				} else {
					tokens = append(tokens, token(t, Equals))
				}
			case "+", "-", "*", "/", "%":
				tokens = append(tokens, token(t, BinaryOperator))
			case "(":
				tokens = append(tokens, token(t, OpenParenthesis))
			case ")":
				tokens = append(tokens, token(t, CloseParenthesis))
			case ":":
				tokens = append(tokens, token(t, Colon))
			case ";":
				tokens = append(tokens, token(t, SemiColon))
			case "{":
				tokens = append(tokens, token(t, OpenBrace))
			case "}":
				tokens = append(tokens, token(t, CloseBrace))
			case "[":
				tokens = append(tokens, token(t, OpenBracket))
			case "]":
				tokens = append(tokens, token(t, CloseBracket))
			case ".":
				tokens = append(tokens, token(t, Dot))
			case ",":
				tokens = append(tokens, token(t, Comma))
			case "!":
				if inputToken[i+1:i+2] == "=" {
					tokens = append(tokens, token("!=", NotEqualsOperator))
					i++
				} else {
					tokens = append(tokens, token(t, Exclamation))
				}
			case ">":
				if inputToken[i+1:i+2] == "=" {
					tokens = append(tokens, token(">=", GreaterOrEqual))
					i++
				} else {
					tokens = append(tokens, token(t, Greater))
				}
			case "<":
				if inputToken[i+1:i+2] == "=" {
					tokens = append(tokens, token("<=", LesserOrEqual))
					i++
				} else {
					tokens = append(tokens, token(t, Lesser))
				}
			case "&":
				if inputToken[i+1:i+2] == "&" {
					tokens = append(tokens, token("&&", AndOperator))
					i++
				} else {
					tokens = append(tokens, token(t, Ampersand))
				}
			case "|":
				if inputToken[i+1:i+2] == "|" {
					tokens = append(tokens, token("||", OrOperator))
					i++
				} else {
					tokens = append(tokens, token(t, Bar))
				}
			case "\"":
				var stringLiteral string = ""
				for j := i + 1; j < len(inputToken) && inputToken[j:j+1] != "\""; j++ {
					stringLiteral += inputToken[j : j+1]
					i++
				}
				tokens = append(tokens, token(stringLiteral, String))
				i++
			default:
				if isNumber(t) {
					var number string = t
					for j := i + 1; j < len(inputToken); j++ {
						if isNumber(inputToken[j : j+1]) {
							number += inputToken[j : j+1]
							i++
						} else {
							break
						}
					}
					tokens = append(tokens, token(number, Number))
				} else if isAlpha(t) {
					var identifier string = t
					for j := i + 1; j < len(inputToken); j++ {
						if isAlpha(inputToken[j : j+1]) {
							identifier += inputToken[j : j+1]
							i++
						} else {
							break
						}
					}
					if tokenType, ok := KEYWORDS[identifier]; ok {
						tokens = append(tokens, token(identifier, tokenType))
					} else {
						tokens = append(tokens, token(identifier, Identifier))
					}
				} else if isSkippable(t) {
					continue
				} else {
					fmt.Println(">> Lexer Error >>")
					fmt.Println("Unknown Token: ", t)
				}
			}

		}
	}
	tokens = append(tokens, token("EndOfFile", EOF))
	return tokens
}

func Main() {

	// fmt.Println(">> Lexer Running >>")

	// file, err := os.Open("input.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// // Create a new Scanner for the file
	// scanner := bufio.NewScanner(file)

	// // Loop over all lines in the file and print them
	// var input string
	// for scanner.Scan() {
	// 	input += scanner.Text()
	// }

	// // Remove the newline character
	// input = strings.Replace(input, "\n", "", -1)

	// // Tokenize the inputToken
	// tokens := Tokenize(input)

	// for _, token := range tokens {

	// 	fmt.Println("Token Type: ", token.Type, "Token Value: ", token.Value)
	// }

}
