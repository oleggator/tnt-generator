package generator

import (
	"fmt"
	"github.com/oleggator/tnt-generator/types"
	"text/template"
)

const StructTemplate = `
typedef struct {
{{- range .Fields }}
	{{ getFieldDefinition . }}
	{{- if IsWithLen .Type }}
	uint32_t {{ .Name }}_len;
	{{- end }}
{{- end}}
} {{ .Name }}_t;
`

func initStructTemplate(parentTemplate *template.Template) {
	funcMap := template.FuncMap{
		"getFieldDefinition": getFieldDefinition,
		"IsWithLen":          IsWithLen,
	}

	template.Must(parentTemplate.New("StructTemplate").Funcs(funcMap).Parse(StructTemplate))
}

func IsWithLen(object interface{}) bool {
	if _, ok := object.(*types.ArrayType); ok {
		return true
	}

	if _, ok := object.(*types.StringType); ok {
		return true
	}

	return false
}

func getFieldDefinition(field *types.CField) (string, error) {
	fieldName := field.Name
	typeName := field.Type.GetName()

	switch field.Type.(type) {
	case *types.PrimitiveType:
		return fmt.Sprintf("%s %s;", typeName, fieldName), nil
	case *types.StringType:
		return fmt.Sprintf("%s %s;", typeName, fieldName), nil
	case *types.ArrayType:
		return fmt.Sprintf("%s %s[ARRAY_LEN];", typeName, fieldName), nil
	case *types.StructType:
		return fmt.Sprintf("%s_t %s;", typeName, fieldName), nil
	}

	return "", fmt.Errorf("unsupported field type %T", field.Type)
}
