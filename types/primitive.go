package types

import "fmt"

type PrimitiveType struct {
	CType
	Const bool
	Name string
	Nullable bool
}

func (st *PrimitiveType) GetName() string {
	return st.Name
}

func (st *PrimitiveType) IsConst() bool {
	return st.Const
}

func (st *PrimitiveType) IsNullable() bool {
	return st.Nullable
}

func (st *PrimitiveType) GetDecodeProcedure(buf string) (string, error) {
	switch st.Name {
	case "int", "integer", "int64_t":
		return fmt.Sprintf("mp_decode_int(%s)", buf), nil
	case "uint", "unsigned", "uint64_t":
		return fmt.Sprintf("mp_decode_uint(%s)", buf), nil
	case "double":
		return fmt.Sprintf("mp_decode_double(%s)", buf), nil
	default:
		return "", fmt.Errorf("unsupported primitive type: %s", st.Name)
	}
}

func (st *PrimitiveType) GetEncodeProcedure(dst, src string) (string, error) {
	switch st.Name {
	case "int", "integer", "int64_t":
		return fmt.Sprintf("mp_encode_int(%s, %s)", dst, src), nil
	case "uint", "unsigned", "uint64_t":
		return fmt.Sprintf("mp_decode_uint(%s, %s)", dst, src), nil
	case "double":
		return fmt.Sprintf("mp_decode_double(%s, %s)", dst, src), nil
	default:
		return "", fmt.Errorf("unsupported primitive type: %s", st.Name)
	}
}
