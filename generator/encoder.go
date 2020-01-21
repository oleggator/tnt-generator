package generator

import (
	"fmt"
	"github.com/oleggator/tnt-generator/types"
	"text/template"
)

const EncoderDocTemplate = `
/* {{ .Name }} struct encoder
 *
 * @param	{{ printf "%-16s" .Name }}struct to encode 	
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */`

const EncoderSignatureTemplate = `
int encode_{{ .Name }}({{ .Name }}_t * {{ .Name }},
	char * buf, char * buf_end)`

// TODO implement field generation
const EncoderTemplate = `
{{ template "EncoderSignatureTemplate" . }}
{
	int err = 0;
	char *end;
	end = mp_encode_array(buf, {{ len .Fields }});

	{{- range .Fields }}
	// field {{ .Name }}
	{{ getEncodeProcedure $ . }}
	{{- end }}
}
`

func initEncoderTemplates(parentTemplate *template.Template) {
	funcMap := template.FuncMap{
		"getEncodeProcedure": getEncodeProcedure,
	}

	template.Must(parentTemplate.New("EncoderSignatureTemplate").Funcs(funcMap).Parse(EncoderSignatureTemplate))
	template.Must(parentTemplate.New("EncoderTemplate").Funcs(funcMap).Parse(EncoderTemplate))
	template.Must(parentTemplate.New("EncoderDocTemplate").Funcs(funcMap).Parse(EncoderDocTemplate))
}

func getEncodeProcedure(cStruct *types.CStruct, field *types.CField) (string, error) {
	fieldName := field.Name
	typeName := field.Type.GetName()

	switch fieldType := field.Type.(type) {
	case *types.ArrayType:
		primitiveType, isPrimitive := fieldType.ItemType.(*types.PrimitiveType)
		if !isPrimitive {
			return "", fmt.Errorf("arrays of only primitive types are supported")
		}

		bufEnd := "end"
		arrLen := fmt.Sprintf("%s->%s_len", cStruct.Name, fieldName)
		index := "i"
		arrElement := fmt.Sprintf("%s->%s[%s]", cStruct.Name, fieldName, index)

		primitiveTypeEncoderProc, err := primitiveType.GetEncodeProcedure(bufEnd, arrElement)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(`
			%s = mp_encode_array(buf, %s);
			for (uint32_t i = 0; i < %s; ++i) {
				%s = %s;
			}
		`, bufEnd, arrLen, arrLen, bufEnd, primitiveTypeEncoderProc), nil

	case *types.StringType:
		src := fmt.Sprintf("%s->%s", cStruct.Name, fieldName)
		procCall, err := fieldType.GetEncodeProcedure("end", src, fmt.Sprintf("%s->%s_len", cStruct.Name, fieldName))
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s = %s;", "end", procCall), nil

	case *types.StructType:
		return fmt.Sprintf(`
			err = encode_%s(&%s->%s, buf, buf_end);
			if (err != 0) { return err; };
		`, typeName, cStruct.Name, fieldName), nil

	case *types.PrimitiveType:
		src := fmt.Sprintf("%s->%s", cStruct.Name, fieldName)
		procCall, err := fieldType.GetEncodeProcedure("end", src)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s = %s;", "end", procCall), nil

	default:
		return "", fmt.Errorf("unsupported type")
	}
}
