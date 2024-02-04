package fields

import (
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/meta"
)

type BoolField struct {
	*meta.Field
}

func (f *BoolField) AddColumn() string {
	return fmt.Sprintf("ADD COLUMN %s BOOLEAN", f.Name)
}
