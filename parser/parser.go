package parser

import (
	AST "Gor/ast"
	LEX "Gor/lexer"
	"fmt"
	"os"
	"strconv"
)

type Parser struct {
	tokens []LEX.Token
}

func (p *Parser) peek() LEX.Token {
	return p.tokens[0]
}

func (p *Parser) consume() LEX.Token {
	var token LEX.Token = p.peek()
	p.tokens = p.tokens[1:]

	if token.Type == LEX.Null {
		token.Value = "null"
	}
	return token
}

// func (p *Parser) not_check(tokenTypes []LEX.TokenType, errorMessage string) bool {

// 	for _, tokenType := range tokenTypes {
// 		if p.peek().Type == tokenType {
// 			return false
// 		}
// 	}

// 	fmt.Println(errorMessage)
// 	os.Exit(1)
// 	return true
// }

func (p *Parser) expect(tokenType LEX.TokenType, errorMessage string) {
	if p.peek().Type == tokenType {
		p.consume()
	} else {
		fmt.Println(errorMessage)
		os.Exit(1)
	}
}

func (p *Parser) not_Eof() bool {
	return p.tokens[0].Type != LEX.EOF
}

func (p *Parser) ProduceAst(sourceCode string) AST.Program {

	p.tokens = LEX.Tokenize(sourceCode)

	var program AST.Program = AST.Program{KindValue: AST.ProgramType, Body: []AST.Stmt{}}

	for p.not_Eof() {
		program.Body = append(program.Body, p.parseStmt())
	}

	return program
}

func (p *Parser) parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: Invalid Number")
		os.Exit(1)
	}
	return i
}

func (p *Parser) parseStmt() AST.Stmt {
	return p.parseExpr()
}

func (p *Parser) parseExpr() AST.Expr {
	return p.parseAdditiveExpr()
}

func (p *Parser) parseAdditiveExpr() AST.Expr {
	left := p.parseMultiplicativeExpr()

	// p.not_check([]LEX.TokenType{LEX.BinaryOperator, LEX.OpenParenthesis, LEX.CloseParenthesis, LEX.Identifier, LEX.Number}, "Error: Invalid Token")

	for p.peek().Value == "+" || p.peek().Value == "-" {
		op := p.consume().Value
		right := p.parseMultiplicativeExpr()
		left = AST.BinaryExpr{KindValue: "BinaryExpr", Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseMultiplicativeExpr() AST.Expr {
	left := p.parsePrimaryExpr()

	// p.not_check([]LEX.TokenType{LEX.BinaryOperator, LEX.OpenParenthesis, LEX.CloseParenthesis, LEX.Identifier, LEX.Number}, "Error: Invalid Token")

	for p.peek().Value == "*" || p.peek().Value == "/" || p.peek().Value == "%" {
		op := p.consume().Value
		right := p.parsePrimaryExpr()
		left = AST.BinaryExpr{KindValue: "BinaryExpr", Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parsePrimaryExpr() AST.Expr {

	tk := p.peek().Type
	switch tk {

	case LEX.Identifier:
		return AST.Identifier{KindValue: "Identifier", Symbol: p.consume().Value}
	case LEX.Null:
		return AST.NullLiteral{KindValue: "NullLiteral", Value: p.consume().Value}
	case LEX.Number:
		return AST.NumericLiteral{KindValue: "NumericLiteral", Value: p.parseInt(p.consume().Value)}
	case LEX.OpenParenthesis:
		p.consume()
		expr := p.parseExpr()
		p.expect(LEX.CloseParenthesis, "Error: Missing Closing Parenthesis")
		return expr
	default:
		fmt.Println("Error: Invalid Token")
		return nil
	}
}

func Main() {
	// fmt.Println(">> Parser Running >>")
}
