package fields

type Fielder interface {
	// this is reposible for adding fields to the object
	AddColumn() string
	Validate(value any) error      // this is used to validate  the value before inset
	Format(value any) (any, error) // this is used to format the value bofore inserting to the db
}

type Field struct {
	// the type difinish for the field
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Default     string `json:"default"`
	Required    string `json:"required"`

	// link field options
	ReferenceObject string `json:"ref_object"`
	ReferenceField  string `json:"ref_field"`
}

func (f *Field) Validate(value any) error {
	return nil
}

func (f *Field) Format(value any) (any, error) {
	return value, nil
}

// This must be implented by every field type
func (*Field) GetDef() string {
	return ""
}
