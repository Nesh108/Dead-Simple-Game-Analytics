package db

import (
    "log"

    "github.com/Nesh108/Dead-Simple-Game-Analytics/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    dbURL := "postgres://pg:pass@localhost:5432/game_logger" // TODO - make env variable

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Event{})

    return db
}
