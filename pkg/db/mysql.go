package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlConenctor struct {
	url string
}

func NewMySqlConnector(url string) *MySqlConenctor {
	return &MySqlConenctor{
		url: url,
	}
}

func (c *MySqlConenctor) Connect() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
