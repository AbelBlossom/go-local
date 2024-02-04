package meta

import (
	"encoding/json"
	"fmt"
)

func ParseField(field *Field) (any, error) {

	switch field.Type {
	case Text:
		return parseData[TextField](field)
	case INT:
		return parseData[IntField](field)
	case Bool:
		return parseData[BoolField](field)
	case Link:
		return parseData[LinkField](field)

	default:
		return nil, fmt.Errorf("invalid field type `%s`", field.Type)
	}
}

func parseData[T any](dt *Field) (*T, error) {
	var val T
	b, err := json.Marshal(dt)
	if err != nil {
		return &val, err
	}
	err = json.Unmarshal(b, &val)
	return &val, err
}

func getType(t FieldType) string {
	switch t {
	case INT:
		return "INT"
	case Bool:
		return "BOOL"
	case Text:
		return "VARCHAR(255)"
	default:
		return "VARCHAR(255)"
	}
}
