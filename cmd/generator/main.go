package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

var (
	buildDate, buildVersion string
)

func main() {
	pkg := flag.String("pkg", "main", "Package to be given for the generated code")
	dir := flag.String("dir", ".", `Path of the directory for finding source files with structs.`)
	dest := flag.String("file", "", "Destination File")
	endianess := flag.String("endianess", "big", "Endianess")
	flag.Parse()

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

	wtr := bufio.NewWriter(f)
	defer wtr.Flush()
	if _, err := wtr.Write(fmted); err != nil {
		log.Fatalf("error while storing the file: %v", err)
	}
}
