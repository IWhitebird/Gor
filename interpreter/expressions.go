package interpreter

import (
	AST "Gor/ast"
	"fmt"
	"os"
)

func Eval_binary_expr(binaryExpr AST.BinaryExpr, env Environment) RuntimeVal {
	var lhs RuntimeVal = Evaluate(binaryExpr.Left, env)
	var rhs RuntimeVal = Evaluate(binaryExpr.Right, env)

	if lhs.Type() == NumberType && rhs.Type() == NumberType {
		return Eval_binary_expr_number(binaryExpr.Operator, lhs, rhs)
	}

	return MK_NULL()
}

func Eval_binary_expr_number(operator string, lhs RuntimeVal, rhs RuntimeVal) RuntimeVal {

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

func Eval_identifier(identifier AST.Identifier, env Environment) RuntimeVal {
	value := env.LookupVar(identifier.Symbol)
	return value
}

func Eval_assignment_expr(assignmentExpr AST.AssignmentExpr, env Environment) RuntimeVal {
	if assignmentExpr.Left.Kind() != AST.IdentifierType {
		fmt.Println("Error: Invalid Assignment")
		os.Exit(1)
	}

	varName := assignmentExpr.Left.(AST.Identifier).Symbol

	return env.AssignVar(varName, Evaluate(assignmentExpr.Right, env))
}

func Eval_object_expr(objectLiteral AST.ObjectLiteral, env Environment) RuntimeVal {
	properties := make(map[string]RuntimeVal)

	for _, property := range objectLiteral.Properties {
		if property.Value == nil {
			properties[property.Key] = env.LookupVar(property.Key)
			continue
		}
		properties[property.Key] = Evaluate(property.Value, env)
	}

	return ObjectVal{TypeVal: ObjectType, Properties: properties}
}

func Eval_call_expr(callExpr AST.CallExpr, env Environment) RuntimeVal {
	var args []RuntimeVal

	for _, arg := range callExpr.Arguments {
		evaluatedArg := Evaluate(arg, env)
		args = append(args, evaluatedArg)
	}

	var function = Evaluate(callExpr.Caller, env).(NativeFuncVal)

	if function.Type() != NativeFuncType {
		fmt.Println("Error: Not a function")
		os.Exit(1)
	}

	result := function.Call(args, &env)

	return result
}
