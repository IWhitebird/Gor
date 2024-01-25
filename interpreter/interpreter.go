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
	case AST.CallExprType:
		return Eval_call_expr(astNode.(AST.CallExpr), env)
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
	case AST.FunctionDeclarationType:
		return Eval_function_declaration(astNode.(AST.FunctionDeclaration), env)

	// Default
	default:
		fmt.Println("Unknown AST Node Type", astNode)
		os.Exit(1)
	}
	return nil
}
