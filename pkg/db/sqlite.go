package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqliteConenctor struct {
	path string
}

func NewSqlConnector(path string) *SqliteConenctor {
	return &SqliteConenctor{
		path: path,
	}
}

func (c *SqliteConenctor) Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(c.path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

}
