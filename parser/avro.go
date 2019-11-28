package parser

// Types
const (
	Array = "array"
	Null = "null"
)

type AvroField struct {
	Name string `json:"name"`
	Type interface{} `json:"type"`
}

type AvroDefinition struct {
	Name        string      `json:"name"`
	LogicalType string      `json:"logicalType"`
	Fields      []AvroField `json:"fields"`
}

var simpleTypesMapping = map[string]string{
	"int":     "int32_t",
	"integer": "int32_t",

	"uint":     "uint32_t",
	"unsigned": "uint32_t",

	"float":  "float",
	"double": "double",

	"string": "char *",
}
