package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func Parse(reader io.Reader) ([]CStruct, error) {
	jsonDecoder := json.NewDecoder(reader)

	var schema []AvroDefinition
	if err := jsonDecoder.Decode(&schema); err != nil {
		return nil, err
	}

	cStructs := make([]CStruct, 0, len(schema))
	cStructsMap := map[string]*CStruct{}
	for _, definition := range schema {
		cStruct := CStruct{}
		cStruct.Name = definition.Name

		for _, avroField := range definition.Fields {
			fieldCType, err := getCType(avroField.Type, cStructsMap, false, false)
			if err != nil {
				return nil, err
			}

			cStruct.Fields = append(cStruct.Fields, CField{
				Name: avroField.Name,
				Type: fieldCType,
			})
		}

		cStructs = append(cStructs, cStruct)
		cStructsMap[cStruct.Name] = &cStruct
	}

	return cStructs, nil
}

func getCType(avroType interface{}, parsedCStructs map[string]*CStruct, nullable bool, constant bool) (CType, error) {
	var err error

	switch avroType := avroType.(type) {
	case string:
		if avroType == Null {
			return CType(&NullType{}), nil
		}

		cTypeName, ok := simpleTypesMapping[avroType]
		if ok {
			return CType(&SimpleType{
				Nullable: nullable,
				Const:    constant,
				Name:     cTypeName,
			}), nil
		}

		if cStruct, ok := parsedCStructs[avroType]; ok {
			return CType(&NestedType{
				Nullable: nullable,
				Const:    constant,
				NestedStruct: cStruct,
			}), nil
		}

		return nil, fmt.Errorf(`there is no "%s" type`, cTypeName)

	case map[string]interface{}:
		typeType, ok := avroType["type"].(string)
		if !ok {
			return nil, fmt.Errorf(`invalid type of type`)
		}

		switch typeType {
		case Array:
			array := &ArrayType{
				Nullable: nullable,
				Const:    constant,
			}

			itemType, ok := avroType["items"]
			if !ok {
				log.Fatalln(`invalid type of type`)
			}

			if array.ItemType, err = getCType(itemType, parsedCStructs, false, constant); err != nil {
				return nil, err
			}

			return CType(array), nil

		default:
			return nil, fmt.Errorf(`type "%s" is not implemented`, typeType)
		}

	case []interface{}:
		if len(avroType) != 2  {
			return nil, fmt.Errorf(`unsupported union type: allowed only [null, X]`)
		}

		element1, err := getCType(avroType[0], parsedCStructs, true, constant)
		if err != nil {
			return nil, err
		}

		element2, err := getCType(avroType[1], parsedCStructs, true, constant)
		if err != nil {
			return nil, err
		}

		_, element1IsNull := element1.(*NullType)
		_, element2IsNull := element2.(*NullType)

		if element1IsNull && element2IsNull {
			return nil, fmt.Errorf(`unsupported union type: allowed only [null, X]`)
		}

		if element1IsNull {
			return element2, nil
		}

		if element2IsNull {
			return element1, nil
		}

	default:
		return nil, fmt.Errorf(`invalid field type`)
	}

	return nil, nil
}

