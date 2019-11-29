package generator

import (
	"text/template"
)

const EncoderDocTemplate = `
/* {{ .Name }} struct encoder
 *
 * @param	{{ .Name }}		struct to encode 	
 * @param	buf					data buffer
 * @param	buf_end				data buffer end
 *
 * @return	result code
 */`

const EncoderSignatureTemplate = `
int encode_{{ .Name }}({{ .Name }}_t * {{ .Name }},
	char ** buf, char ** buf_end)`


// TODO implement field generation
const EncoderTemplate = `
{{ template "EncoderSignatureTemplate" . }}
{
	{{- range .Fields }}
	// field {{ .Name }}
	{{- end }}
}
`

func initEncoderTemplates(parentTemplate *template.Template) {
	funcMap := template.FuncMap{}

	template.Must(parentTemplate.New("EncoderSignatureTemplate").Funcs(funcMap).Parse(EncoderSignatureTemplate))
	template.Must(parentTemplate.New("EncoderTemplate").Funcs(funcMap).Parse(EncoderTemplate))
	template.Must(parentTemplate.New("EncoderDocTemplate").Funcs(funcMap).Parse(EncoderDocTemplate))
}
