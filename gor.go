package main

import (
	"Gor/ast"
	"Gor/lexer"
	"Gor/parser"
	"fmt"
)

func main() {
	fmt.Println(">> Welcome To Gor Language >:D")

	lexer.Main()
	ast.Main()
	parser.Main()
}
