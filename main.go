package main

import (
	"fmt"
	"log"

	"github.com/AbelBlossom/go-local/pkg/db"
	"github.com/AbelBlossom/go-local/pkg/meta"
)

func main() {
	if err := db.ConenctDB(db.NewSqlConnector("./test.db")); err != nil {
		fmt.Println("cannot connect to db")
		log.Fatal(err)
	}

	if err := meta.Migrate(); err != nil {
		log.Fatal(err)
	}
	obj := meta.Object{
		Name: "todo",
		Fields: []meta.Field{
			{
				Name:     "name",
				Label:    "Name",
				Type:     meta.Text,
				Required: true,
			},
			{
				Name:    "completed",
				Label:   "Completed",
				Type:    meta.Bool,
				Default: "true",
			},
			{
				Name:     "todo_id",
				Label:    "Todo ID",
				Type:     meta.Text,
				Unique:   true,
				Required: true,
			},
		},
	}

	err := meta.CreateObejct(&obj)
	if err != nil {
		log.Fatal(err)
	}
	obj2 := meta.Object{
		Name: "test",
		Fields: []meta.Field{
			{
				Name:            "todo",
				Type:            meta.Link,
				ReferenceObject: "todo",
				ReferenceField:  "school_id",
			},
		},
	}

	err = meta.CreateObejct(&obj2)
	if err != nil {
		log.Fatal(err)
	}
}
