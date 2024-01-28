package fields

import "fmt"

type TextField struct {
	*Field
}

func NewTextField() *TextField {
	return &TextField{
		Field: &Field{},
	}
}

func (f *TextField) AddColumn() string {
	return fmt.Sprintf("ADD COLUMN %s VARCHAR(255)", f.Name)
}
