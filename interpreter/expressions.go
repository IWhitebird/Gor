package interpreter

import (
	AST "Gor/ast"
	"fmt"
	"os"
	"reflect"
)

func Eval_binary_expr(binaryExpr AST.BinaryExpr, env Environment) RuntimeVal {
	var lhs RuntimeVal = Evaluate(binaryExpr.Left, env)

	var rhs RuntimeVal = Evaluate(binaryExpr.Right, env)

	return Eval_binary_expr_number(binaryExpr.Operator, lhs, rhs)
}

func Eval_equals(lhs RuntimeVal, rhs RuntimeVal, operator string) RuntimeVal {

	if lhs.Type() == NumberType && rhs.Type() == NumberType {
		switch operator {
		case "==":
			return MK_BOOL(lhs.(NumberVal).Value == rhs.(NumberVal).Value)
		case "!=":
			return MK_BOOL(lhs.(NumberVal).Value != rhs.(NumberVal).Value)
		case ">=":
			return MK_BOOL(lhs.(NumberVal).Value >= rhs.(NumberVal).Value)
		case "<=":
			return MK_BOOL(lhs.(NumberVal).Value <= rhs.(NumberVal).Value)
		case "&&":
			return MK_NUMBER(rhs.(NumberVal).Value)
		case "||":
			return MK_NUMBER(lhs.(NumberVal).Value)
		}
	}

	if lhs.Type() == BoolType && rhs.Type() == BoolType {
		switch operator {
		case "==":
			return MK_BOOL(lhs.(BoolVal).Value == rhs.(BoolVal).Value)
		case "!=":
			return MK_BOOL(lhs.(BoolVal).Value != rhs.(BoolVal).Value)
		case ">=":
			var lef int
			if lhs.(BoolVal).Value {
				lef = 1
			} else {
				lef = 0
			}
			var rig int
			if rhs.(BoolVal).Value {
				rig = 1
			} else {
				rig = 0
			}
			return MK_BOOL(lef >= rig)
		case "<=":
			var lef int
			if lhs.(BoolVal).Value {
				lef = 1
			} else {
				lef = 0
			}
			var rig int
			if rhs.(BoolVal).Value {
				rig = 1
			} else {
				rig = 0
			}
			return MK_BOOL(lef <= rig)
		case "&&":
			return MK_BOOL(lhs.(BoolVal).Value && rhs.(BoolVal).Value)
		case "||":
			return MK_BOOL(lhs.(BoolVal).Value || rhs.(BoolVal).Value)
		}
	}

	if lhs.Type() == StringType && rhs.Type() == StringType {
		switch operator {
		case "==":
			return MK_BOOL(lhs.(StringVal).Value == rhs.(StringVal).Value)
		case "!=":
			return MK_BOOL(lhs.(StringVal).Value != rhs.(StringVal).Value)
		case ">=":
			return MK_BOOL(lhs.(StringVal).Value >= rhs.(StringVal).Value)
		case "<=":
			return MK_BOOL(lhs.(StringVal).Value <= rhs.(StringVal).Value)
		case "&&":
			return MK_STRING(rhs.(StringVal).Value)
		case "||":
			return MK_STRING(lhs.(StringVal).Value)
		}
	}

	if lhs.Type() == ObjectType && rhs.Type() == ObjectType {
		switch operator {
		case "==":
			return MK_BOOL(reflect.DeepEqual(lhs.(ObjectVal).Properties, rhs.(ObjectVal).Properties))
		case "!=":
			return MK_BOOL(!reflect.DeepEqual(lhs.(ObjectVal).Properties, rhs.(ObjectVal).Properties))
		case ">=":
			leftSize := len(lhs.(ObjectVal).Properties)
			rightSize := len(rhs.(ObjectVal).Properties)
			return MK_BOOL(leftSize >= rightSize)
		case "<=":
			leftSize := len(lhs.(ObjectVal).Properties)
			rightSize := len(rhs.(ObjectVal).Properties)
			return MK_BOOL(leftSize <= rightSize)
		case "&&":
			return MK_OBJECT(rhs.(ObjectVal).Properties)
		case "||":
			return MK_OBJECT(lhs.(ObjectVal).Properties)
		}
	}

	if lhs.Type() == FunctionType && rhs.Type() == FunctionType {
		switch operator {
		case "==":
			return MK_BOOL(lhs.(FunctionVal).Name == rhs.(FunctionVal).Name)
		case "!=":
			return MK_BOOL(lhs.(FunctionVal).Name != rhs.(FunctionVal).Name)
		case ">=":
			return MK_BOOL(lhs.(FunctionVal).Name >= rhs.(FunctionVal).Name)
		case "<=":
			return MK_BOOL(lhs.(FunctionVal).Name <= rhs.(FunctionVal).Name)
		case "&&":
			return MK_NATIVE_FUNC(rhs.(NativeFuncVal).Call)
		case "||":
			return MK_NATIVE_FUNC(lhs.(NativeFuncVal).Call)
		}
	}

	return MK_NULL()
}

func Eval_binary_expr_number(operator string, lhs RuntimeVal, rhs RuntimeVal) RuntimeVal {

	switch operator {
	case "==":
		return Eval_equals(lhs, rhs, operator)
	case "!=":
		return Eval_equals(lhs, rhs, operator)
	case "&&":
		return Eval_equals(lhs, rhs, operator)
	case "||":
		return Eval_equals(lhs, rhs, operator)
	case ">=":
		return Eval_equals(lhs, rhs, operator)
	case "<=":
		return Eval_equals(lhs, rhs, operator)
	}

	switch operator {
	case "+":
		return MK_NUMBER(lhs.(NumberVal).Value + rhs.(NumberVal).Value)
	case "-":
		return MK_NUMBER(lhs.(NumberVal).Value - rhs.(NumberVal).Value)
	case "*":
		return MK_NUMBER(lhs.(NumberVal).Value * rhs.(NumberVal).Value)
	case "/":
		return MK_NUMBER(lhs.(NumberVal).Value / rhs.(NumberVal).Value)
	case "%":
		return MK_NUMBER(lhs.(NumberVal).Value % rhs.(NumberVal).Value)
	case ">":
		return MK_BOOL(lhs.(NumberVal).Value > rhs.(NumberVal).Value)
	case "<":
		return MK_BOOL(lhs.(NumberVal).Value < rhs.(NumberVal).Value)
	case "&":
		if lhs.Type() == BoolType && rhs.Type() == BoolType {
			return MK_BOOL(lhs.(BoolVal).Value && rhs.(BoolVal).Value)
		}
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_NUMBER(lhs.(NumberVal).Value & rhs.(NumberVal).Value)
		}
	case "|":
		if lhs.Type() == BoolType && rhs.Type() == BoolType {
			return MK_BOOL(lhs.(BoolVal).Value || rhs.(BoolVal).Value)
		}
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_NUMBER(lhs.(NumberVal).Value | rhs.(NumberVal).Value)
		}
	}

	return MK_NULL()
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

	caller := Evaluate(callExpr.Caller, env)

	switch caller := caller.(type) {
	case NativeFuncVal:
		if caller.Type() == NativeFuncType {
			result := caller.Call(args, &env)
			return result
		}

	case FunctionVal:
		scope := NewEnvironment(&env)

		if len(caller.Parameters) != len(args) {
			return MK_NULL()
		}

		for i, param := range caller.Parameters {
			scope.DeclareVar(param, args[i], false)
		}

		var result RuntimeVal = MK_NULL()

		for _, statement := range caller.Body {
			result = Evaluate(statement, *scope)
		}

		return result
	}
	return MK_NULL()
}
