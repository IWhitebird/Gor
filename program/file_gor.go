package program

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
)

func CompleteFile(file_path string) ([2]interface{}, error) {

	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}

	inputFile, inputFileErr := os.ReadFile(file_path)
	if inputFileErr != nil {
		log.Fatal(inputFileErr)
	}

	program := parser.ProduceAst(string(inputFile))

	bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return [2]interface{}{}, err
	}

	data := ITR.Evaluate(program, env)

	return [2]interface{}{data, string(bodyJSON)}, nil
}
