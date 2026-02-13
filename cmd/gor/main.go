package main

import (
	"flag"
	"fmt"
	"os"

	Gor "github.com/iwhitebird/Gor"
)

var version = "1.0.0"

func usage() {
	fmt.Print(`Gor - An interpreted programming language written in Go

Usage:
  gor <file.gor>         Run a Gor source file
  gor --repl             Start the interactive REPL
  gor --ast <file.gor>   Print the AST of a source file

Flags:
`)
	flag.PrintDefaults()
}

func main() {
	repl := flag.Bool("repl", false, "Start the interactive REPL")
	ast := flag.Bool("ast", false, "Print the AST as JSON instead of executing")
	ver := flag.Bool("version", false, "Print version")

	flag.Usage = usage
	flag.Parse()

	if *ver {
		fmt.Println("gor", version)
		return
	}

	if *repl {
		Gor.Repl()
		return
	}

	args := flag.Args()

	if len(args) == 0 {
		usage()
		os.Exit(0)
	}

	filePath := args[0]

	if *ast {
		Gor.PrintAST(filePath)
		return
	}

	Gor.RunFromFile(filePath)
}
