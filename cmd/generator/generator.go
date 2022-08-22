package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/knadh/stuffbin"
)

// genMarshalField is each field in genMarshalStruct.
type genMarshalField struct {
	Append            bool
	Func              string
	Offset            string
	Length            int64
	FieldName         string
	HasToBeMarshalled bool
}

// genMarshalStruct is the struct passed down to template for generating marshalling code.
type genMarshalStruct struct {
	StructName  string
	Fields      []genMarshalField
	NeedsBuffer bool
}

// HasFieldToBeMarshalled returns true when any of the fields has HasToBeMarshalled as true.
func (g *genMarshalStruct) HasFieldToBeMarshalled() bool {
	for _, f := range g.Fields {
		if f.HasToBeMarshalled || f.Append {
			return true
		}
	}
	return false
}

// genUnmarshalField is each field in genMarshalStruct.
type genUnmarshalField struct {
	Append    bool
	String    bool
	Struct    bool
	Float64   bool
	Func      string
	Offset    string
	Length    int64
	FieldName string
	Type      string
}

// genMarshalStruct is the struct passed down to template for generating marshalling code.
type genUnmarshalStruct struct {
	StructName  string
	NeedsBuffer bool
	Fields      []genUnmarshalField
}

func generateCode(dest io.Writer, fs stuffbin.FileSystem, pkg string, structs map[string][]structField, endianess string) error {
	tmplContext := make(map[string]interface{})
	addInitialContext(tmplContext, pkg, endianess)

	// Add structs to tmplContext for marshalling
	sts, err := makeMarshallingStructs(structs)
	if err != nil {
		return err
	}
	tmplContext["MarshalStructs"] = sts

	// Add structs to tmplContext for unmarshalling
	usts, err := makeUnmarshallingStructs(structs)
	if err != nil {
		return err
	}
	tmplContext["UnmarshalStructs"] = usts

	return saveResource("struct", []string{"/templates/gen.tmpl"}, dest, tmplContext, fs)
}

func makeMarshallingStructs(sts map[string][]structField) ([]genMarshalStruct, error) {
	gensts := make([]genMarshalStruct, 0, len(sts))
	for name, fields := range sts {
		st := genMarshalStruct{
			StructName: name,
		}

		fds := make([]genMarshalField, 0, len(fields))
		for _, f := range fields {
			fd := genMarshalField{}

			switch f.Type {
			case "[]byte", "string":
				fd.Append = true
				len, err := getLength(f, "len")
				if err != nil {
					return nil, err
				}
				fd.Offset = len
				fd.FieldName = fmt.Sprintf("s.%s", f.Name)
			case "int32", "uint32":
				fd.Func = "PutUint32"
				fd.Offset = "4"
				fd.FieldName = fmt.Sprintf("uint32(s.%s)", f.Name)
				st.NeedsBuffer = true
			case "int64", "uint64":
				fd.Func = "PutUint64"
				fd.Offset = "8"
				fd.FieldName = fmt.Sprintf("uint64(s.%s)", f.Name)
				st.NeedsBuffer = true
			case "int16", "uint16":
				fd.Func = "PutUint16"
				fd.Offset = "2"
				fd.FieldName = fmt.Sprintf("uint16(s.%s)", f.Name)
				st.NeedsBuffer = true
			case "float64":
				fd.Offset = "8"
				fd.FieldName = fmt.Sprintf("math.Float64bits(s.%s)", f.Name)
				fd.Func = "PutUint64"
				st.NeedsBuffer = true
			default:
				fd.Append = true
				fd.FieldName = fmt.Sprintf("s.%s", f.Name)
				// Since this field is an embedded struct/custom type, it has to be marshalled.
				fd.HasToBeMarshalled = true
			}

			fds = append(fds, fd)
		}
		st.Fields = fds

		gensts = append(gensts, st)
	}
	return gensts, nil
}

func makeUnmarshallingStructs(sts map[string][]structField) ([]genUnmarshalStruct, error) {
	gensts := make([]genUnmarshalStruct, 0, len(sts))
	for name, fields := range sts {
		st := genUnmarshalStruct{
			StructName: name,
		}

		fds := make([]genUnmarshalField, 0, len(fields))
		for _, f := range fields {
			fd := genUnmarshalField{
				FieldName: f.Name,
				Type:      f.Type,
			}

			switch f.Type {
			case "[]byte":
				fd.Append = true
				len, err := getLength(f, "len")
				if err != nil {
					return nil, err
				}
				fd.Offset = len
			case "string":
				len, err := getLength(f, "len")
				if err != nil {
					return nil, err
				}
				fd.Offset = len
				fd.String = true
			case "int32", "uint32":
				fd.Func = "Uint32"
				fd.Offset = "4"
				st.NeedsBuffer = true
			case "int64", "uint64":
				fd.Func = "Uint64"
				fd.Offset = "8"
				st.NeedsBuffer = true
			case "int16", "uint16":
				fd.Func = "Uint16"
				fd.Offset = "2"
				st.NeedsBuffer = true
			case "float64":
				fd.Float64 = true
				fd.Func = "Uint64"
				fd.Offset = "8"
				st.NeedsBuffer = true
			default:
				// This might be a struct or a custom type.
				fd.Struct = true
			}

			fds = append(fds, fd)
		}

		st.Fields = fds
		gensts = append(gensts, st)
	}
	return gensts, nil
}

// func marshalStruct(typ string, sts map[string][]structField, f structField) (genMarshalField, error) {
// 	fd := genMarshalField{}

// 	// Check if the type is a struct in the package.
// 	_, ok := sts[typ]
// 	if !ok {
// 		return genMarshalField{}, fmt.Errorf("error while generating marshalling code: %s", typ)
// 	}
// 	fd.Append = true
// 	len, err := getLength(f, "len")
// 	if err != nil {
// 		return genMarshalField{}, err
// 	}
// 	fd.Offset = len
// 	fd.FieldName = fmt.Sprintf("s.%s.Marshal()", f.Name)
// }

func addInitialContext(tmplContext map[string]interface{}, pkg, endianess string) {
	tmplContext["Pkg"] = pkg
	tmplContext["BuildString"] = buildString

	if endianess == "big" {
		tmplContext["Endian"] = "BigEndian"
	} else {
		tmplContext["Endian"] = "LittleEndian"
	}
}

func getLength(fd structField, tag string) (string, error) {
	attrs := strings.Split(fd.Tag, " ")
	// Find length in attrs
	for _, attr := range attrs {
		as := strings.Split(attr, "=")
		if len(as) != 2 {
			continue
		}
		if as[0] == tag {
			// Return if length is an integer,
			// if not assume that this is a field name.
			if _, err := strconv.ParseInt(as[1], 10, 64); err != nil {
				return fmt.Sprintf("int(s.%s)", as[1]), nil
			}
			return as[1], nil
		}
	}
	return "", fmt.Errorf("error while finding length attribute for field %s", fd.Name)
}
