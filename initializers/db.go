package initializers

import (
	"fmt"
	"os"
	"tadeobennett/celebrateease/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// globally accessible
var DB *gorm.DB

// connect to the postgres database
func ConnectToDatabase() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}


//creates a table in our database that matches the Post struct
func SyncDb() {
	DB.AutoMigrate(&models.Post{})
}
