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

// token returns a Token based on the given string.
func token(t string) Token {
	switch t {
	case "=":
		return Token{Equals, "="}
	case "+", "-", "*", "/":
		return Token{BinaryOperator, t}
	case "(":
		return Token{OpenParenthesis, "("}
	case ")":
		return Token{CloseParenthesis, ")"}
	default:
		if isNumber(t) {
			return Token{Number, t}
		} else if tokenType, exists := KEYWORDS[t]; exists {
			return Token{tokenType, t}
		} else {
			return Token{Identifier, t}
		}
	}
}

// Tokenize generates a slice of tokens from the given input string.
func Tokenize(inputToken string) []Token {
	var tokens []Token
	for _, t := range strings.Fields(inputToken) {
		t = strings.Trim(t, " \t\n")
		if t != "" {
			fmt.Println(t)
			tokens = append(tokens, token(t))
		}
	}
	return tokens
}

// isNumber checks if a given string represents a number.
func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
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
