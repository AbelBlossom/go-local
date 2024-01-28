package fields

import "fmt"

type BoolField struct {
	*Field
}

func (f *BoolField) AddColumn() string {
	return fmt.Sprintf("ADD COLUMN %s BOOLEAN", f.Name)
}
