package main

import (
	ITR "Gor/interpreter"
	PSR "Gor/parser"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	outputFile, OutputFileErr := os.Create("Ast.json")

	fmt.Println(">> Welcome To Gor Language >:D")

	// Parser Instance
	parser := PSR.Parser{}

	for {
		fmt.Print(">> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		program := parser.ProduceAst(input)

		// fmt.Println("Node Value : ", program.KindValue)

		bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		// fmt.Println("Node Body JSON : ", string(bodyJSON))

		// Write JSON data to the file
		if OutputFileErr == nil {
			_, writeErr := outputFile.Write(bodyJSON)
			if writeErr != nil {
				fmt.Println("Error writing to file:", writeErr)
			}
		} else {
			fmt.Println("Error opening file:", OutputFileErr)
		}

		evaluatedProgram := ITR.Eval_program(program)

		fmt.Println("Evaluated Program : ", evaluatedProgram)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
