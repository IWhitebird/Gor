package Gor

import (
	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
	PGM "github.com/iwhitebird/Gor/program"
)

func Repl() {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	PGM.Repl(parser, env)
}

type Result struct {
	Output string
	AST    string
	Error  error
}

func RunFromInput(input string) <-chan Result {
	resultChan := make(chan Result)

	go func() {
		defer close(resultChan)
		Output, AST, err := PGM.CompleteInput(input)
		resultChan <- Result{Output, AST, err}
	}()

	return resultChan
}

// func RunFromFile(file_path string) <-chan Result {
// 	resultChan := make(chan Result)

// 	go func() {
// 		defer close(resultChan)

// 		result, err := PGM.CompleteFile(file_path)
// 		resultChan <- Result{result, err}
// 	}()

// 	return resultChan
// }
