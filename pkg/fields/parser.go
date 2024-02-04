package fields

import (
	"encoding/json"
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/meta"
)

func ParseField(field meta.Field) (any, error) {

	switch field.Type {
	case meta.Text:
		return parseData[TextField](field)
	case meta.INT:
		return parseData[IntField](field)
	case meta.Bool:
		return parseData[BoolField](field)

	default:
		return nil, fmt.Errorf("invalid field type `%s`", field.Type)
	}
}

func parseData[T any](dt meta.Field) (*T, error) {
	var val T
	b, err := json.Marshal(dt)
	if err != nil {
		return &val, err
	}
	err = json.Unmarshal(b, &val)
	return &val, err
}
