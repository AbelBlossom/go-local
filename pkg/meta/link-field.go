package meta

import (
	"errors"
	"fmt"
)

type LinkField struct {
	*Field
}

func (f *LinkField) AddColumn() (string, error) {
	// referene type
	field, err := GetField(f.ReferenceObject, f.ReferenceField)
	if err != nil {
		return "", err
	}

	t := getType(field.Type)
	return fmt.Sprintf("ADD COLUMN %s %s FOREIGN KEY REFERENCES %s(%s)",
		f.Name, t, f.ReferenceObject, f.ReferenceField), nil
}

func (f *LinkField) ValidateField() error {
	// get the field meta for the reference table
	if err := ObjectExists(f.ReferenceObject); err != nil {
		return err
	}
	if f.ReferenceField != "" && f.ReferenceField != "id" {
		field, err := GetField(f.ReferenceObject, f.ReferenceField)
		if err != nil {
			return err
		}
		if !field.Unique {
			return errors.New("only unique fields can be referenced")
		}
	}
	return nil
}
