package database

import (
	"fmt"
	"last/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Set up is creating conection for the database
func SetupDataBaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to Load Environmental Variables")
	}
	dbServer := os.Getenv("DB_SERVER")
	dbPassord := os.Getenv("DB_PASSORD")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
		dbServer, dbUsername, dbPassord, dbPort, dbName)
	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Conneted to the Database")

	fmt.Println("Migrating models...")

	db.AutoMigrate(&models.Author{}, &models.User{})
	return db
}

func CloseDataBaseConnection(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}
	dbSql.Close()
}
