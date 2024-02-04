package fields

import (
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/meta"
)

type IntField struct {
	*meta.Field
}

func (f *IntField) AddColumn() string {
	return fmt.Sprintf("ADD COLUMN %s INT", f.Name)
}
