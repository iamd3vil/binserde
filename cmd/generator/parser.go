package main

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

type structField struct {
	Name string
	Type string
	Tag  string
}

func parseNode(node ast.Node) (map[string][]structField, error) {
	structs := make(map[string][]structField)

	var outerErr error
	ast.Inspect(node, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			switch t.Type.(type) {
			case *ast.StructType:
				s := t.Type.(*ast.StructType)
				stFields, err := parseStructSpec(t.Name.String(), s)
				if err != nil {
					outerErr = err
					return false
				}
				if len(stFields) != 0 {
					structs[t.Name.String()] = stFields
				}
			}
		}
		return true

	})

	return structs, outerErr
}

// parseStructSpec parses the struct and returns the fields
func parseStructSpec(structName string, s *ast.StructType) ([]structField, error) {
	stFields := []structField{}
	for _, f := range s.Fields.List {
		// Get `validate` tag. If it doesn't have that fields, ignore that field
		tag := reflect.StructTag(strings.Replace(f.Tag.Value, "`", "", -1)).Get("bin")
		var fieldType string
		switch f.Type.(type) {
		case *ast.Ident:
			fieldType = f.Type.(*ast.Ident).Name
		default:
			return nil, fmt.Errorf("invalid type in struct %s: %v", structName, f.Type)
		}
		stFields = append(stFields, structField{
			Name: f.Names[0].Name,
			Tag:  tag,
			Type: fieldType,
		})
	}
	return stFields, nil
}
