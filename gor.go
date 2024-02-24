package main

import (
	ITR "Gor/interpreter"
	PSR "Gor/parser"
	PGM "Gor/program"
)

func main() {
	var env = ITR.EnviromentSetup()
	var parser = PSR.Parser{}
	PGM.CompleteFile(parser, env)
	// Repl(parser, env)
	// Testo()
}
