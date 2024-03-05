package program

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
)

func Repl(parser PSR.Parser, env *ITR.Environment) <-chan [2]interface{} {
	fmt.Println("Gor REPL~")
	var resultChan = make(chan [2]interface{})

	go func() {
		defer close(resultChan)

		var scanner = bufio.NewScanner(os.Stdin)

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
				resultChan <- [2]interface{}{nil, err}
				continue
			}

			evaluatedProgram := ITR.Evaluate(program, env)

			resultChan <- [2]interface{}{evaluatedProgram, string(bodyJSON)}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error:", err)
		}
	}()

	return resultChan
}
