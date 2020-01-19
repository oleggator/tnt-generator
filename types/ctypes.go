package types

type CType interface {
	IsConst() bool
	IsNullable() bool
	GetName() string
}

type CStruct struct {
	Name   string
	Fields []CField
}

type CField struct {
	Type CType
	Name string
}
