package Gor

import (
	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
	PGM "github.com/iwhitebird/Gor/program"
)

type Result struct {
	Result string
	Error  error
}

func Repl() {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	PGM.Repl(parser, env)
}

func RunFromInput(input string) <-chan Result {
	resultChan := make(chan Result)

	go func() {
		defer close(resultChan)

		result, err := PGM.CompleteInput(input)
		resultChan <- Result{result, err}
	}()

	return resultChan
}

func RunFromFile(file_path string) <-chan Result {
	resultChan := make(chan Result)

	go func() {
		defer close(resultChan)

		result, err := PGM.CompleteFile(file_path)
		resultChan <- Result{result, err}
	}()

	return resultChan
}
