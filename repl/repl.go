package repl

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
)

func Repl(parser PSR.Parser, env *ITR.Environment) {
	fmt.Println("Gor REPL~")
	var scanner = bufio.NewScanner(os.Stdin)
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

		evaluatedProgram := ITR.Evaluate(program, env)

		fmt.Println("Evaluated Program : ", evaluatedProgram)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

}
