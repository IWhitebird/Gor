package program

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	ITR "github.com/iwhitebird/Gor/core/interpreter"
	PSR "github.com/iwhitebird/Gor/core/parser"
)

func Repl(parser PSR.Parser, env *ITR.Environment) {
	fmt.Println("Gor REPL~")

	var scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("~> ")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			break
		}

		program := parser.ProduceAst(input)
		evaluatedProgram := ITR.Evaluate(program, env)

		fmt.Println(evaluatedProgram)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

}
