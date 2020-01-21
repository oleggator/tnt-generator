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
	char *end = buf;
	end = mp_encode_array(end, {{ len .Fields }});

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

	const buf = "buf"
	const bufEnd = "buf_end"
	const dataEnd = "end"
	src := fmt.Sprintf("%s->%s", cStruct.Name, fieldName)

	switch fieldType := field.Type.(type) {
	case *types.ArrayType:
		primitiveType, isPrimitive := fieldType.ItemType.(*types.PrimitiveType)
		if !isPrimitive {
			return "", fmt.Errorf("arrays of only primitive types are supported")
		}

		arrLen := fmt.Sprintf("%s_len", src)
		const index = "i"
		arrElement := fmt.Sprintf("%s[%s]", src, index)

		primitiveTypeEncoderProc, err := primitiveType.GetEncodeProcedure(dataEnd, arrElement)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(`
			%s = mp_encode_array(%s, %s);
			for (uint32_t %s = 0; %s < %s; ++%s) {
				%s = %s;
			} `,
			dataEnd, dataEnd, arrLen,
			index, index, arrLen, index,
			dataEnd, primitiveTypeEncoderProc,
		), nil

	case *types.StringType:
		strLen := fmt.Sprintf("%s_len", src)
		procCall, err := fieldType.GetEncodeProcedure(dataEnd, src, strLen)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s = %s;", dataEnd, procCall), nil

	case *types.StructType:
		return fmt.Sprintf(`
			err = encode_%s(&%s->%s, %s, %s);
			if (err != 0) { return err; };
		`, typeName, cStruct.Name, fieldName, buf, bufEnd), nil

	case *types.PrimitiveType:
		procCall, err := fieldType.GetEncodeProcedure(dataEnd, src)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s = %s;", dataEnd, procCall), nil

	default:
		return "", fmt.Errorf("unsupported type")
	}
}
