package fields

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FieldType string

const (
	Text FieldType = "TEXT"
	INT  FieldType = "INT"
	Bool FieldType = "BOOL"
)

func ParseField(data map[string]any) (any, error) {
	fieldType, ok := data["type"]
	if !ok {
		return nil, errors.New("field type must be provided")
	}

	switch fieldType {
	case Text:
		return parseData[TextField](data)
	case INT:
		return parseData[IntField](data)
	case Bool:
		return parseData[BoolField](data)

	default:
		return nil, fmt.Errorf("invalid field type `%s`", fieldType)
	}
}

func parseData[T any](dt map[string]any) (*T, error) {
	var val T
	b, err := json.Marshal(dt)
	if err != nil {
		return &val, err
	}
	err = json.Unmarshal(b, &val)
	return &val, err
}
