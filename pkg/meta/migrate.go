package meta

import (
	"errors"

	"github.com/AbelBlossom/go-local/pkg/db"
)

func Migrate() error {
	if db.DB == nil {
		return errors.New("connect to db")
	}
	return db.DB.AutoMigrate(&Object{}, &Field{})
}
