package main

import (
	ITR "Gor/interpreter"
	PSR "Gor/parser"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var env = ITR.EnviromentSetup()
var parser = PSR.Parser{}
var scanner = bufio.NewScanner(os.Stdin)

func CompleteFile() {
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

	evaluatedProgram := ITR.Eval_program(program, *env)

	// if evaluatedProgram.Type() == ITR.NumberType {
	// 	fmt.Println(evaluatedProgram.(ITR.NumberVal).Value)
	// 	return
	// }
	fmt.Println("Evaluated Program : ", evaluatedProgram)
	fmt.Println("Wrapped Program : ", ITR.RuntimeVal_Wrapper(evaluatedProgram))

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func Repl() {
	fmt.Println("Gor REPL~")
	outputFile, OutputFileErr := os.Create("Ast.json")
	if OutputFileErr != nil {
		fmt.Println("Error creating file:", OutputFileErr)
		return
	}
	defer outputFile.Close()

	for {
		fmt.Print("~> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		program := parser.ProduceAst(input)

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

		evaluatedProgram := ITR.Eval_program(program, *env)

		fmt.Println("Evaluated Program : ", evaluatedProgram)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

}

func main() {
	CompleteFile()
	// Repl()
}
