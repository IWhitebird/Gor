package interpreter

import (
	"encoding/json"

	AST "github.com/iwhitebird/Gor/core/ast"
)

type ValueType int

const (
	NullType ValueType = iota
	NumberType
	StringType
	BoolType
	ObjectType
	VectorType
	NativeFuncType
	FunctionType
	ReturnType
)

type RuntimeVal interface {
	Type() ValueType
}

type NullVal struct {
	Value *string
}

func (n NullVal) Type() ValueType {
	return NullType
}

type NumberVal struct {
	Value int
}

func (n NumberVal) Type() ValueType {
	return NumberType
}

type StringVal struct {
	Value string
}

func (s StringVal) Type() ValueType {
	return StringType
}

type BoolVal struct {
	Value bool
}

func (b BoolVal) Type() ValueType {
	return BoolType
}

type VectorVal struct {
	Elements []RuntimeVal
}

func (v VectorVal) Type() ValueType {
	return VectorType
}

type ObjectVal struct {
	Properties map[string]RuntimeVal
}

func (o ObjectVal) Type() ValueType {
	return ObjectType
}

type FunctionCall func(args []RuntimeVal, env *Environment) RuntimeVal

type NativeFuncVal struct {
	Call FunctionCall
}

func (n NativeFuncVal) Type() ValueType {
	return NativeFuncType
}

type FunctionVal struct {
	Name       string
	Parameters []string
	Body       AST.BlockStmt
	Env        *Environment // FIX: pointer instead of value copy to avoid aliased maps
}

func (f FunctionVal) Type() ValueType {
	return FunctionType
}

type ReturnVal struct {
	Value RuntimeVal
}

func (r ReturnVal) Type() ValueType {
	return ReturnType
}

// Cached singletons to avoid repeated allocations
var nullSingleton = NullVal{Value: nil}
var trueVal = BoolVal{Value: true}
var falseVal = BoolVal{Value: false}

func MK_RETURN(value RuntimeVal) ReturnVal {
	return ReturnVal{Value: value}
}

func MK_NULL() NullVal {
	return nullSingleton
}

func MK_NUMBER(value int) NumberVal {
	return NumberVal{Value: value}
}

func MK_STRING(value string) StringVal {
	return StringVal{Value: value}
}

func MK_BOOL(value bool) BoolVal {
	if value {
		return trueVal
	}
	return falseVal
}

func MK_OBJECT(properties ...map[string]RuntimeVal) ObjectVal {
	var value map[string]RuntimeVal

	if len(properties) > 0 {
		value = properties[0]
	} else {
		value = make(map[string]RuntimeVal)
	}

	return ObjectVal{Properties: value}
}

func MK_NATIVE_FUNC(call FunctionCall) NativeFuncVal {
	return NativeFuncVal{Call: call}
}

func RuntimeVal_Wrapper(val RuntimeVal) interface{} {
	switch val.Type() {
	case NullType:
		return "null"
	case NumberType:
		return val.(NumberVal).Value
	case StringType:
		return val.(StringVal).Value
	case BoolType:
		return val.(BoolVal).Value
	case ObjectType:
		return val.(ObjectVal).Properties
	}

	return nil
}

func ConvertRuntimeValToMap(inputMap map[string]RuntimeVal) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range inputMap {
		if value.Type() == ObjectType {
			newMap[key] = ConvertRuntimeValToMap(value.(ObjectVal).Properties)
		} else {
			newMap[key] = RuntimeVal_Wrapper(value)
		}
	}
	return newMap
}

func ConvertMapToJson(inputMap map[string]RuntimeVal) string {
	newMap := ConvertRuntimeValToMap(inputMap)
	jsonData, _ := json.Marshal(newMap)
	return string(jsonData)
}
