package main

import (
	"fmt"
	"log"

	"github.com/AbelBlossom/go-local/pkg/db"
	"github.com/AbelBlossom/go-local/pkg/inter"
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

	err := inter.CreateObejct(meta.Object{
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
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
