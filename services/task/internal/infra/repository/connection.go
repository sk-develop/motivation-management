package repository

import (
	"os"

	"github.com/sk-develop/motivation-management/shared/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	conn, err := gorm.Open(postgres.Open(os.Getenv("PG_DSN")), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	return conn
}
