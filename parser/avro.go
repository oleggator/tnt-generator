package parser

// Types
const (
	Array  = "array"
	String = "string"
	Null   = "null"
)

type AvroField struct {
	Name string      `json:"name"`
	Type interface{} `json:"type"`
}

type AvroDefinition struct {
	Name        string      `json:"name"`
	LogicalType string      `json:"logicalType"`
	Fields      []AvroField `json:"fields"`
}

var primitiveTypesMapping = map[string]string{
	"int":     "int64_t",
	"integer": "int64_t",

	"uint":     "uint64_t",
	"unsigned": "uint64_t",

	"double": "double",
}
