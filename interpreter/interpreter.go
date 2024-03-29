package interpreter

import (
	"fmt"
	"os"

	AST "github.com/iwhitebird/Gor/ast"
)

func Evaluate(astNode AST.Stmt, env *Environment) RuntimeVal {

	switch astNode.Kind() {
	// Expressions

	case AST.NumericLiteralType:
		return NumberVal{TypeVal: NumberType, Value: astNode.(AST.NumericLiteral).Value}
	case AST.StringLiteralType:
		return StringVal{TypeVal: StringType, Value: astNode.(AST.StringLiteral).Value}
	case AST.ReturnStmtType:
		return ReturnVal{TypeVal: ReturnType, Value: Evaluate(astNode.(AST.ReturnStmt).Value, env)}
	case AST.VectorLiteralType:
		return Eval_vector_expr(astNode.(AST.VectorLiteral), env)
	case AST.IndexExprType:
		return Eval_index_expr(astNode.(AST.IndexExpr), env)
	case AST.OjectLiteralType:
		return Eval_object_expr(astNode.(AST.ObjectLiteral), env)
	case AST.MemberExprType:
		return Eval_member_expr(astNode.(AST.MemberExpr), env)
	case AST.CallExprType:
		return Eval_call_expr(astNode.(AST.CallExpr), env)
	case AST.AssignmentExprType:
		return Eval_assignment_expr(astNode.(AST.AssignmentExpr), env)
	case AST.IdentifierType:
		return Eval_identifier(astNode.(AST.Identifier), env)
	case AST.BinaryExprType:
		return Eval_binary_expr(astNode.(AST.BinaryExpr), env)
	case AST.BlockStmtType:
		return Eval_block_statement(astNode.(AST.BlockStmt), env)
	case AST.IfStmtType:
		return Eval_if_statement(astNode.(AST.IfStmt), env)
	case AST.ForStmtType:
		return Eval_for_statement(astNode.(AST.ForStmt), env)

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
