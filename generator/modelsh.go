package generator

import (
	"github.com/oleggator/tnt-generator/types"
	"os"
	"text/template"
)

var modelsHTemplate = NewModelsHTemplates()

const ModelsHTemplate = `
#pragma once

#include <stdint.h>
#include <tarantool/module.h>

#define ARRAY_LEN 1024

{{ range . }}
    {{ template "StructTemplate" . }}

    {{ template "EncoderDocTemplate" . }}
    {{ template "EncoderSignatureTemplate" . }};

    {{ template "DecoderDocTemplate" . }}
    {{ template "DecoderSignatureTemplate" . }};
{{ end }}
`

func NewModelsHTemplates() *template.Template {
	funcMap := template.FuncMap{}

	tpl := template.Must(template.New("ModelsHTemplate").Funcs(funcMap).Parse(ModelsHTemplate))
	initEncoderTemplates(tpl)
	initDecoderTemplates(tpl)
	initStructTemplate(tpl)

	return tpl
}

func GenerateModelsH(path string, structs []types.CStruct) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}

	defer file.Close()


	return modelsHTemplate.Execute(file, structs)
}
