package meta

import (
	"fmt"
)

type TextField struct {
	*Field
}

func (f *TextField) AddColumn() (string, error) {
	return fmt.Sprintf("ADD COLUMN %s VARCHAR(255)", f.Name), nil
}
