package database

import (
	"goapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {

	dsn := "root:password@/goapp"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}
	database.AutoMigrate(&models.User{})
}
