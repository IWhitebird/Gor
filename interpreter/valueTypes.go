package interpreter

import (
	AST "Gor/ast"
)

type ValueType string

const (
	NullType       ValueType = "null"
	NumberType     ValueType = "number"
	StringType     ValueType = "string"
	BoolType       ValueType = "bool"
	ObjectType     ValueType = "object"
	NativeFuncType ValueType = "nativeFunc"
	FunctionType   ValueType = "function"
	ReturnType     ValueType = "return"
)

type RuntimeVal interface {
	Type() ValueType
}

type NullVal struct {
	TypeVal ValueType
	Value   *string
}

func (n NullVal) Type() ValueType {
	return n.TypeVal
}

type NumberVal struct {
	TypeVal ValueType
	Value   int
}

func (n NumberVal) Type() ValueType {
	return n.TypeVal
}

type StringVal struct {
	TypeVal ValueType
	Value   string
}

func (s StringVal) Type() ValueType {
	return s.TypeVal
}

type BoolVal struct {
	TypeVal ValueType
	Value   bool
}

func (b BoolVal) Type() ValueType {
	return b.TypeVal
}

type ObjectVal struct {
	TypeVal    ValueType
	Properties map[string]RuntimeVal
}

func (o ObjectVal) Type() ValueType {
	return o.TypeVal
}

type FunctionCall func(args []RuntimeVal, env *Environment) RuntimeVal

type NativeFuncVal struct {
	TypeVal ValueType
	Call    FunctionCall
}

func (n NativeFuncVal) Type() ValueType {
	return n.TypeVal
}

type FunctionVal struct {
	TypeVal    ValueType
	Name       string
	Parameters []string
	Body       []AST.Stmt
	Env        Environment
}

func (f FunctionVal) Type() ValueType {
	return f.TypeVal
}

type ReturnVal struct {
	TypeVal ValueType
	Value   RuntimeVal
}

func (r ReturnVal) Type() ValueType {
	return r.TypeVal
}

// Instant Make Function

func MK_RETURN(value RuntimeVal) ReturnVal {
	return ReturnVal{
		TypeVal: ReturnType,
		Value:   value,
	}
}

func MK_NULL() NullVal {
	return NullVal{
		TypeVal: NullType,
		Value:   nil,
	}
}

func MK_NUMBER(values ...int) NumberVal {
	var value int

	if len(values) > 0 {
		value = values[0]
	} else {
		// Default value if not provided
		value = 0
	}

	return NumberVal{
		TypeVal: NumberType,
		Value:   value,
	}
}

func MK_STRING(values ...string) StringVal {
	var value string

	if len(values) > 0 {
		value = values[0]
	} else {
		// Default value if not provided
		value = ""
	}

	return StringVal{
		TypeVal: StringType,
		Value:   value,
	}
}

func MK_BOOL(values ...bool) BoolVal {
	var value bool

	if len(values) > 0 {
		value = values[0]
	} else {
		// Default value if not provided
		value = true
	}

	return BoolVal{
		TypeVal: "BoolType",
		Value:   value,
	}
}

func MK_OBJECT(properties ...map[string]RuntimeVal) ObjectVal {
	var value map[string]RuntimeVal

	if len(properties) > 0 {
		value = properties[0]
	} else {
		// Default value if not provided
		value = make(map[string]RuntimeVal)
	}

	return ObjectVal{
		TypeVal:    ObjectType,
		Properties: value,
	}
}

func MK_NATIVE_FUNC(call FunctionCall) NativeFuncVal {
	return NativeFuncVal{
		TypeVal: NativeFuncType,
		Call:    call,
	}
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
	case ReturnType:
		return RuntimeVal_Wrapper(val.(ReturnVal).Value)
	}

	// Return a default value in case the type is not recognized
	return nil
}
