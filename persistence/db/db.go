package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dsn string) *Database {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Print{})

	database := &Database{
		db: db,
	}

	return database
}
