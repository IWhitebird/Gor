package Gor

import (
	ITR "github.com/iwhitebird/Gor/interpreter"
	PSR "github.com/iwhitebird/Gor/parser"
	PGM "github.com/iwhitebird/Gor/program"
)

func New() {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	PGM.CompleteFile(parser, env)
	// Repl(parser, env)
	// Testo()
}

func main() {
	New()
}
