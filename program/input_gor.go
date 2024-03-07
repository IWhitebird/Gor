package program

import (
	"encoding/json"
	"fmt"

	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
)

func CompleteInput(input string) (string, error) {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}

	program := parser.ProduceAst(input)

	bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", err
	}

	ITR.Evaluate(program, env)

	return string(bodyJSON), nil
}
