package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/db-finpro?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
