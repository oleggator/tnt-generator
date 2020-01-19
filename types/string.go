package types

import "fmt"

type StringType struct {
	CType
	Const bool
	Nullable bool
}

func (st *StringType) IsNullable() bool {
	return st.Const
}

func (st *StringType) IsConst() bool {
	return st.Const
}

func (at *StringType) GetName() string {
	if at.IsConst() {
		return "const char *"
	}
	return "char *"
}

func (st *StringType) GetDecodeProcedure(buf string, dst string) (string, error) {
	return fmt.Sprintf("mp_decode_str(%s, %s)", buf, dst), nil
}
