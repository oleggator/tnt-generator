package generator

import (
	"github.com/oleggator/tnt-generator/parser"
	"os"
	"text/template"
)

var modelsCTemplate = NewModelsCTemplates()

const ModelsCTemplate = `
#include <stdint.h>
#include <msgpuck.h>
#include <stdio.h>

#include "models.h"

{{ range . }}
    {{- template "EncoderTemplate" . }}
    {{- template "DecoderTemplate" . }}
{{ end -}}
`

func NewModelsCTemplates() *template.Template {
	funcMap := template.FuncMap{}

	tpl := template.Must(template.New("ModelsCTemplate").Funcs(funcMap).Parse(ModelsCTemplate))
	initEncoderTemplates(tpl)
	initDecoderTemplates(tpl)
	initStructTemplate(tpl)

	return tpl
}

func GenerateModelsC(path string, structs []parser.CStruct) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	defer file.Sync()


	return modelsCTemplate.Execute(file, structs)
}
