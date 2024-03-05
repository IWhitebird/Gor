package Gor

import (
	"fmt"

	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
	PGM "github.com/iwhitebird/Gor/program"
)

func Repl() <-chan [2]interface{} {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	resultChan := PGM.Repl(parser, env)
	return resultChan
}

func RunFromInput(input string) interface{} {
	data, err := PGM.CompleteInput(input)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return data
}

func RunFromFile(file_path string) interface{} {
	data, err := PGM.CompleteFile(file_path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return data
}
