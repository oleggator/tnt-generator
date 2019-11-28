package parser


type CType interface {
	IsConst() bool
	IsNullable() bool
}


type NullType struct{
	CType
}

func (nt *NullType) IsConst() bool {
	return false
}

func (nt *NullType) IsNullable() bool {
	return false
}


type ArrayType struct {
	CType
	Const bool
	ItemType CType
	Nullable bool
}

func (st *ArrayType) IsNullable() bool {
	return st.Const
}

func (st *ArrayType) IsConst() bool {
	return st.Const
}

func (st *ArrayType) GetItemType() CType {
	return st.ItemType
}


type SimpleType struct {
	CType
	Const bool
	Name string
	Nullable bool
}

func (st *SimpleType) GetName() string {
	return st.Name
}

func (st *SimpleType) IsConst() bool {
	return st.Const
}


func (st *SimpleType) IsNullable() bool {
	return st.Nullable
}


type NestedType struct {
	CType
	NestedStruct *CStruct
	Const bool
	Nullable bool
}

func (st *NestedType) GetName() string {
	return st.NestedStruct.Name
}

func (st *NestedType) IsConst() bool {
	return st.Const
}

func (st *NestedType) IsNullable() bool {
	return st.Nullable
}


type CField struct {
	Type CType
	Name string
}

type CStruct struct {
	Name string
	Fields []CField
}

type Payload struct {

}
