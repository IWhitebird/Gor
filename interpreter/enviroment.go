package interpreter

import (
	"fmt"
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

	if env.Constants[varname] {
		fmt.Println("ERROR : Cannot assign to constant variable.", varname)
		return nil
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
		fmt.Println("ERROR Cannot resolve ,as it does not exist.", varname)
	}

	return env.ParentEnv.Resolve(varname)
}

func EnviromentSetup() *Environment {
	// Root Environment Instance
	parentEnv := NewEnvironment(nil)
	// Declare Variables
	parentEnv.DeclareVar("a", MK_NUMBER(10))
	parentEnv.DeclareVar("b", MK_NUMBER(20))
	parentEnv.DeclareVar("null", MK_NULL())
	parentEnv.DeclareVar("true", MK_BOOL(true))
	parentEnv.DeclareVar("false", MK_BOOL(false))

	// Environment Instance
	env := NewEnvironment(parentEnv)

	return env
}
