package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConenctor interface {
	Connect() (*gorm.DB, error)
}

func ConenctDB(conn DBConenctor) error {
	_db, err := conn.Connect()
	if err != nil {
		return err
	}
	DB = _db
	// migrate metadata
	return nil
}
