package parser

import (
	"strconv"

	AST "github.com/iwhitebird/Gor/core/ast"
	LEX "github.com/iwhitebird/Gor/core/lexer"
)

type Parser struct {
	tokens []LEX.Token
	pos    int
}

func (p *Parser) peek() LEX.Token {
	return p.tokens[p.pos]
}

func (p *Parser) consume() LEX.Token {
	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) expect(tokenType LEX.TokenType, errorMessage string) LEX.Token {
	if p.peek().Type == tokenType {
		return p.consume()
	}
	panic(errorMessage)
}

func (p *Parser) not_Eof() bool {
	return p.tokens[p.pos].Type != LEX.EOF
}

func (p *Parser) ProduceAst(sourceCode string) AST.Program {
	p.tokens = LEX.Tokenize(sourceCode)
	p.pos = 0

	var program AST.Program = AST.Program{Body: make([]AST.Stmt, 0, 16)}

	for p.not_Eof() {
		program.Body = append(program.Body, p.parseStmt())
	}

	return program
}

func (p *Parser) parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic("Error: Invalid Number")
	}
	return i
}

func (p *Parser) parseStmt() AST.Stmt {
	switch p.peek().Type {
	case LEX.Let:
		return p.parseVariableDeclaration()
	case LEX.Const:
		return p.parseVariableDeclaration()
	case LEX.Return:
		return p.parseReturnStatement()
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

func (p *Parser) parseVectorDeclaration() AST.Stmt {
	p.consume()

	elements := make([]AST.Expr, 0, 4)

	for p.not_Eof() && p.peek().Type != LEX.CloseBracket {
		elements = append(elements, p.parseExpr())

		if p.peek().Type != LEX.CloseBracket {
			p.expect(LEX.Comma, "Error: Missing Comma")
		}
	}

	p.expect(LEX.CloseBracket, "Error: Missing Closing Bracket")

	return AST.VectorLiteral{Elements: elements}
}

func (p *Parser) parseReturnStatement() AST.Stmt {
	p.consume()
	value := p.parseExpr()
	return AST.ReturnStmt{Value: value}
}

func (p *Parser) parseIfStatement() AST.Stmt {
	p.consume()

	p.expect(LEX.OpenParenthesis, "Error Missing Opening Parenthesis")

	left_expr := p.parseExpr()

	p.expect(LEX.CloseParenthesis, "Error Missing Closing Parenthesis")

	body := p.parseBlockStmt()

	var alternate AST.Stmt

	if p.peek().Type == LEX.Else {
		p.consume()

		if p.peek().Type == LEX.If {
			alternate = p.parseIfStatement().(AST.IfStmt)
		} else {
			alternate = p.parseBlockStmt().(AST.BlockStmt)
		}
	}

	return AST.IfStmt{Test: left_expr, Body: body.(AST.BlockStmt), Alternate: alternate}
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

	return AST.ForStmt{Init: init, Test: test, Update: update, Body: body.(AST.BlockStmt)}
}

func (p *Parser) parseFunctionDeclaration() AST.Stmt {
	p.consume()
	iden := p.expect(LEX.Identifier, "Error: Missing Identifier")

	args := p.parseArguments()
	params := make([]string, 0, len(args))

	for _, arg := range args {
		if arg.Kind() != AST.IdentifierType {
			panic("Error: Function Parameters must be Identifiers")
		}
		params = append(params, arg.(AST.Identifier).Symbol)
	}

	body := p.parseBlockStmt()

	return AST.FunctionDeclaration{Identifier: iden.Value, Parameters: params, Body: body.(AST.BlockStmt)}
}

func (p *Parser) parseBlockStmt() AST.Stmt {
	p.expect(LEX.OpenBrace, "Error: Missing Opening Brace")

	body := make([]AST.Stmt, 0, 8)

	for p.not_Eof() && p.peek().Type != LEX.CloseBrace {
		body = append(body, p.parseStmt())
	}

	p.expect(LEX.CloseBrace, "Error: Missing Closing Brace")

	return AST.BlockStmt{Body: body}
}

func (p *Parser) parseVariableDeclaration() AST.Stmt {

	isConst := p.peek().Type == LEX.Const
	p.consume()

	isIdentifier := p.expect(LEX.Identifier, "Error: Missing Identifier")

	if p.peek().Type != LEX.Equals {

		if isConst {
			panic("Error: Const Variable Declaration cannot be without a value")
		}

		return AST.VariableDeclaration{
			Constant:   isConst,
			Identifier: isIdentifier.Value,
		}
	}

	p.expect(LEX.Equals, "Error: Missing Equals")

	dec := AST.VariableDeclaration{
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
		return AST.AssignmentExpr{Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseAndStatement() AST.Expr {
	left := p.parseComparisonExpr()

	// FIX: use for loop instead of if so chained &&/|| works (a && b && c)
	for p.peek().Value == "&&" || p.peek().Value == "||" {
		op := p.consume().Value
		right := p.parseComparisonExpr()
		left = AST.BinaryExpr{Operator: op, Left: left, Right: right}
	}
	return left
}

func (p *Parser) parseComparisonExpr() AST.Expr {
	left := p.parseAdditiveExpr()

	for p.peek().Value == "==" || p.peek().Value == "!=" || p.peek().Value == "<" || p.peek().Value == ">" || p.peek().Value == "<=" || p.peek().Value == ">=" {
		op := p.consume().Value
		right := p.parseAdditiveExpr()
		left = AST.BinaryExpr{Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseOjbectExpr() AST.Expr {

	if p.peek().Type != LEX.OpenBrace {
		return p.parseAndStatement()
	}

	// OPEN BRACE
	p.consume()
	properties := make([]AST.Property, 0, 4)

	// { key }
	for p.not_Eof() && p.peek().Type != LEX.CloseBrace {
		key := p.expect(LEX.Identifier, "Error: Missing Identifier")

		if p.peek().Type == LEX.Comma {
			p.consume()
			prop := AST.Property{Key: key.Value, Value: nil}
			properties = append(properties, prop)
			continue
		} else if p.peek().Type == LEX.CloseBrace {
			prop := AST.Property{Key: key.Value, Value: nil}
			properties = append(properties, prop)
			continue
		} else if p.peek().Type == LEX.Colon {
			p.consume()
			value := p.parseExpr()
			prop := AST.Property{Key: key.Value, Value: value}
			properties = append(properties, prop)
		}

		if p.peek().Type != LEX.CloseBrace {
			p.expect(LEX.Comma, "Error: Missing Comma")
		}
	}

	p.expect(LEX.CloseBrace, "Error: Missing Closing Brace")

	return AST.ObjectLiteral{Properties: properties}
}

func (p *Parser) parseAdditiveExpr() AST.Expr {
	left := p.parseMultiplicativeExpr()

	for p.peek().Value == "+" || p.peek().Value == "-" || p.peek().Value == "&" || p.peek().Value == "|" {
		op := p.consume().Value
		right := p.parseMultiplicativeExpr()
		left = AST.BinaryExpr{Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseMultiplicativeExpr() AST.Expr {
	left := p.parseCallMemberExpr()

	for p.peek().Value == "*" || p.peek().Value == "/" || p.peek().Value == "%" {
		op := p.consume().Value
		right := p.parseCallMemberExpr()
		left = AST.BinaryExpr{Operator: op, Left: left, Right: right}
	}

	return left
}

func (p *Parser) parseCallMemberExpr() AST.Expr {
	member := p.parseMemberAndVectorExpr()

	if p.peek().Type == LEX.OpenParenthesis {
		return p.parseCallExpr(member)
	}
	return member
}

func (p *Parser) parseCallExpr(caller AST.Expr) AST.Expr {
	var call AST.Expr = AST.CallExpr{Caller: caller, Arguments: p.parseArguments()}

	// FIX: check for OpenParenthesis instead of OpenBrace for chained calls like foo()()
	if p.peek().Type == LEX.OpenParenthesis {
		call = p.parseCallExpr(call)
	}
	return call
}

func (p *Parser) parseArguments() []AST.Expr {

	p.expect(LEX.OpenParenthesis, "Error: Missing Opening Parenthesis")

	var arguments []AST.Expr

	if p.peek().Type != LEX.CloseParenthesis {
		arguments = p.parseArgumentsList()
	}

	p.expect(LEX.CloseParenthesis, "Error: Missing Closing Parenthesis")

	return arguments
}

func (p *Parser) parseArgumentsList() []AST.Expr {
	args := []AST.Expr{p.parseAssignmentExpr()}

	for p.peek().Type == LEX.Comma {
		p.consume()
		args = append(args, p.parseAssignmentExpr())
	}

	return args
}

func (p *Parser) parseMemberAndVectorExpr() AST.Expr {
	var object AST.Expr = p.parsePrimaryExpr()

	for p.peek().Type == LEX.Dot || p.peek().Type == LEX.OpenBracket {

		operator := p.consume()

		if operator.Type == LEX.Dot {

			property := p.parsePrimaryExpr()
			computed := false

			if property.Kind() != AST.IdentifierType {
				panic("Error: Property must be an Identifier")
			}

			object = AST.MemberExpr{Object: object, Property: property, Computed: computed}

		}

		if operator.Type == LEX.OpenBracket {
			index := p.parseExpr()

			object = AST.IndexExpr{Array: object, Index: index}

			p.expect(LEX.CloseBracket, "Error: Missing Closing Bracket")
		}

	}

	return object
}

func (p *Parser) parsePrimaryExpr() AST.Expr {

	tk := p.peek().Type
	switch tk {

	case LEX.Identifier:
		return AST.Identifier{Symbol: p.consume().Value}
	case LEX.Number:
		return AST.NumericLiteral{Value: p.parseInt(p.consume().Value)}
	case LEX.String:
		return AST.StringLiteral{Value: p.consume().Value}
	case LEX.OpenBracket:
		return p.parseVectorDeclaration()
	case LEX.OpenParenthesis:
		p.consume()
		expr := p.parseExpr()
		p.expect(LEX.CloseParenthesis, "Error: Missing Closing Parenthesis")
		return expr

	default:
		panic("Error: Invalid Token")
	}
}
