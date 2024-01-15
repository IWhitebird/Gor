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
	Number     TokenType = iota // 0
	Identifier                  // 2
	String

	// Operators
	Equals           // 3
	OpenParenthesis  // 4
	CloseParenthesis // 5
	Colon            // 1
	BinaryOperator   // 6

	// Keywords
	Let    // 7
	Const  // 11
	If     // 8
	Else   // 9
	Return // 10

	// End of File
	EOF // 11
)

// KEYWORDS maps keyword strings to their corresponding token types.
const (
	letStr    = "let"
	ifStr     = "if"
	elseStr   = "else"
	returnStr = "return"
	constStr  = "const"
)

var KEYWORDS = map[string]TokenType{
	letStr:    Let,
	ifStr:     If,
	elseStr:   Else,
	returnStr: Return,
	constStr:  Const,
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
	return s == " " || s == "\t" || s == "\n"
}

// Tokenize generates a slice of tokens from the given input string.
func Tokenize(inputToken string) []Token {
	var tokens []Token
	for i := 0; i < len(inputToken); i++ {

		var t string = inputToken[i : i+1]

		if t == " " || t == "\t" || t == "\n" {
			continue
		}

		if t != "" {

			switch t {
			case "=":
				tokens = append(tokens, token(t, Equals))
			case "+", "-", "*", "/", "%":
				tokens = append(tokens, token(t, BinaryOperator))
			case "(":
				tokens = append(tokens, token(t, OpenParenthesis))
			case ")":
				tokens = append(tokens, token(t, CloseParenthesis))
			case ":":
				tokens = append(tokens, token(t, Colon))
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
