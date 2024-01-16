package ast

/*
NodeType represents the type of a node.
*/
type NodeType string

const (

	// Statements
	ProgramType             NodeType = "Program"
	VariableDeclarationType NodeType = "VariableDeclaration"

	// Expressions
	AssignmentExprType NodeType = "AssignmentExpr"
	NumericLiteralType NodeType = "NumericLiteral"
	IdentifierType     NodeType = "Identifier"
	BinaryExprType     NodeType = "BinaryExpr"
)

/*
Does not contain Value at RunTime
*/
type Stmt interface {
	Kind() NodeType
}

/*
Program contains Many Statements
*/
type Program struct {
	KindValue NodeType
	Body      []Stmt
}

func (p Program) Kind() NodeType {
	return p.KindValue
}

/*
VariableDeclaration represents a variable declaration in the source.
*/
type VariableDeclaration struct {
	KindValue  NodeType
	Constant   bool
	Identifier string
	Value      Expr
}

func (v VariableDeclaration) Kind() NodeType {
	return v.KindValue
}

/*
Expression will result in value at runtime.
*/
type Expr interface {
	Stmt
}

/*
AssignmentExpr represents an assignment operation.
*/

type AssignmentExpr struct {
	KindValue NodeType
	Left      Expr
	Right     Expr
}

func (a AssignmentExpr) Kind() NodeType {
	return a.KindValue
}

/*
BinaryExpr represents an operation with two sides separated by an operator.
*/
type BinaryExpr struct {
	KindValue NodeType
	Left      Expr
	Right     Expr
	Operator  string // needs to be of type BinaryOperator
}

func (b BinaryExpr) Kind() NodeType {
	return b.KindValue
}

/*
Identifier represents a user-defined variable or symbol in the source.
*/
type Identifier struct {
	KindValue NodeType
	Symbol    string
}

func (i Identifier) Kind() NodeType {
	return i.KindValue
}

/*
NumericLiteral represents a numeric constant inside the source code.
*/
type NumericLiteral struct {
	KindValue NodeType
	Value     int
}

func (n NumericLiteral) Kind() NodeType {
	return n.KindValue
}

func Main() {
	// fmt.Println(">> AST Running >>")
}
