package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "simple.go", nil, 0)
	if err != nil {
		// Whoops!
	}
	ast.Walk(new(FuncVisitor), file)
	printer.Fprint(os.Stdout, fset, file)
	fmt.Println("-------------------------------------")
	ast.Walk(new(ImportVisitor), file)
	printer.Fprint(os.Stdout, fset, file)
}

type FuncVisitor struct {
}

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	case *ast.FuncDecl:
		t.Name = ast.NewIdent(strings.Title(t.Name.Name))
	}

	return v
}

type ImportVisitor struct{}

func (i *ImportVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	case *ast.GenDecl:
		if t.Tok == token.IMPORT {
			fmt.Println(len(t.Specs))
			newSpecs := []ast.Spec{}
			for _, spec := range t.Specs {
				newSpecs = append(newSpecs, spec)
				fmt.Printf(">>>:% #v\n", spec.(*ast.ImportSpec).Path)
			}
			newPackage := &ast.BasicLit{token.NoPos, token.STRING, `"strings"`}
			newSpecs = append(newSpecs, &ast.ImportSpec{Path: newPackage})
			t.Specs = newSpecs
		}
		return nil
	}

	return i
}
