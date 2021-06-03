package main

import (
	"fmt"
	"io"

	"github.com/knadh/stuffbin"
)

type field struct {
	Append    bool
	Func      string
	Offset    string
	Length    int64
	FieldName string
}

type genStruct struct {
	StructName string
	Fields     []field
}

func generateCode(dest io.Writer, fs stuffbin.FileSystem, pkg string, structs map[string][]structField, endianess string) error {
	tmplContext := make(map[string]interface{})
	tmplContext["Pkg"] = pkg
	tmplContext["BuildDate"] = buildDate
	tmplContext["BuildVersion"] = buildVersion

	if endianess == "big" {
		tmplContext["Endian"] = "BigEndian"
	} else {
		tmplContext["Endian"] = "LittleEndian"
	}

	sts := makeStructs(structs)
	fmt.Println(sts)
	tmplContext["Structs"] = sts

	return saveResource("struct", []string{"/templates/gen.tmpl"}, dest, tmplContext, fs)
}

func makeStructs(sts map[string][]structField) []genStruct {
	gensts := make([]genStruct, 0, len(sts))
	for name, fields := range sts {
		st := genStruct{
			StructName: name,
		}

		fds := make([]field, 0, len(fields))
		for _, f := range fields {
			fd := field{
				FieldName: f.Name,
			}

			switch f.Type {
			case "[]byte", "string":
				fd.Append = true
				fd.Offset = fmt.Sprintf("len(s.%s)", f.Name)
			case "int32", "uint32":
				fd.Func = "PutUint32"
				fd.Offset = "4"
				fd.FieldName = fmt.Sprintf("uint32(s.%s)", f.Name)
			case "int64", "uint64":
				fd.Func = "PutUint64"
				fd.Offset = "8"
				fd.FieldName = fmt.Sprintf("uint64(s.%s)", f.Name)
			}

			fds = append(fds, fd)
		}
		st.Fields = fds

		gensts = append(gensts, st)
	}
	return gensts
}
