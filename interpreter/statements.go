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
}
