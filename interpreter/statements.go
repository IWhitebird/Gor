package interpreter

import (
	AST "github.com/iwhitebird/Gor/ast"
)

func Eval_program(program AST.Program, env *Environment) RuntimeVal {
	var lastEvaluated RuntimeVal = MK_NULL()

	for _, statement := range program.Body {

		lastEvaluated = Evaluate(statement, env)

		if lastEvaluated.Type() == ReturnType {
			return lastEvaluated.(ReturnVal).Value
		}

	}

	return lastEvaluated
}

func Eval_block_statement(blockStmt AST.BlockStmt, env *Environment) RuntimeVal {
	var lastEvaluated RuntimeVal = MK_NULL()
	for _, statement := range blockStmt.Body {

		// fmt.Println("Evaluating AST Node", statement.Kind())

		lastEvaluated = Evaluate(statement, env)

		if lastEvaluated.Type() == ReturnType {
			return lastEvaluated
		}

	}

	return lastEvaluated
}

func Eval_variable_declaration(variableDeclaration AST.VariableDeclaration, env *Environment) RuntimeVal {
	if variableDeclaration.Value == nil {
		return MK_NULL()
	}
	value := Evaluate(variableDeclaration.Value, env)
	return env.DeclareVar(variableDeclaration.Identifier, value, variableDeclaration.Constant)
}

func Eval_function_declaration(functionDeclaration AST.FunctionDeclaration, env *Environment) RuntimeVal {

	function := FunctionVal{
		TypeVal:    FunctionType,
		Name:       functionDeclaration.Identifier,
		Parameters: functionDeclaration.Parameters,
		Body:       functionDeclaration.Body,
		Env:        *env,
	}

	return env.DeclareVar(functionDeclaration.Identifier, function, false)
}

func Eval_if_statement(declaration AST.IfStmt, env *Environment) RuntimeVal {

	test := Evaluate(declaration.Test, env)

	if test.(BoolVal).Value {
		return Eval_body(declaration.Body, env, true)
	} else if declaration.Alternate != nil {
		return Eval_body(declaration.Alternate, env, true)
	} else {
		return MK_NULL()
	}
}

func Eval_body(body AST.Stmt, env *Environment, newEnv bool) RuntimeVal {
	var scope *Environment

	if newEnv {
		scope = NewEnvironment(env)
	} else {
		scope = env
	}

	var result RuntimeVal = MK_NULL()

	switch body.Kind() {
	case AST.BlockStmtType:
		result = Eval_block_statement(body.(AST.BlockStmt), scope)
	case AST.IfStmtType:
		result = Eval_if_statement(body.(AST.IfStmt), scope)
	}

	return result
}

func Eval_for_statement(declaration AST.ForStmt, env *Environment) RuntimeVal {
	newenv := NewEnvironment(env)

	Eval_variable_declaration(declaration.Init.(AST.VariableDeclaration), newenv)
	test := Evaluate(declaration.Test, newenv)
	update := declaration.Update
	body := declaration.Body

	if !test.(BoolVal).Value {
		return MK_NULL()
	}

	for {

		new_newenv := NewEnvironment(newenv)
		Eval_body(body, new_newenv, false)
		Eval_assignment_expr(update.(AST.AssignmentExpr), newenv)

		test = Evaluate(declaration.Test, newenv)

		if !test.(BoolVal).Value {
			break
		}
	}
	return MK_NULL()
}
