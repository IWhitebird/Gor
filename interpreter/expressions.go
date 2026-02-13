package interpreter

import (
	"fmt"
	"os"

	AST "github.com/iwhitebird/Gor/ast"
)

func Eval_binary_expr(binaryExpr AST.BinaryExpr, env *Environment) RuntimeVal {
	var lhs RuntimeVal = Evaluate(binaryExpr.Left, env)
	var rhs RuntimeVal = Evaluate(binaryExpr.Right, env)

	return Eval_binary_op(binaryExpr.Operator, lhs, rhs)
}

func runtimeValsEqual(lhs RuntimeVal, rhs RuntimeVal) bool {
	if lhs.Type() != rhs.Type() {
		return false
	}
	switch lhs.Type() {
	case NumberType:
		return lhs.(NumberVal).Value == rhs.(NumberVal).Value
	case StringType:
		return lhs.(StringVal).Value == rhs.(StringVal).Value
	case BoolType:
		return lhs.(BoolVal).Value == rhs.(BoolVal).Value
	case NullType:
		return true
	case ObjectType:
		lProps := lhs.(ObjectVal).Properties
		rProps := rhs.(ObjectVal).Properties
		if len(lProps) != len(rProps) {
			return false
		}
		for k, lv := range lProps {
			rv, ok := rProps[k]
			if !ok || !runtimeValsEqual(lv, rv) {
				return false
			}
		}
		return true
	case VectorType:
		lElems := lhs.(VectorVal).Elements
		rElems := rhs.(VectorVal).Elements
		if len(lElems) != len(rElems) {
			return false
		}
		for i := range lElems {
			if !runtimeValsEqual(lElems[i], rElems[i]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func Eval_binary_op(operator string, lhs RuntimeVal, rhs RuntimeVal) RuntimeVal {
	switch operator {

	// Comparison operators — work across types
	case "==":
		return MK_BOOL(runtimeValsEqual(lhs, rhs))
	case "!=":
		return MK_BOOL(!runtimeValsEqual(lhs, rhs))

	// Logical operators
	case "&&":
		if IsTruthy(lhs) {
			return rhs
		}
		return lhs
	case "||":
		if IsTruthy(lhs) {
			return lhs
		}
		return rhs

	// Numeric comparison
	case ">":
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_BOOL(lhs.(NumberVal).Value > rhs.(NumberVal).Value)
		}
		if lhs.Type() == StringType && rhs.Type() == StringType {
			return MK_BOOL(lhs.(StringVal).Value > rhs.(StringVal).Value)
		}
		fmt.Println("Error: '>' requires numbers or strings")
		os.Exit(1)
	case "<":
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_BOOL(lhs.(NumberVal).Value < rhs.(NumberVal).Value)
		}
		if lhs.Type() == StringType && rhs.Type() == StringType {
			return MK_BOOL(lhs.(StringVal).Value < rhs.(StringVal).Value)
		}
		fmt.Println("Error: '<' requires numbers or strings")
		os.Exit(1)
	case ">=":
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_BOOL(lhs.(NumberVal).Value >= rhs.(NumberVal).Value)
		}
		if lhs.Type() == StringType && rhs.Type() == StringType {
			return MK_BOOL(lhs.(StringVal).Value >= rhs.(StringVal).Value)
		}
		fmt.Println("Error: '>=' requires numbers or strings")
		os.Exit(1)
	case "<=":
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_BOOL(lhs.(NumberVal).Value <= rhs.(NumberVal).Value)
		}
		if lhs.Type() == StringType && rhs.Type() == StringType {
			return MK_BOOL(lhs.(StringVal).Value <= rhs.(StringVal).Value)
		}
		fmt.Println("Error: '<=' requires numbers or strings")
		os.Exit(1)

	// Arithmetic — + also handles string concatenation
	case "+":
		if lhs.Type() == NumberType && rhs.Type() == NumberType {
			return MK_NUMBER(lhs.(NumberVal).Value + rhs.(NumberVal).Value)
		}
		if lhs.Type() == StringType && rhs.Type() == StringType {
			return MK_STRING(lhs.(StringVal).Value + rhs.(StringVal).Value)
		}
		fmt.Println("Error: '+' requires two numbers or two strings")
		os.Exit(1)
	case "-":
		return MK_NUMBER(lhs.(NumberVal).Value - rhs.(NumberVal).Value)
	case "*":
		return MK_NUMBER(lhs.(NumberVal).Value * rhs.(NumberVal).Value)
	case "/":
		if rhs.(NumberVal).Value == 0 {
			fmt.Println("Error: Division by zero")
			os.Exit(1)
		}
		return MK_NUMBER(lhs.(NumberVal).Value / rhs.(NumberVal).Value)
	case "%":
		if rhs.(NumberVal).Value == 0 {
			fmt.Println("Error: Modulo by zero")
			os.Exit(1)
		}
		return MK_NUMBER(lhs.(NumberVal).Value % rhs.(NumberVal).Value)

	// Bitwise operators
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

func Eval_identifier(identifier AST.Identifier, env *Environment) RuntimeVal {
	value := env.LookupVar(identifier.Symbol)
	return value
}

func Eval_assignment_expr(assignmentExpr AST.AssignmentExpr, env *Environment) RuntimeVal {

	switch assignmentExpr.Left.Kind() {

	case AST.IdentifierType:
		return env.AssignVar(assignmentExpr.Left.(AST.Identifier).Symbol, Evaluate(assignmentExpr.Right, env))
	case AST.MemberExprType:
		var object RuntimeVal = MK_NULL()
		object = Eval_assign_member_expr(assignmentExpr.Left.(AST.MemberExpr), env)
		object.(ObjectVal).Properties[assignmentExpr.Left.(AST.MemberExpr).Property.(AST.Identifier).Symbol] = Evaluate(assignmentExpr.Right, env)
		return object
	case AST.IndexExprType:
		array := Evaluate(assignmentExpr.Left.(AST.IndexExpr).Array, env)
		index := Evaluate(assignmentExpr.Left.(AST.IndexExpr).Index, env)
		array.(VectorVal).Elements[int(index.(NumberVal).Value)] = Evaluate(assignmentExpr.Right, env)
		return array
	}
	fmt.Println("Error: Invalid Assignment")
	os.Exit(1)
	return MK_NULL()
}

func Eval_object_expr(objectLiteral AST.ObjectLiteral, env *Environment) RuntimeVal {
	properties := make(map[string]RuntimeVal)

	for _, property := range objectLiteral.Properties {
		if property.Value == nil {
			properties[property.Key] = env.LookupVar(property.Key)
			continue
		}
		properties[property.Key] = Evaluate(property.Value, env)
	}

	return ObjectVal{Properties: properties}
}

func Eval_call_expr(callExpr AST.CallExpr, env *Environment) RuntimeVal {
	args := make([]RuntimeVal, 0, len(callExpr.Arguments))

	for _, arg := range callExpr.Arguments {
		evaluatedArg := Evaluate(arg, env)
		args = append(args, evaluatedArg)
	}

	caller := Evaluate(callExpr.Caller, env)

	switch caller := caller.(type) {
	case NativeFuncVal:
		result := caller.Call(args, env)
		return result

	case FunctionVal:
		scope := NewEnvironment(caller.Env)

		if len(caller.Parameters) != len(args) {
			fmt.Println("Error: Expected", len(caller.Parameters), "arguments, got", len(args))
			os.Exit(1)
		}

		for i, param := range caller.Parameters {
			scope.DeclareVar(param, args[i], false)
		}

		result := Evaluate(caller.Body, scope)

		// FIX: handle functions that don't explicitly return a value
		if result.Type() == ReturnType {
			return result.(ReturnVal).Value
		}
		return result
	}
	fmt.Println("Error: Cannot call a non-function value")
	os.Exit(1)
	return MK_NULL()
}

func Eval_member_expr(memberExpr AST.MemberExpr, env *Environment) RuntimeVal {

	object := Evaluate(memberExpr.Object, env)

	if memberExpr.Property.Kind() == AST.IdentifierType {
		return object.(ObjectVal).Properties[memberExpr.Property.(AST.Identifier).Symbol]
	}

	fmt.Println("Error: Invalid Member Expression")
	return object
}

func Eval_assign_member_expr(memberExpr AST.MemberExpr, env *Environment) RuntimeVal {

	object := Evaluate(memberExpr.Object, env)

	if memberExpr.Property.Kind() == AST.IdentifierType {
		return object
	}

	fmt.Println("Error: Invalid Member Expression")
	return object
}

func Eval_vector_expr(vectorLiteral AST.VectorLiteral, env *Environment) RuntimeVal {
	elements := make([]RuntimeVal, 0, len(vectorLiteral.Elements))

	for _, element := range vectorLiteral.Elements {
		evaluatedElement := Evaluate(element, env)
		elements = append(elements, evaluatedElement)
	}

	return VectorVal{Elements: elements}
}

func Eval_index_expr(indexExpr AST.IndexExpr, env *Environment) RuntimeVal {
	array := Evaluate(indexExpr.Array, env)
	index := Evaluate(indexExpr.Index, env)

	if array.Type() == VectorType && index.Type() == NumberType {
		idx := index.(NumberVal).Value
		elems := array.(VectorVal).Elements
		if idx < 0 || idx >= len(elems) {
			fmt.Println("Error: Index out of bounds:", idx)
			os.Exit(1)
		}
		return elems[idx]
	}

	fmt.Println("Error: Invalid Index Expression")
	os.Exit(1)
	return MK_NULL()
}
