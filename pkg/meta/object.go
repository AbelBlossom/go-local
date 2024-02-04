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
