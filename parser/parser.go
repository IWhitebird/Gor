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

func (p *Parser) expect(tokenType LEX.TokenType, errorMessage string) LEX.Token {
	if p.peek().Type == tokenType {
		return p.consume()
	} else {
		fmt.Println(errorMessage)
		os.Exit(1)
	}
	return LEX.Token{}
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

	switch p.peek().Type {
	case LEX.Let:
		return p.parseVariableDeclaration()
	case LEX.Const:
		return p.parseVariableDeclaration()
	default:
		return p.parseExpr()
	}
}

func (p *Parser) parseVariableDeclaration() AST.Stmt {

	isConst := p.peek().Type == LEX.Const
	p.consume()

	isIdentifier := p.expect(LEX.Identifier, "Error: Missing Identifier")

	if p.peek().Type == LEX.Colon {
		p.consume()

		if isConst {
			fmt.Println("Error: Const Variable Declaration cannot be without a value")
			os.Exit(1)
		}

		return AST.VariableDeclaration{
			KindValue:  AST.VariableDeclarationType,
			Constant:   isConst,
			Identifier: isIdentifier.Value,
		}
	}

	p.expect(LEX.Equals, "Error: Missing Equals")

	dec := AST.VariableDeclaration{
		KindValue:  AST.VariableDeclarationType,
		Constant:   isConst,
		Identifier: isIdentifier.Value,
		Value:      p.parseExpr(),
	}

	p.expect(LEX.Colon, "Error: Missing Colon")

	return dec

}

func (p *Parser) parseExpr() AST.Expr {
	return p.parseAssignmentExpr()
}

func (p *Parser) parseAssignmentExpr() AST.Expr {
	left := p.parseAdditiveExpr()

	if p.peek().Type == LEX.Equals {
		p.consume()
		right := p.parseAdditiveExpr()
		return AST.AssignmentExpr{KindValue: AST.AssignmentExprType, Left: left, Right: right}
	}

	return left
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
	case LEX.Number:
		return AST.NumericLiteral{KindValue: "NumericLiteral", Value: p.parseInt(p.consume().Value)}
	case LEX.OpenParenthesis:
		p.consume()
		expr := p.parseExpr()
		p.expect(LEX.CloseParenthesis, "Error: Missing Closing Parenthesis")
		return expr
	default:
		fmt.Println("Error: Invalid Token")
		os.Exit(1)
		return nil
	}
}

func Main() {
	// fmt.Println(">> Parser Running >>")
}
