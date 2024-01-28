package inter

import (
	"errors"
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/db"
	"github.com/AbelBlossom/go-local/pkg/fields"
	"github.com/AbelBlossom/go-local/pkg/object"
	"gorm.io/gorm"
)

func CreateObejct(obj object.Object) error {
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
		// build the fields
		for _, f := range obj.Fields {
			v, err := fields.ParseField(f)
			if err != nil {
				return err
			}
			fieder, ok := v.(fields.Fielder)
			if !ok {
				return errors.New("invalid field")
			}
			if err != nil {
				return err
			}
			if err := tx.Exec(fmt.Sprintf("ALTER TABLE %s %s;", obj.Name, fieder.AddColumn())).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
