package parser

import (
	AST "Gor/ast"
	Lex "Gor/lexer"
	"fmt"
	"os"
	"strconv"
)

type Parser struct {
	tokens []Lex.Token
}

func (p *Parser) peek() Lex.Token {
	return p.tokens[0]
}

func (p *Parser) consume() Lex.Token {
	var token Lex.Token = p.tokens[0]
	p.tokens = p.tokens[1:]
	return token
}

func (p *Parser) not_Eof() bool {
	return p.tokens[0].Type != Lex.EOF
}

func (p *Parser) produceAst(sourceCode string) AST.Program {

	p.tokens = Lex.Tokenize(sourceCode)

	var program AST.Program = AST.Program{KindValue: AST.ProgramType, Body: []AST.Stmt{}}

	for p.not_Eof() {
		program.Body = append(p.parseStmt())
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
	return p.parsePrimaryExpr()
}

func (p *Parser) parsePrimaryExpr() AST.Expr {

	tk := p.peek().Type

	switch tk {

	case Lex.Identifier:
		return AST.Identifier{KindValue: "Identifier", Symbol: p.consume().Value}

	case Lex.Number:
		return AST.NumericLiteral{KindValue: "Number", Value: p.parseInt(p.consume().Value)}

	case Lex.BinaryOperator:
		return AST.BinaryExpr{
			KindValue: "BinaryExpr", 
			Left: p.consume().Value, 
			Right: p.consume().Value, 
			Operator: p.consume().Value
		}
	
	default:
		fmt.Println("Error: Invalid Token")
		os.Exit(1)

	}

	return AST.Expr{}
}

func Main() {
	fmt.Println(">> Welcome To Gor Parser >:D")
	fmt.Println(Lex.Tokenize("var a = 1 + 2"))
}
