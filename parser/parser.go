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
	case LEX.Function:
		return p.parseFunctionDeclaration()
	case LEX.If:
		return p.parseIfStatement()
	case LEX.For:
		return p.parseForStatement()
	default:
		return p.parseExpr()
	}
}

func (p *Parser) parseIfStatement() AST.Stmt {
	p.consume()

	p.expect(LEX.OpenParenthesis, "Error Missing Opening Parenthesis")

	left_expr := p.parseExpr()

	p.expect(LEX.CloseParenthesis, "Error Missing Closing Parenthesis")

	body := p.parseBlockStmt()

	alternate := []AST.Stmt{}

	if p.peek().Type == LEX.Else {
		p.consume()

		if p.peek().Type == LEX.If {
			alternate = append(alternate, p.parseIfStatement())
		} else {
			alternate = p.parseBlockStmt()
		}
	}

	return AST.IfStmt{KindValue: AST.IfStmtType, Test: left_expr, Body: body, Alternate: alternate}

}

func (p *Parser) parseForStatement() AST.Stmt {
	p.consume()
	p.expect(LEX.OpenParenthesis, "Error Missing Opening Parenthesis")

	init := p.parseVariableDeclaration()

	p.expect(LEX.SemiColon, "Error Missing Semicolon")

	test := p.parseExpr()

	p.expect(LEX.SemiColon, "Error Missing Semicolon")

	update := p.parseExpr()

	p.expect(LEX.CloseParenthesis, "Error Missing Closing Parenthesis")

	body := p.parseBlockStmt()

	return AST.ForStmt{KindValue: AST.ForStmtType, Init: init, Test: test, Update: update, Body: body}
}

func (p *Parser) parseFunctionDeclaration() AST.Stmt {
	p.consume()
	p.expect(LEX.Identifier, "Error: Missing Identifier")

	args := p.parseArguments()
	params := []string{}

	for _, arg := range args {
		if arg.Kind() != AST.IdentifierType {
			fmt.Println("Error: Function Parameters must be Identifiers")
			os.Exit(1)
		}
		params = append(params, arg.(AST.Identifier).Symbol)
	}

	body := p.parseBlockStmt()

	return AST.FunctionDeclaration{KindValue: AST.FunctionDeclarationType, Identifier: p.peek().Value, Parameters: params, Body: body}
}

func (p *Parser) parseBlockStmt() []AST.Stmt {
	p.expect(LEX.OpenBrace, "Error: Missing Opening Brace")

	var body []AST.Stmt = []AST.Stmt{}

	for p.not_Eof() && p.peek().Type != LEX.CloseBrace {
		body = append(body, p.parseStmt())
	}

	p.expect(LEX.CloseBrace, "Error: Missing Closing Brace")

	return body
}

func (p *Parser) parseVariableDeclaration() AST.Stmt {

	isConst := p.peek().Type == LEX.Const
	p.consume()

	isIdentifier := p.expect(LEX.Identifier, "Error: Missing Identifier")

	if p.peek().Type != LEX.Equals {

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

	return dec

}

func (p *Parser) parseExpr() AST.Expr {
	return p.parseAssignmentExpr()
}

func (p *Parser) parseAssignmentExpr() AST.Expr {
	left := p.parseOjbectExpr()

	if p.peek().Type == LEX.Equals {
		p.consume()
		right := p.parseAssignmentExpr()
		return AST.AssignmentExpr{KindValue: AST.AssignmentExprType, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseAndStatement() AST.Expr {
	left := p.parseAdditiveExpr()

	if p.peek().Value == "&&" || p.peek().Value == "||" {
		op := p.consume().Value
		right := p.parseAdditiveExpr()
		left = AST.BinaryExpr{KindValue: "BinaryExpr", Operator: op, Left: left, Right: right}
	}
	return left
}

func (p *Parser) parseOjbectExpr() AST.Expr {

	// NO OPEN BRACE

	if p.peek().Type != LEX.OpenBrace {
		return p.parseAndStatement()
	}

	// OPEN BRACE
	p.consume()
	properties := []AST.Property{}

	// { key }
	for p.not_Eof() && p.peek().Type != LEX.CloseBrace {
		key := p.expect(LEX.Identifier, "Error: Missing Identifier")

		if p.peek().Type == LEX.Comma {
			p.consume()
			prop := AST.Property{KindValue: AST.PropertyType, Key: key.Value, Value: nil}
			properties = append(properties, prop)
			continue
		} else if p.peek().Type == LEX.CloseBrace {
			prop := AST.Property{KindValue: AST.PropertyType, Key: key.Value, Value: nil}
			properties = append(properties, prop)
			continue
		} else if p.peek().Type == LEX.Colon {
			p.consume()
			value := p.parseExpr()
			prop := AST.Property{KindValue: AST.PropertyType, Key: key.Value, Value: value}
			properties = append(properties, prop)
		}

		if p.peek().Type != LEX.CloseBrace {
			p.expect(LEX.Comma, "Error: Missing Comma")
		}
	}

	p.expect(LEX.CloseBrace, "Error: Missing Closing Brace")

	return AST.ObjectLiteral{KindValue: AST.OjectLiteralType, Properties: properties}
}

func (p *Parser) parseAdditiveExpr() AST.Expr {
	left := p.parseMultiplicativeExpr()

	for p.peek().Value == "+" || p.peek().Value == "-" || p.peek().Value == "==" || p.peek().Value == "!=" || p.peek().Value == "<" || p.peek().Value == ">" || p.peek().Value == "<=" || p.peek().Value == ">=" {
		op := p.consume().Value
		right := p.parseMultiplicativeExpr()
		left = AST.BinaryExpr{KindValue: "BinaryExpr", Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseMultiplicativeExpr() AST.Expr {
	left := p.parseCallMemberExpr()

	for p.peek().Value == "*" || p.peek().Value == "/" || p.peek().Value == "%" {
		op := p.consume().Value
		right := p.parseCallMemberExpr()
		left = AST.BinaryExpr{KindValue: "BinaryExpr", Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseCallMemberExpr() AST.Expr {
	member := p.parseMemberExpr()

	if p.peek().Type == LEX.OpenParenthesis {
		return p.parseCallExpr(member)
	}
	return member
}

func (p *Parser) parseCallExpr(caller AST.Expr) AST.Expr {
	var call AST.Expr = AST.CallExpr{KindValue: "CallExpr", Caller: caller, Arguments: p.parseArguments()}

	if p.peek().Type == LEX.OpenBrace {
		call = p.parseCallExpr(call)
	}
	return call
}

func (p *Parser) parseArguments() []AST.Expr {

	p.expect(LEX.OpenParenthesis, "Error: Missing Opening Parenthesis")

	var arguments []AST.Expr = []AST.Expr{}

	if p.peek().Type != LEX.CloseParenthesis {
		arguments = p.parseArgumentsList()
	}

	p.expect(LEX.CloseParenthesis, "Error: Missing Closing Parenthesis")

	return arguments
}

func (p *Parser) parseArgumentsList() []AST.Expr {
	var args []AST.Expr = []AST.Expr{p.parseAssignmentExpr()}

	for p.peek().Type == LEX.Comma {
		p.consume()
		args = append(args, p.parseAssignmentExpr())
	}

	return args
}

func (p *Parser) parseMemberExpr() AST.Expr {
	var object AST.Expr = p.parsePrimaryExpr()

	for p.peek().Type == LEX.Dot || p.peek().Type == LEX.OpenBracket {

		operator := p.consume()
		var property AST.Expr
		var computed bool

		if operator.Type == LEX.Dot {

			property = p.parsePrimaryExpr()
			computed = false

			if property.Kind() != AST.IdentifierType {
				fmt.Println("Error: Property must be an Identifier")
				os.Exit(1)
			}

		}

		if operator.Type == LEX.OpenBracket {
			property = p.parseExpr()
			computed = true
			p.expect(LEX.CloseBracket, "Error: Missing Closing Bracket")
		}

		object = AST.MemberExpr{KindValue: "MemberExpr", Object: object, Property: property, Computed: computed}
	}

	return object
}

func (p *Parser) parsePrimaryExpr() AST.Expr {

	tk := p.peek().Type
	switch tk {

	case LEX.Identifier:
		return AST.Identifier{KindValue: "Identifier", Symbol: p.consume().Value}
	case LEX.Number:
		return AST.NumericLiteral{KindValue: "NumericLiteral", Value: p.parseInt(p.consume().Value)}
	case LEX.String:
		return AST.StringLiteral{KindValue: "StringLiteral", Value: p.consume().Value}
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
