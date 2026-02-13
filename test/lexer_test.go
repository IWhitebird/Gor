package test

import (
	"testing"

	LEX "github.com/iwhitebird/Gor/lexer"
)

func TestLexerEmptyInput(t *testing.T) {
	tokens := LEX.Tokenize("")
	if len(tokens) != 1 || tokens[0].Type != LEX.EOF {
		t.Fatalf("expected single EOF token, got %v", tokens)
	}
}

func TestLexerWhitespaceOnly(t *testing.T) {
	tokens := LEX.Tokenize("   \t\n  ")
	if len(tokens) != 1 || tokens[0].Type != LEX.EOF {
		t.Fatalf("expected single EOF token, got %v", tokens)
	}
}

func TestLexerNumber(t *testing.T) {
	tokens := LEX.Tokenize("42")
	if tokens[0].Type != LEX.Number || tokens[0].Value != "42" {
		t.Fatalf("expected Number 42, got %v", tokens[0])
	}
}

func TestLexerMultiDigitNumber(t *testing.T) {
	tokens := LEX.Tokenize("12345")
	if tokens[0].Type != LEX.Number || tokens[0].Value != "12345" {
		t.Fatalf("expected Number 12345, got %v", tokens[0])
	}
}

func TestLexerString(t *testing.T) {
	tokens := LEX.Tokenize(`"hello world"`)
	if tokens[0].Type != LEX.String || tokens[0].Value != "hello world" {
		t.Fatalf("expected String 'hello world', got %v", tokens[0])
	}
}

func TestLexerEmptyString(t *testing.T) {
	tokens := LEX.Tokenize(`""`)
	if tokens[0].Type != LEX.String || tokens[0].Value != "" {
		t.Fatalf("expected empty String, got %v", tokens[0])
	}
}

func TestLexerIdentifier(t *testing.T) {
	tokens := LEX.Tokenize("myVar")
	if tokens[0].Type != LEX.Identifier || tokens[0].Value != "myVar" {
		t.Fatalf("expected Identifier myVar, got %v", tokens[0])
	}
}

func TestLexerIdentifierWithUnderscore(t *testing.T) {
	tokens := LEX.Tokenize("my_var_2")
	if tokens[0].Type != LEX.Identifier || tokens[0].Value != "my_var_2" {
		t.Fatalf("expected Identifier my_var_2, got %v", tokens[0])
	}
}

func TestLexerKeywords(t *testing.T) {
	tests := []struct {
		input    string
		expected LEX.TokenType
	}{
		{"let", LEX.Let},
		{"const", LEX.Const},
		{"fn", LEX.Function},
		{"if", LEX.If},
		{"else", LEX.Else},
		{"return", LEX.Return},
		{"for", LEX.For},
	}
	for _, tt := range tests {
		tokens := LEX.Tokenize(tt.input)
		if tokens[0].Type != tt.expected {
			t.Errorf("keyword %q: expected type %d, got %d", tt.input, tt.expected, tokens[0].Type)
		}
	}
}

func TestLexerOperators(t *testing.T) {
	tokens := LEX.Tokenize("== != && || >= <=")
	expected := []struct {
		typ LEX.TokenType
		val string
	}{
		{LEX.EqualsOperator, "=="},
		{LEX.NotEqualsOperator, "!="},
		{LEX.AndOperator, "&&"},
		{LEX.OrOperator, "||"},
		{LEX.GreaterOrEqual, ">="},
		{LEX.LesserOrEqual, "<="},
	}
	for i, e := range expected {
		if tokens[i].Type != e.typ || tokens[i].Value != e.val {
			t.Errorf("token %d: expected (%d, %q), got (%d, %q)", i, e.typ, e.val, tokens[i].Type, tokens[i].Value)
		}
	}
}

func TestLexerSingleCharOperators(t *testing.T) {
	tokens := LEX.Tokenize("+ - * / % = > < ! & |")
	expectedTypes := []LEX.TokenType{
		LEX.BinaryOperator, LEX.BinaryOperator, LEX.BinaryOperator,
		LEX.BinaryOperator, LEX.BinaryOperator, LEX.Equals,
		LEX.Greater, LEX.Lesser, LEX.Exclamation,
		LEX.Ampersand, LEX.Bar,
	}
	for i, et := range expectedTypes {
		if tokens[i].Type != et {
			t.Errorf("token %d: expected type %d, got type %d (val=%q)", i, et, tokens[i].Type, tokens[i].Value)
		}
	}
}

func TestLexerDelimiters(t *testing.T) {
	tokens := LEX.Tokenize("( ) { } [ ] : ; , .")
	expectedTypes := []LEX.TokenType{
		LEX.OpenParenthesis, LEX.CloseParenthesis,
		LEX.OpenBrace, LEX.CloseBrace,
		LEX.OpenBracket, LEX.CloseBracket,
		LEX.Colon, LEX.SemiColon, LEX.Comma, LEX.Dot,
	}
	for i, et := range expectedTypes {
		if tokens[i].Type != et {
			t.Errorf("token %d: expected type %d, got type %d", i, et, tokens[i].Type)
		}
	}
}

func TestLexerComment(t *testing.T) {
	tokens := LEX.Tokenize("# this is a comment\n42")
	if tokens[0].Type != LEX.Number || tokens[0].Value != "42" {
		t.Fatalf("expected Number 42 after comment, got %v", tokens[0])
	}
}

func TestLexerCommentAtEndOfInput(t *testing.T) {
	tokens := LEX.Tokenize("42 # trailing comment")
	if tokens[0].Type != LEX.Number || tokens[0].Value != "42" {
		t.Fatalf("expected Number 42, got %v", tokens[0])
	}
	if tokens[1].Type != LEX.EOF {
		t.Fatalf("expected EOF after comment, got %v", tokens[1])
	}
}

func TestLexerOperatorAtEndOfInput(t *testing.T) {
	// Edge case: operator at end of input should not panic
	tokens := LEX.Tokenize("x =")
	if tokens[0].Type != LEX.Identifier {
		t.Fatalf("expected Identifier, got %v", tokens[0])
	}
	if tokens[1].Type != LEX.Equals {
		t.Fatalf("expected Equals, got %v", tokens[1])
	}
}

func TestLexerExclamationAtEnd(t *testing.T) {
	tokens := LEX.Tokenize("!")
	if tokens[0].Type != LEX.Exclamation {
		t.Fatalf("expected Exclamation, got %v", tokens[0])
	}
}

func TestLexerGreaterAtEnd(t *testing.T) {
	tokens := LEX.Tokenize(">")
	if tokens[0].Type != LEX.Greater {
		t.Fatalf("expected Greater, got %v", tokens[0])
	}
}

func TestLexerLesserAtEnd(t *testing.T) {
	tokens := LEX.Tokenize("<")
	if tokens[0].Type != LEX.Lesser {
		t.Fatalf("expected Lesser, got %v", tokens[0])
	}
}

func TestLexerAmpersandAtEnd(t *testing.T) {
	tokens := LEX.Tokenize("&")
	if tokens[0].Type != LEX.Ampersand {
		t.Fatalf("expected Ampersand, got %v", tokens[0])
	}
}

func TestLexerBarAtEnd(t *testing.T) {
	tokens := LEX.Tokenize("|")
	if tokens[0].Type != LEX.Bar {
		t.Fatalf("expected Bar, got %v", tokens[0])
	}
}

func TestLexerComplexExpression(t *testing.T) {
	tokens := LEX.Tokenize("let x = 1 + 2 * 3")
	if len(tokens) != 9 { // let x = 1 + 2 * 3 EOF
		t.Fatalf("expected 8 tokens, got %d", len(tokens))
	}
}
