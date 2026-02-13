package Gor

import (
	"encoding/json"
	"fmt"
	"os"

	ITR "github.com/iwhitebird/Gor/core/interpreter"
	PSR "github.com/iwhitebird/Gor/core/parser"
	PGM "github.com/iwhitebird/Gor/core/program"
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

func RunFromFile(filePath string) {
	inputFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not read file %q: %v\n", filePath, err)
		os.Exit(1)
	}

	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}

	program := parser.ProduceAst(string(inputFile))
	ITR.Evaluate(program, env)
}

func PrintAST(filePath string) {
	inputFile, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not read file %q: %v\n", filePath, err)
		os.Exit(1)
	}

	var parser = PSR.Parser{}
	program := parser.ProduceAst(string(inputFile))

	bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling AST: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(bodyJSON))
}
