package generator

import (
	"text/template"
)

const EncoderSignatureTemplate = `
int encode_{{ .Name }}({{ .Name }}_t *{{ .Name }}, char **buf, char **buf_end);
`


// TODO implement field generation
const EncoderTemplate = `
{{ template "EncoderSignatureTemplate" . }} {
	{{ range .Fields }}
	
	{{ end }}
}
`

func initEncoderTemplates(parentTemplate *template.Template) {
	funcMap := template.FuncMap{}

	template.Must(parentTemplate.New("EncoderSignatureTemplate").Funcs(funcMap).Parse(EncoderSignatureTemplate))
	template.Must(parentTemplate.New("EncoderTemplate").Funcs(funcMap).Parse(EncoderTemplate))
}
