package interpreter

import (
	"fmt"
	"os"
)

type Environment struct {
	ParentEnv *Environment
	Variables map[string]RuntimeVal
	Constants map[string]bool
}

func NewEnvironment(parentEnv *Environment) *Environment {
	return &Environment{
		ParentEnv: parentEnv,
		Variables: make(map[string]RuntimeVal),
		Constants: make(map[string]bool),
	}
}

func (env *Environment) DeclareVar(varname string, value RuntimeVal, optionalParams ...bool) RuntimeVal {
	if _, exists := env.Variables[varname]; exists {
		fmt.Println("ERROR : Cannot declare variable, As it already is defined.", varname)
		os.Exit(1)
	}

	isConst := false

	if len(optionalParams) > 0 {
		isConst = optionalParams[0]
	}

	if isConst {
		env.Constants[varname] = true
	}
	env.Variables[varname] = value
	return value
}

func (env *Environment) AssignVar(varname string, value RuntimeVal) RuntimeVal {
	resolvedEnv := env.Resolve(varname)

	// FIX: check Constants on the resolved env, not the current one
	if resolvedEnv.Constants[varname] {
		fmt.Println("ERROR : Cannot assign to constant variable.", varname)
		os.Exit(1)
	}

	resolvedEnv.Variables[varname] = value
	return value
}

func (env *Environment) LookupVar(varname string) RuntimeVal {
	resolvedEnv := env.Resolve(varname)
	return resolvedEnv.Variables[varname]
}

func (env *Environment) Resolve(varname string) *Environment {
	if _, exists := env.Variables[varname]; exists {
		return env
	}

	if env.ParentEnv == nil {
		fmt.Println("ERROR Cannot resolve, as it does not exist.", varname)
		os.Exit(1)
	}

	return env.ParentEnv.Resolve(varname)
}

// IsTruthy converts any RuntimeVal to a boolean for use in if/for conditions.
func IsTruthy(val RuntimeVal) bool {
	switch val.Type() {
	case BoolType:
		return val.(BoolVal).Value
	case NumberType:
		return val.(NumberVal).Value != 0
	case StringType:
		return val.(StringVal).Value != ""
	case NullType:
		return false
	case ObjectType:
		return len(val.(ObjectVal).Properties) > 0
	case VectorType:
		return len(val.(VectorVal).Elements) > 0
	default:
		return true
	}
}

func EnviromentSetup() *Environment {
	// Root Environment Instance
	parentEnv := NewEnvironment(nil)
	// Declare Variables
	parentEnv.DeclareVar("null", MK_NULL())
	parentEnv.DeclareVar("true", MK_BOOL(true))
	parentEnv.DeclareVar("false", MK_BOOL(false))

	parentEnv.DeclareVar("print", MK_NATIVE_FUNC(func(args []RuntimeVal, env *Environment) RuntimeVal {
		// FIX: print ALL arguments, not just the first one (removed early return)
		for _, arg := range args {
			switch arg.Type() {
			case NumberType:
				fmt.Println(arg.(NumberVal).Value)
			case BoolType:
				fmt.Println(arg.(BoolVal).Value)
			case StringType:
				fmt.Println(arg.(StringVal).Value)
			case ObjectType:
				convertedJson := ConvertMapToJson(arg.(ObjectVal).Properties)
				fmt.Println(convertedJson)
			case VectorType:
				fmt.Println(arg.(VectorVal).Elements)
			case NullType:
				fmt.Println("null")
			default:
				fmt.Println(arg)
			}
		}

		if len(args) == 0 {
			fmt.Println("null")
		}

		if len(args) > 0 {
			return args[len(args)-1]
		}
		return MK_NULL()
	},
	))

	parentEnv.DeclareVar("swap", MK_NATIVE_FUNC(func(args []RuntimeVal, env *Environment) RuntimeVal {
		fmt.Println(args)
		return MK_NULL()
	},
	))

	// Environment Instance
	env := NewEnvironment(parentEnv)

	return env
}
