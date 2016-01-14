package main

import (
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "simple.go", nil, 0)
	if err != nil {
		// Whoops!
	}
	printer.Fprint(os.Stdout, fset, file)
}
