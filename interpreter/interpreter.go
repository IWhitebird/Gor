package interpreter

import (
	AST "Gor/ast"
	"fmt"
	"os"
)

func Evaluate(astNode AST.Stmt, env Environment) RuntimeVal {
	switch astNode.Kind() {

	// Expressions

	case AST.NumericLiteralType:
		return NumberVal{TypeVal: NumberType, Value: astNode.(AST.NumericLiteral).Value}
	case AST.OjectLiteralType:
		return Eval_object_expr(astNode.(AST.ObjectLiteral), env)
	case AST.AssignmentExprType:
		return Eval_assignment_expr(astNode.(AST.AssignmentExpr), env)
	case AST.IdentifierType:
		return Eval_identifier(astNode.(AST.Identifier), env)
	case AST.BinaryExprType:
		return Eval_binary_expr(astNode.(AST.BinaryExpr), env)

	// Statements

	case AST.ProgramType:
		return Eval_program(astNode.(AST.Program), env)
	case AST.VariableDeclarationType:
		return Eval_variable_declaration(astNode.(AST.VariableDeclaration), env)

	// Default
	default:
		fmt.Println("Unknown AST Node Type", astNode)
		os.Exit(1)
	}
	return nil
}
