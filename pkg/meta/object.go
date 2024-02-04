package meta

import (
	"fmt"

	"github.com/AbelBlossom/go-local/pkg/db"
	"github.com/AbelBlossom/go-local/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Object struct {
	Name   string  `json:"name" gorm:"unique"`
	Fields []Field `gorm:"-"`
}

func LeanObject(name string) *Object {
	return &Object{
		Name: name,
	}
}

func (Object) TableName() string {
	return utils.ObjectMetaTable
}

func (obj *Object) GetFields() ([]Field, error) {
	var fields []Field
	err := db.DB.Table(utils.FieldMetaTable).Where("object = ?", obj.Name).Find(&fields).Error
	return fields, err
}

func (o *Object) BuildQuery() (clause.Expr, error) {
	create_template := `
		CREATE TABLE %s (
			id int NOT NULL PRIMARY KEY
		);
		`
	q := gorm.Expr(fmt.Sprintf(create_template, o.Name), o.Name)
	return q, nil
}

func (o *Object) AddFields(fields ...Field) error {
	if len(fields) == 0 {
		return nil
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		for _, f := range fields {
			if err := AddField(tx, &f, o); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
