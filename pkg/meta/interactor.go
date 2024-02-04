package meta

import (
	"errors"
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/db"
	"gorm.io/gorm"
)

func CreateObejct(obj *Object) error {
	if db.DB == nil {
		return errors.New("connect to a db")
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		create_q, err := obj.BuildQuery()
		if err != nil {
			return err
		}
		if err := tx.Exec(create_q.SQL).Error; err != nil {
			return err
		}

		//create obejct meta
		if err := tx.Create(&obj).Error; err != nil {
			return err
		}
		// build the fields
		for _, f := range obj.Fields {
			if err = AddField(tx, &f, obj); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func AddField(db *gorm.DB, field *Field, object *Object) error {
	field.Object = object.Name
	fmt.Println(field)
	v, err := ParseField(field)
	if err != nil {
		return err
	}
	fielder, ok := v.(Fielder)
	if !ok {
		return fmt.Errorf("fild %v is does not impl Fileder", field)
	}
	if err != nil {
		return err
	}

	if err := fielder.ValidateField(); err != nil {
		return err
	}

	// constraints
	constraints := ""
	if field.Required {
		constraints += "NOT NULL"
	}
	if field.Default != "" {
		constraints += fmt.Sprintf("DEFAULT %s", field.Default)
	}

	adder, err := fielder.AddColumn()
	if err != nil {
		return err
	}
	if err := db.Exec(fmt.Sprintf("ALTER TABLE %s %s %s;", object.Name, adder, constraints)).Error; err != nil {
		return err
	}

	if field.Unique {
		if err := db.Exec(fmt.Sprintf("ALTER TABLE %s ADD UNIQUE (%s);", object.Name, field.Name)).Error; err != nil {
			return err
		}
	}

	if field.Type == Link {
		if field.ReferenceField == "" {
			field.ReferenceField = "id"
		}
		if err := db.Exec(fmt.Sprintf("ALTER TABLE %s ADD FOREIGN KEY (%s) REFERENCES %s(%s);",
			object.Name, field.Name, field.ReferenceObject, field.ReferenceField)).Error; err != nil {
			return err
		}
	}

	// create the field meta
	if err := db.Create(&field).Error; err != nil {
		return err
	}
	return nil
}

func ObjectExists(name string) error {
	var count int64
	err := db.DB.Model(&Object{}).Where(Object{Name: name}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("object not found")
	}
	return nil
}

func GetField(objectName, fieldName string) (*Field, error) {
	var field Field
	err := db.DB.Model(&Field{}).Where(Field{Object: objectName, Name: fieldName}).First(&field).Error
	return &field, err
}
