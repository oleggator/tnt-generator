package generator

import (
	"fmt"
	"github.com/oleggator/tnt-generator/types"
	"text/template"
)

const DecoderDocTemplate = `
/* {{ .Name }} struct decoder
 *
 * @param	{{ printf "%-16s" .Name }}struct to decode 	
 * @param	buf             data buffer
 * @param	buf_end         data buffer end
 *
 * @return	result code
 */`

const DecoderSignatureTemplate = `
int decode_{{ .Name }}({{ .Name }}_t * {{ .Name }},
	const char ** buf, const char ** buf_end)`


// TODO implement field generation
const DecoderTemplate = `
{{ template "DecoderSignatureTemplate" . }}
{
	int err = 0;

	uint32_t field_count = mp_decode_array(buf);
	if (field_count != {{ len .Fields }}) {
		goto wrong_field_count_error;
	}

	{{- range .Fields }}
	// field {{ .Name }}
	{{ getDecodeProcedure $ . }}
	{{- end }}

	return 0;

wrong_field_count_error:
/*	say_error("wrong '%s' fields count - %d, must be %d", ".Name", field_count, {{ len .Fields }});*/
	return 1;
too_big_array:
	return 2;
}
`

func initDecoderTemplates(parentTemplate *template.Template) {
	funcMap := template.FuncMap{
		"getDecodeProcedure": getDecodeProcedure,
	}

	template.Must(parentTemplate.New("DecoderSignatureTemplate").Funcs(funcMap).Parse(DecoderSignatureTemplate))
	template.Must(parentTemplate.New("DecoderTemplate").Funcs(funcMap).Parse(DecoderTemplate))
	template.Must(parentTemplate.New("DecoderDocTemplate").Funcs(funcMap).Parse(DecoderDocTemplate))
}

func getDecodeProcedure(cStruct *types.CStruct, field *types.CField) (string, error) {
	fieldName := field.Name
	typeName := field.Type.GetName()

	switch fieldType := field.Type.(type) {
	case *types.ArrayType:
		primitiveType, isPrimitive := fieldType.ItemType.(*types.PrimitiveType)
		if !isPrimitive {
			return "", fmt.Errorf("arrays of only primitive types are supported")
		}

		primitiveTypeDecoderProc, err := primitiveType.GetDecodeProcedure("buf")
		if err != nil {
			return "", err
		}

		return fmt.Sprintf(`
			%s->%s_len = mp_decode_array(buf);
			for (uint32_t i = 0; i < %s->%s_len; ++i) {
				%s->%s[i] = %s;
			}
	`,
	cStruct.Name, fieldName,
	cStruct.Name, fieldName,
	cStruct.Name, fieldName,
	primitiveTypeDecoderProc), nil

	case *types.StringType:
		procCall, err := fieldType.GetDecodeProcedure("buf", fmt.Sprintf("&%s->%s_len", cStruct.Name, fieldName))
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s->%s = %s;", cStruct.Name, fieldName, procCall), nil

	case *types.StructType:
		return fmt.Sprintf(`
			err = decode_%s(&%s->%s, buf, buf_end);
			if (err != 0) { return err; };
		`, typeName, cStruct.Name, fieldName), nil

	case *types.PrimitiveType:
		procCall, err := fieldType.GetDecodeProcedure("buf")
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s->%s = %s;", cStruct.Name, fieldName, procCall), nil

	default:
		return "", fmt.Errorf("unsupported type")
	}
}
