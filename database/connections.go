package database

import (
	"fmt"
	"log"

	"github.com/ValentinAltamirano1/WashUp-Api/model"
    "gorm.io/driver/postgres"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    const (
        host     = "localhost"
        port     = 5432
        user     = "postgres"
        password = "valenalt"
        dbname   = "washup"
    )

    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

    if err != nil {
        log.Panic("error connecting to the database: ", err)
    }

    _, err = db.DB()
    if err != nil {
        log.Fatal(err)
    }

    DB = db

    
    db.AutoMigrate(&model.User{})
}