package interpreter

type ValueType string

const (
	NullType   ValueType = "null"
	NumberType ValueType = "number"
)

type RuntimeVal interface {
	Type() ValueType
}

type NullVal struct {
	TypeVal ValueType
	Value   string
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
