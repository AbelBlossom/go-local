package fields

import (
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/meta"
)

type TextField struct {
	*meta.Field
}

func NewTextField() *TextField {
	return &TextField{
		Field: &meta.Field{},
	}
}

func (f *TextField) AddColumn() string {
	return fmt.Sprintf("ADD COLUMN %s VARCHAR(255)", f.Name)
}
