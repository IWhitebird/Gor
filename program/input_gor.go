package program

import (
	"encoding/json"
	"fmt"

	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
)

func CompleteInput(input string) ([2]interface{}, error) {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	program := parser.ProduceAst(input)

	bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return [2]interface{}{}, err
	}

	evaluatedProgram := ITR.Evaluate(program, env)

	data := [2]interface{}{evaluatedProgram, string(bodyJSON)}

	return data, nil
}
