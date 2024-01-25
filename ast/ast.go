package ast

/*
NodeType represents the type of a node.
*/
type NodeType string

const (

	// Statements
	ProgramType             NodeType = "Program"
	VariableDeclarationType NodeType = "VariableDeclaration"
	FunctionDeclarationType NodeType = "FunctionDeclaration"
	// Expressions
	AssignmentExprType NodeType = "AssignmentExpr"
	MemberExprType     NodeType = "MemberExpr"
	CallExprType       NodeType = "CallExpr"

	// Literals
	OjectLiteralType   NodeType = "ObjectLiteral"
	PropertyType       NodeType = "Property"
	NumericLiteralType NodeType = "NumericLiteral"
	IdentifierType     NodeType = "Identifier"
	BinaryExprType     NodeType = "BinaryExpr"
	BlockStmtType      NodeType = "BlockStmt"
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
FunctionDeclaration represents a function declaration in the source.
*/

type FunctionDeclaration struct {
	KindValue  NodeType
	Identifier string
	Parameters []string
	Body       BlockStmt
}

func (f FunctionDeclaration) Kind() NodeType {
	return f.KindValue
}

/*
BlockStmt represents a block statement in the source.
*/
type BlockStmt struct {
	KindValue NodeType
	Body      []Stmt
}

func (b BlockStmt) Kind() NodeType {
	return b.KindValue
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

/*
ObjectLiteral represents an object literal inside the source code.
*/

type ObjectLiteral struct {
	KindValue  NodeType
	Properties []Property
}

func (o ObjectLiteral) Kind() NodeType {
	return o.KindValue
}

/*
Property represents a property inside an object literal.
*/

type Property struct {
	KindValue NodeType
	Key       string
	Value     Expr
}

func (p Property) Kind() NodeType {
	return p.KindValue
}

/*
MemberExpr represents a member expression inside the source code.
*/

type MemberExpr struct {
	KindValue NodeType
	Object    Expr
	Property  Expr
	Computed  bool
}

func (m MemberExpr) Kind() NodeType {
	return m.KindValue
}

/*
CallExpr represents a function call inside the source code.
*/

type CallExpr struct {
	KindValue NodeType
	Caller    Expr
	Arguments []Expr
}

func (c CallExpr) Kind() NodeType {
	return c.KindValue
}
