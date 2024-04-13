package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URI")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
