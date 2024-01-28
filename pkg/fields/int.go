package fields

import "fmt"

type IntField struct {
	*Field
}

func (f *IntField) AddColumn() string {
	return fmt.Sprintf("ADD COLUMN %s INT", f.Name)
}
