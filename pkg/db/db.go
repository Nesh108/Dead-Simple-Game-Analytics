package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
	"github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	user := os.Getenv("PG_USER")
	pw := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	dbName := os.Getenv("PG_DB")
	dbURL := "postgres://" + user + ":" + pw + "@" + host + ":" + port + "/" + dbName

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	migrateErr := db.AutoMigrate(&models.Event{})
	if migrateErr != nil {
		fmt.Println(migrateErr)
		errorMessage := fmt.Sprintf("DeadSimpleGameAnalytics: %s (%s).", "Failed to AutoMigrate DB", migrateErr.Error())
		services.SendTelegramMessages(errorMessage)
	}

	return db
}
