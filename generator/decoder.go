package generator

import (
	"text/template"
)

const DecoderDocTemplate = `
/* {{ .Name }} struct decoder
 *
 * @param	{{ .Name }}		struct to decode 	
 * @param	buf					data buffer
 * @param	buf_end				data buffer end
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
	{{- range .Fields }}
	// field {{ .Name }}
	{{- end }}
}
`

func initDecoderTemplates(parentTemplate *template.Template) {
	funcMap := template.FuncMap{}

	template.Must(parentTemplate.New("DecoderSignatureTemplate").Funcs(funcMap).Parse(DecoderSignatureTemplate))
	template.Must(parentTemplate.New("DecoderTemplate").Funcs(funcMap).Parse(DecoderTemplate))
	template.Must(parentTemplate.New("DecoderDocTemplate").Funcs(funcMap).Parse(DecoderDocTemplate))
}
