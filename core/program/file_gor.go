package program

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	ITR "github.com/iwhitebird/Gor/core/interpreter"
	PSR "github.com/iwhitebird/Gor/core/parser"
)

func CompleteFile(file_path string) (string, error) {

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
		return "", err
	}

	ITR.Evaluate(program, env)

	return string(bodyJSON), nil
}
