package program

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	ITR "github.com/iwhitebird/Gor/core/interpreter"
	PSR "github.com/iwhitebird/Gor/core/parser"
)

func CompleteInput(input string) (string, string, error) {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}

	program := parser.ProduceAst(input)

	// Redirect stdout to a buffer
	// Redirect stdout to a buffer
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ITR.Evaluate(program, env)

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	bodyJSON, err := json.MarshalIndent(program.Body, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", "", err
	}

	return string(out), string(bodyJSON), nil
}
