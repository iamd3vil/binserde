package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"

	"github.com/spf13/pflag"
)

var (
	buildString string
)

func main() {
	fg := pflag.NewFlagSet("binserde", pflag.ExitOnError)
	fg.Usage = func() {
		fmt.Println(fg.FlagUsages())
		os.Exit(1)
	}
	pkg := fg.String("pkg", "main", "Package to be given for the generated code")
	dir := fg.String("dir", ".", `Path of the directory for finding source files with structs.`)
	dest := fg.String("file", "", "Destination File")
	endianess := fg.String("endianess", "big", "Endianess")
	version := fg.Bool("version", false, "Version")
	fg.Parse(os.Args[1:])

	if *version {
		fmt.Println(buildString)
		os.Exit(0)
	}

	fs, err := initFileSystem()
	if err != nil {
		log.Fatalf("error while getting template for generating code: %v", err)
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, *dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Couldn't parse directory: %v", err)
	}
	sts, err := parseNode(pkgs[*pkg])
	if err != nil {
		log.Fatalf("error while parsing: %v", err)
	}

	code := bytes.NewBuffer([]byte{})
	if err := generateCode(code, fs, *pkg, sts, *endianess); err != nil {
		log.Fatalf("error while generating code: %v", err)
	}

	fPath := *dest
	if fPath == "" {
		fPath = fmt.Sprintf("%s_binserde_gen.go", *pkg)
	}

	f, err := os.Create(fPath)
	if err != nil {
		log.Printf("error while creating file: %v\n", err)
	}
	defer f.Close()

	fmted, err := format.Source(code.Bytes())
	if err != nil {
		log.Fatalf("error while formatting code: %v", err)
	}

	if _, err := f.Write(fmted); err != nil {
		log.Fatalf("error while storing the file: %v", err)
	}
}
