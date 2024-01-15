package interpreter

import (
	"fmt"
)

type Environment struct {
	ParentEnv *Environment
	Variables map[string]RuntimeVal
}

func NewEnvironment(parentEnv *Environment) *Environment {
	return &Environment{
		ParentEnv: parentEnv,
		Variables: make(map[string]RuntimeVal),
	}
}
func (env *Environment) DeclareVar(varname string, value RuntimeVal) RuntimeVal {
	if _, exists := env.Variables[varname]; exists {
		fmt.Println("ERROR : Cannot declare variable, As it already is defined.", varname)
	}

	env.Variables[varname] = value
	return value
}

func (env *Environment) AssignVar(varname string, value RuntimeVal) RuntimeVal {
	resolvedEnv := env.Resolve(varname)
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
