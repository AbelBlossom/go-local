package inter

import (
	"errors"
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/db"
	"github.com/AbelBlossom/go-local/pkg/fields"
	"github.com/AbelBlossom/go-local/pkg/meta"
	"gorm.io/gorm"
)

func CreateObejct(obj meta.Object) error {
	if db.DB == nil {
		return errors.New("connect to a db")
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		create_q, err := obj.BuildQuery()
		if err != nil {
			return err
		}
		if err := tx.Exec(create_q.SQL, create_q.Vars...).Error; err != nil {
			return err
		}

		//create obejct meta
		if err := tx.Create(&obj).Error; err != nil {
			return err
		}
		// build the fields
		for _, f := range obj.Fields {
			f.Object = obj.Name
			v, err := fields.ParseField(f)
			if err != nil {
				return err
			}
			fieder, ok := v.(meta.Fielder)
			if !ok {
				return errors.New("invalid field")
			}
			if err != nil {
				return err
			}

			// constraints
			constraints := ""
			if f.Required {
				constraints += "NOT NULL"
			}

			if f.Unique {
				constraints += " UNIQUE"
			}

			if f.Default != "" {
				constraints += fmt.Sprintf("DEFAULT '%s'", f.Default)
			}
			if err := tx.Exec(fmt.Sprintf("ALTER TABLE %s %s %s;", obj.Name, fieder.AddColumn(), constraints)).Error; err != nil {
				return err
			}

			// create the field meta
			if err := tx.Create(&f).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
