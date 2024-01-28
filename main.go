package main

import (
	"fmt"
	"log"

	"github.com/AbelBlossom/go-local/pkg/db"
	"github.com/AbelBlossom/go-local/pkg/fields"
	"github.com/AbelBlossom/go-local/pkg/inter"
	"github.com/AbelBlossom/go-local/pkg/object"
)

func main() {
	if err := db.ConenctDB(db.NewSqlConnector("./test.db")); err != nil {
		fmt.Println(err)
		fmt.Println("cannot connect to db")
	}

	err := inter.CreateObejct(object.Object{
		Name: "todo",
		Fields: []map[string]any{
			{
				"type": fields.Text,
				"name": "content",
			},
			{
				"type": fields.Bool,
				"name": "is_completed",
			},
			{
				"type": fields.INT,
				"name": "likes",
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
