package generator

import (
	"text/template"
)

const DecoderSignatureTemplate = `
int decode_{{ .Name }}({{ .Name }}_t *{{ .Name }}, const char **args, const char **args_end);
`


// TODO implement field generation
const DecoderTemplate = `
{{ template "DecoderSignatureTemplate" . }} {
	{{ range .Fields }}
	
	{{ end }}
}
`

func initDecoderTemplates(parentTemplate *template.Template) {
	funcMap := template.FuncMap{}

	template.Must(parentTemplate.New("DecoderSignatureTemplate").Funcs(funcMap).Parse(DecoderSignatureTemplate))
	template.Must(parentTemplate.New("DecoderTemplate").Funcs(funcMap).Parse(DecoderTemplate))
}
