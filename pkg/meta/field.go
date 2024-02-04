package meta

import (
	"github.com/AbelBlossom/go-local/pkg/utils"
)

type Fielder interface {
	// this is reposible for adding fields to the object
	AddColumn() (string, error)
	Validate(value any) error // this is used to validate  the value before inset
	ValidateField() error
	Format(value any) (any, error) // this is used to format the value bofore inserting to the db
}

type Field struct {
	// the type difinish for the field
	Name        string    `json:"name"`
	Object      string    `json:"object,omitempty"`
	Type        FieldType `json:"type"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Default     string    `json:"default"`
	Required    bool      `json:"required"`
	Unique      bool      `json:"unique"`

	// link field options
	ReferenceObject string `json:"ref_object"`
	ReferenceField  string `json:"ref_field"`
}

func (Field) TableName() string {
	return utils.FieldMetaTable

}
func (f *Field) Validate(value any) error {
	return nil
}

func (f *Field) ValidateField() error {
	return nil
}

func (f *Field) Format(value any) (any, error) {
	return value, nil
}
