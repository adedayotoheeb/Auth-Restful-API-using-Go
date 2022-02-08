package main

import (
	"last/application"
	"last/database"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupDataBaseConnection()
)

func main() {
	application.StartApplication()
	defer database.CloseDataBaseConnection(db)
}
