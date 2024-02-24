package program

import (
	ITR "Gor/interpreter"
	PSR "Gor/parser"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CompleteFile(parser PSR.Parser, env *ITR.Environment) {
	var scanner = bufio.NewScanner(os.Stdin)
	//AST FILE
	outputFile, OutputFileErr := os.Create("Ast.json")
	if OutputFileErr != nil {
		fmt.Println("Error creating file:", OutputFileErr)
		return
	}
	defer outputFile.Close()

	//From File

	inputFile, inputFileErr := os.ReadFile("input.gor")
	if inputFileErr != nil {
		log.Fatal(inputFileErr)
	}

	program := parser.ProduceAst(string(inputFile))

	bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to the file
	if OutputFileErr == nil {
		_, writeErr := outputFile.Write(bodyJSON)
		if writeErr != nil {
			fmt.Println("Error writing to file:", writeErr)
		}
	} else {
		fmt.Println("Error opening file:", OutputFileErr)
	}

	ITR.Evaluate(program, env)

	// if evaluatedProgram.Type() == ITR.NumberType {
	// 	fmt.Println(evaluatedProgram.(ITR.NumberVal).Value)
	// 	return
	// }
	// fmt.Println("Evaluated Program : ", evaluatedProgram)
	// fmt.Println("Wrapped Program : ", ITR.RuntimeVal_Wrapper(evaluatedProgram))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
