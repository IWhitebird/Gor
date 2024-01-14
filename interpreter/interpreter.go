package interpreter

import (
	AST "Gor/ast"
	"fmt"
	"os"
)

func Eval_program(program AST.Program) RuntimeVal {
	var lastEvaluated RuntimeVal = NullVal{TypeVal: NullType, Value: "null"}
	for _, statements := range program.Body {
		lastEvaluated = evaluate(statements)
	}
	return lastEvaluated
}

func eval_binary_expr(binaryExpr AST.BinaryExpr) RuntimeVal {
	var lhs RuntimeVal = evaluate(binaryExpr.Left)
	var rhs RuntimeVal = evaluate(binaryExpr.Right)

	if lhs.Type() == NumberType && rhs.Type() == NumberType {
		return eval_binary_expr_number(binaryExpr.Operator, lhs, rhs)
	}

	return NullVal{TypeVal: NullType, Value: "null"}
}

func eval_binary_expr_number(operator string, lhs RuntimeVal, rhs RuntimeVal) RuntimeVal {

	var value int

	switch operator {
	case "+":
		value = lhs.(NumberVal).Value + rhs.(NumberVal).Value
	case "-":
		value = lhs.(NumberVal).Value - rhs.(NumberVal).Value
	case "*":
		value = lhs.(NumberVal).Value * rhs.(NumberVal).Value
	case "/":
		if rhs.(NumberVal).Value == 0 {
			fmt.Println("Error : Division by zero")
			os.Exit(1)
		}
		value = lhs.(NumberVal).Value / rhs.(NumberVal).Value
	case "%":
		value = lhs.(NumberVal).Value % rhs.(NumberVal).Value
	}

	return NumberVal{TypeVal: NumberType, Value: value}

}

func evaluate(astNode AST.Stmt) RuntimeVal {
	switch astNode.Kind() {
	case AST.NumericLiteralType:
		return NumberVal{TypeVal: NumberType, Value: astNode.(AST.NumericLiteral).Value}
	case AST.NullLiretalType:
		return NullVal{TypeVal: NullType, Value: astNode.(AST.NullLiteral).Value}
	case AST.BinaryExprType:
		return eval_binary_expr(astNode.(AST.BinaryExpr))
	case AST.ProgramType:
		return Eval_program(astNode.(AST.Program))
	default:
		fmt.Println("Unknown AST Node Type", astNode)
		os.Exit(1)
	}
	return nil
}
