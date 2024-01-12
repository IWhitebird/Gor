package lexer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TokenType represents the type of a token.
type TokenType int

const (
	Number           TokenType = iota // 0
	Identifier                        // 1
	Equals                            // 2
	OpenParenthesis                   // 3
	CloseParenthesis                  // 4
	BinaryOperator                    // 5

	// Keywords
	Let    // 6
	If     // 7
	Else   // 8
	Return // 9
)

const (
	letStr    = "let"
	ifStr     = "if"
	elseStr   = "else"
	returnStr = "return"
)

// KEYWORDS maps keyword strings to their corresponding token types.
var KEYWORDS = map[string]TokenType{
	letStr:    Let,
	ifStr:     If,
	elseStr:   Else,
	returnStr: Return,
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
	return strings.ToLower(s) == strings.ToUpper(s)
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
			case "+", "-", "*", "/":
				tokens = append(tokens, token(t, BinaryOperator))
			case "(":
				tokens = append(tokens, token(t, OpenParenthesis))
			case ")":
				tokens = append(tokens, token(t, CloseParenthesis))

			}

		}
	}
	return tokens
}

func Main() {
	// Take Input from User
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print(">> ")
	// input, _ := reader.ReadString('\n')

	// Take Input from File

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Loop over all lines in the file and print them
	var input string
	for scanner.Scan() {
		input += scanner.Text()
	}

	// Remove the newline character
	input = strings.Replace(input, "\n", "", -1)

	// Tokenize the inputToken
	tokens := Tokenize(input)

	for _, token := range tokens {
		fmt.Println("Token Type: ", token.Type, " Token Value: ", token.Value)
	}

}
