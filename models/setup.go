// models/setup.go

package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Singapore",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}),
		&gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Student{})

	if err != nil {
		return
	}

	DB = database
}
