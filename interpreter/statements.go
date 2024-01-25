package interpreter

import (
	AST "Gor/ast"
)

func Eval_program(program AST.Program, env Environment) RuntimeVal {
	var lastEvaluated RuntimeVal = MK_NULL()
	for _, statements := range program.Body {
		lastEvaluated = Evaluate(statements, env)
	}
	return lastEvaluated
}

func Eval_variable_declaration(variableDeclaration AST.VariableDeclaration, env Environment) RuntimeVal {
	if variableDeclaration.Value == nil {
		return MK_NULL()
	}
	value := Evaluate(variableDeclaration.Value, env)
	return env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)
}

func Eval_function_declaration(functionDeclaration AST.FunctionDeclaration, env Environment) RuntimeVal {

	function := FunctionVal{
		TypeVal:    FunctionType,
		Name:       functionDeclaration.Identifier,
		Parameters: functionDeclaration.Parameters,
		Body:       functionDeclaration.Body,
		Env:        env,
	}

	return env.DeclareVar(functionDeclaration.Identifier, function, false)
}
