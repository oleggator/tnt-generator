package parser

import (
	"encoding/json"
	"fmt"
	"github.com/oleggator/tnt-generator/types"
	"io"
	"log"
)

func Parse(reader io.Reader) ([]types.CStruct, error) {
	jsonDecoder := json.NewDecoder(reader)

	var schema []AvroDefinition
	if err := jsonDecoder.Decode(&schema); err != nil {
		return nil, err
	}

	cStructs := make([]types.CStruct, 0, len(schema))
	cStructsMap := map[string]*types.CStruct{}
	for _, definition := range schema {
		cStruct := types.CStruct{}
		cStruct.Name = definition.Name

		for _, avroField := range definition.Fields {
			fieldCType, err := getCType(avroField.Type, cStructsMap, false, false)
			if err != nil {
				return nil, err
			}

			cStruct.Fields = append(cStruct.Fields, types.CField{
				Name: avroField.Name,
				Type: fieldCType,
			})
		}

		cStructs = append(cStructs, cStruct)
		cStructsMap[cStruct.Name] = &cStruct
	}

	return cStructs, nil
}

func getCType(avroType interface{}, parsedCStructs map[string]*types.CStruct, nullable bool, constant bool) (types.CType, error) {
	var err error

	switch avroType := avroType.(type) {
	case string:
		if avroType == Null {
			return nil, nil
		}

		if avroType == String {
			return types.CType(&types.StringType{
				Const:    true,
				Nullable: nullable,
			}), nil
		}

		if cTypeName, ok := primitiveTypesMapping[avroType]; ok {
			return types.CType(&types.PrimitiveType{
				Nullable: nullable,
				Const:    constant,
				Name:     cTypeName,
			}), nil
		}

		if cStruct, ok := parsedCStructs[avroType]; ok {
			return types.CType(&types.StructType{
				Nullable: nullable,
				Const:    constant,
				Struct:   cStruct,
			}), nil
		}

		return nil, fmt.Errorf(`there is no "%s" type`, avroType)

	case map[string]interface{}: // complex type
		typeType, ok := avroType["type"].(string)
		if !ok {
			return nil, fmt.Errorf(`invalid type of type`)
		}

		switch typeType {
		case Array:
			array := &types.ArrayType{
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

			return types.CType(array), nil

		default:
			return nil, fmt.Errorf(`type "%s" is not implemented`, typeType)
		}

	case []interface{}: // Union
		if len(avroType) != 2 {
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

		if element1 == nil && element2 == nil {
			return nil, fmt.Errorf(`unsupported union type: allowed only [null, X]`)
		}

		if element1 == nil {
			return element2, nil
		}

		if element2 == nil {
			return element1, nil
		}

	default:
		return nil, fmt.Errorf(`invalid field type`)
	}

	return nil, nil
}
