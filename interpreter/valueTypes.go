package interpreter

type ValueType string

const (
	NullType       ValueType = "null"
	NumberType     ValueType = "number"
	BoolType       ValueType = "bool"
	ObjectType     ValueType = "object"
	NativeFuncType ValueType = "nativeFunc"
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

// Instant Make Function

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

func MK_NATIVE_FUNC(call FunctionCall) NativeFuncVal {
	return NativeFuncVal{
		TypeVal: NativeFuncType,
		Call:    call,
	}
}
