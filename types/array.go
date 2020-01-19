package types

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

func (at *ArrayType) GetName() string {
	return at.ItemType.GetName()
}
