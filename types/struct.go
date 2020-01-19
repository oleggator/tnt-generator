package types

type StructType struct {
	CType
	Struct   *CStruct
	Const    bool
	Nullable bool
}

func (st *StructType) GetName() string {
	return st.Struct.Name
}

func (st *StructType) IsConst() bool {
	return st.Const
}

func (st *StructType) IsNullable() bool {
	return st.Nullable
}
