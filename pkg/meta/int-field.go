package meta

import (
	"fmt"
)

type IntField struct {
	*Field
}

func (f *IntField) AddColumn() (string, error) {
	return fmt.Sprintf("ADD COLUMN %s INT", f.Name), nil
}
