package Gor

import (
	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
	PGM "github.com/iwhitebird/Gor/program"
)

func Repl() {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	PGM.Repl(parser, env)
}

func RunFromInput(input string) (string, error) {
	return PGM.CompleteInput(input)
}

func RunFromFile(file_path string) (string, error) {
	return PGM.CompleteFile(file_path)
}
