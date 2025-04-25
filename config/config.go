// config/config.go
package config

import (
	"fmt"
	"jwt-gin-practice/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func LoadConfig() {
	// Load configuration from environment or file
	godotenv.Load(".env") // charge les variables

}

func ConnectDatabase() {
	var err error

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Println("We are connected to the database ", DB)
	}

	err = DB.AutoMigrate(&models.User{}).Error

	if err != nil {
		err = fmt.Errorf("gorm AutoMigrate: %w", err)
	}
}
