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

	err = db.AutoMigrate(&User{}, &Print{})
	if err != nil {
		panic(err)
	}

	database := &Database{
		db: db,
	}

	return database
}
