package ast

import "encoding/json"

/*
NodeType represents the type of a node.
*/
type NodeType int

const (
	// Statements
	ProgramType NodeType = iota
	VariableDeclarationType
	FunctionDeclarationType
	BlockStmtType
	IfStmtType
	ForStmtType
	ReturnStmtType

	// Expressions
	IndexExprType
	AssignmentExprType
	MemberExprType
	CallExprType

	// Literals
	VectorLiteralType
	OjectLiteralType
	PropertyType
	NumericLiteralType
	StringLiteralType
	IdentifierType
	BinaryExprType
)

var nodeTypeNames = [...]string{
	ProgramType:             "Program",
	VariableDeclarationType: "VariableDeclaration",
	FunctionDeclarationType: "FunctionDeclaration",
	BlockStmtType:           "BlockStmt",
	IfStmtType:              "IfStmt",
	ForStmtType:             "ForStmt",
	ReturnStmtType:          "ReturnStmt",
	IndexExprType:           "IndexExpr",
	AssignmentExprType:      "AssignmentExpr",
	MemberExprType:          "MemberExpr",
	CallExprType:            "CallExpr",
	VectorLiteralType:       "VectorLiteral",
	OjectLiteralType:        "ObjectLiteral",
	PropertyType:            "Property",
	NumericLiteralType:      "NumericLiteral",
	StringLiteralType:       "StringLiteral",
	IdentifierType:          "Identifier",
	BinaryExprType:          "BinaryExpr",
}

func (n NodeType) String() string {
	if int(n) < len(nodeTypeNames) {
		return nodeTypeNames[n]
	}
	return "Unknown"
}

func (n NodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

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
	Body []Stmt
}

func (p Program) Kind() NodeType {
	return ProgramType
}

func (p Program) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind NodeType `json:"kind"`
		Body []Stmt   `json:"body"`
	}{ProgramType, p.Body})
}

/*
BlockStmt represents a block of statements inside the source code.
*/

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) Kind() NodeType {
	return BlockStmtType
}

func (b BlockStmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind NodeType `json:"kind"`
		Body []Stmt   `json:"body"`
	}{BlockStmtType, b.Body})
}

/*
VariableDeclaration represents a variable declaration in the source.
*/
type VariableDeclaration struct {
	Constant   bool
	Identifier string
	Value      Expr
}

func (v VariableDeclaration) Kind() NodeType {
	return VariableDeclarationType
}

func (v VariableDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind       NodeType `json:"kind"`
		Constant   bool     `json:"constant"`
		Identifier string   `json:"identifier"`
		Value      Expr     `json:"value"`
	}{VariableDeclarationType, v.Constant, v.Identifier, v.Value})
}

/*
FunctionDeclaration represents a function declaration in the source.
*/

type FunctionDeclaration struct {
	Identifier string
	Parameters []string
	Body       BlockStmt
}

func (f FunctionDeclaration) Kind() NodeType {
	return FunctionDeclarationType
}

func (f FunctionDeclaration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind       NodeType  `json:"kind"`
		Identifier string    `json:"identifier"`
		Parameters []string  `json:"parameters"`
		Body       BlockStmt `json:"body"`
	}{FunctionDeclarationType, f.Identifier, f.Parameters, f.Body})
}

/*
Return represents a return statement in the source.
*/

type ReturnStmt struct {
	Value Expr
}

func (r ReturnStmt) Kind() NodeType {
	return ReturnStmtType
}

func (r ReturnStmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  NodeType `json:"kind"`
		Value Expr     `json:"value"`
	}{ReturnStmtType, r.Value})
}

/*
IfStmt represents an if statement in the source.
*/
type IfStmt struct {
	Test      Expr
	Body      BlockStmt
	Alternate Stmt
}

func (i IfStmt) Kind() NodeType {
	return IfStmtType
}

func (i IfStmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind      NodeType  `json:"kind"`
		Test      Expr      `json:"test"`
		Body      BlockStmt `json:"body"`
		Alternate Stmt      `json:"alternate"`
	}{IfStmtType, i.Test, i.Body, i.Alternate})
}

/*
ForStmt represents a for statement in the source.
*/

type ForStmt struct {
	Init   Expr
	Test   Expr
	Update Expr
	Body   BlockStmt
}

func (f ForStmt) Kind() NodeType {
	return ForStmtType
}

func (f ForStmt) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind   NodeType  `json:"kind"`
		Init   Expr      `json:"init"`
		Test   Expr      `json:"test"`
		Update Expr      `json:"update"`
		Body   BlockStmt `json:"body"`
	}{ForStmtType, f.Init, f.Test, f.Update, f.Body})
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
	Left  Expr
	Right Expr
}

func (a AssignmentExpr) Kind() NodeType {
	return AssignmentExprType
}

func (a AssignmentExpr) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  NodeType `json:"kind"`
		Left  Expr     `json:"left"`
		Right Expr     `json:"right"`
	}{AssignmentExprType, a.Left, a.Right})
}

/*
BinaryExpr represents an operation with two sides separated by an operator.
*/
type BinaryExpr struct {
	Left     Expr
	Right    Expr
	Operator string
}

func (b BinaryExpr) Kind() NodeType {
	return BinaryExprType
}

func (b BinaryExpr) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind     NodeType `json:"kind"`
		Operator string   `json:"operator"`
		Left     Expr     `json:"left"`
		Right    Expr     `json:"right"`
	}{BinaryExprType, b.Operator, b.Left, b.Right})
}

/*
Identifier represents a user-defined variable or symbol in the source.
*/
type Identifier struct {
	Symbol string
}

func (i Identifier) Kind() NodeType {
	return IdentifierType
}

func (i Identifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind   NodeType `json:"kind"`
		Symbol string   `json:"symbol"`
	}{IdentifierType, i.Symbol})
}

/*
NumericLiteral represents a numeric constant inside the source code.
*/
type NumericLiteral struct {
	Value int
}

func (n NumericLiteral) Kind() NodeType {
	return NumericLiteralType
}

func (n NumericLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  NodeType `json:"kind"`
		Value int      `json:"value"`
	}{NumericLiteralType, n.Value})
}

/*
StringLiteral represents a string constant inside the source code.
*/
type StringLiteral struct {
	Value string
}

func (s StringLiteral) Kind() NodeType {
	return StringLiteralType
}

func (s StringLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  NodeType `json:"kind"`
		Value string   `json:"value"`
	}{StringLiteralType, s.Value})
}

/*
ObjectLiteral represents an object literal inside the source code.
*/

type ObjectLiteral struct {
	Properties []Property
}

func (o ObjectLiteral) Kind() NodeType {
	return OjectLiteralType
}

func (o ObjectLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind       NodeType   `json:"kind"`
		Properties []Property `json:"properties"`
	}{OjectLiteralType, o.Properties})
}

/*
Property represents a property inside an object literal.
*/

type Property struct {
	Key   string
	Value Expr
}

func (p Property) Kind() NodeType {
	return PropertyType
}

func (p Property) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  NodeType `json:"kind"`
		Key   string   `json:"key"`
		Value Expr     `json:"value"`
	}{PropertyType, p.Key, p.Value})
}

/*
MemberExpr represents a member expression inside the source code.
*/

type MemberExpr struct {
	Object   Expr
	Property Expr
	Computed bool
}

func (m MemberExpr) Kind() NodeType {
	return MemberExprType
}

func (m MemberExpr) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind     NodeType `json:"kind"`
		Object   Expr     `json:"object"`
		Property Expr     `json:"property"`
		Computed bool     `json:"computed"`
	}{MemberExprType, m.Object, m.Property, m.Computed})
}

/*
CallExpr represents a function call inside the source code.
*/

type CallExpr struct {
	Caller    Expr
	Arguments []Expr
}

func (c CallExpr) Kind() NodeType {
	return CallExprType
}

func (c CallExpr) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind      NodeType `json:"kind"`
		Caller    Expr     `json:"caller"`
		Arguments []Expr   `json:"arguments"`
	}{CallExprType, c.Caller, c.Arguments})
}

/*
VectorLiteral represents a vector literal inside the source code.
*/

type VectorLiteral struct {
	Elements []Expr
}

func (vc VectorLiteral) Kind() NodeType {
	return VectorLiteralType
}

func (vc VectorLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind     NodeType `json:"kind"`
		Elements []Expr   `json:"elements"`
	}{VectorLiteralType, vc.Elements})
}

/*
IndexExpr represents an index expression inside the source code.
*/

type IndexExpr struct {
	Array Expr
	Index Expr
}

func (ie IndexExpr) Kind() NodeType {
	return IndexExprType
}

func (ie IndexExpr) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  NodeType `json:"kind"`
		Array Expr     `json:"array"`
		Index Expr     `json:"index"`
	}{IndexExprType, ie.Array, ie.Index})
}
