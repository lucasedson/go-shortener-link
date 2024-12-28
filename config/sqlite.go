package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	schema "github.com/lucasedson/go-shortener-link/schemas"
)

func InitializeSqlite() error {

	dbPath := "./database.db"
	var err error

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		print(err)
		return err
	}

	err = db.AutoMigrate(&schema.Link{})

	if err != nil {
		print(err)
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
