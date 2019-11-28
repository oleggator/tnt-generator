package generator

import (
	"fmt"
	"github.com/oleggator/tnt-generator/parser"
	"text/template"
)

const StructTemplate = `
typedef struct {
{{- range .Fields }}
	{{ getFieldDefinition . }}
	{{- if isArray .Type }}
	uint32_t {{ .Name }}_len;
	{{- end }}
{{- end}}
} {{ .Name }}_t;
`

func initStructTemplate(parentTemplate *template.Template) {
	funcMap := template.FuncMap{
		"getFieldDefinition": getFieldDefinition,
		"isArray":            isArray,
	}

	template.Must(parentTemplate.New("StructTemplate").Funcs(funcMap).Parse(StructTemplate))
}

func isArray(object interface{}) bool {
	_, ok := object.(*parser.ArrayType)
	return ok
}

func getFieldDefinition(field *parser.CField) (string, error) {
	fieldName := field.Name

	switch fieldType := field.Type.(type) {
	case *parser.SimpleType:
		typeName, err := getTypeName(fieldType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s;", typeName, fieldName), nil

	case *parser.ArrayType:
		typeName, err := getTypeName(fieldType.ItemType)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s %s[ARRAY_LEN];", typeName, fieldName), nil

	case *parser.NestedType:
		typeName, err := getTypeName(fieldType)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s_t %s;", typeName, fieldName), nil
	}

	return "", fmt.Errorf("unsupported field %T", field.Type)
}

func getTypeName(cType parser.CType) (string, error) {
	switch fieldType := cType.(type) {
	case *parser.SimpleType:
		return fieldType.Name, nil

	case *parser.ArrayType:
		return getTypeName(fieldType.ItemType)

	case *parser.NestedType:
		return fieldType.NestedStruct.Name, nil
	}

	return "", fmt.Errorf("unsupported field %T", cType)
}
