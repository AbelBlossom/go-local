package object

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ???
type Objecter interface {
	Create() error
}

type Object struct {
	Name string `json:"name"`

	//TODO: fields should be map[string]any so that it can be more dynamic
	Fields []map[string]any
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
