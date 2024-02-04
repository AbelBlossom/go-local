package meta

import (
	"fmt"
)

type BoolField struct {
	*Field
}

func (f *BoolField) AddColumn() (string, error) {
	return fmt.Sprintf("ADD COLUMN %s BOOLEAN", f.Name), nil
}
