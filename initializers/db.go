package initializers

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	// "strconv"
	// "strings"
	"tadeobennett/celebrateease/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// globally accessible
var DB *gorm.DB

// returns the dsn
func getDSNString() string {
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	// port, err := strconv.Atoi(strings.TrimSpace(os.Getenv("DB_PORT")))
	// port := os.Getenv("DB_PORT")
	// if err != nil {
	// 	log.Panic("Port is of the wrong type")
	// }
	dsn := flag.String("dsn", "postgres://"+user+":"+password+"@"+host+"/"+dbname+"?sslmode=disable", "PostgreSQL DSN (Data Source Name)")
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	return *dsn
}

// connect to the postgres database
func ConnectToDatabase() (*sql.DB, error) {
	var err error

	dsn := getDSNString()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to connect to database")
		return nil, err
	}
	// Extract the *sql.DB object from the gorm.DB object
	psqldb, err := DB.DB()
	if err != nil {
		// Log the error and return it
		log.Fatalf("Failed to get *sql.DB object: %v", err)
		return nil, err
	}
	// Check if the connection is alive by pinging the database
	// if err := psqldb.Ping(); err != nil {
	// 	fmt.Println("Failed to ping the database:", err)
	// }else{
	// 	println("ping is good")
	// }

	// Return the *sql.DB object
	return psqldb, err
}

// creates tables in our database that matches the structs in models
func SyncDb() {
	DB.AutoMigrate(&models.Celebrant{})
	DB.AutoMigrate(&models.Page{})
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.User{})
}
