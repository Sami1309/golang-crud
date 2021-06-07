package middleware

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go-crud-app/models"
)

var DB *gorm.DB

func ConnectDB() {

	//Establish database connection via sqlite3
	database, err := gorm.Open("sqlite3", "books.db")

	//Check if failed to connect
	if err != nil {
		panic("Failed to connect to database")
	}

	//Establish book schema via gorm
	database.AutoMigrate(&models.Book{})

	//Allow DB access from models package
	DB = database
}
