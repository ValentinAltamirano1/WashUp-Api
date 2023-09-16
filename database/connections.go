package database

import (
	"fmt"
	"log"

	"github.com/ValentinAltamirano1/WashUp-Api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Dbinstance struct {
    Db *gorm.DB
}

var DB Dbinstance

func Connect() Dbinstance {
    const (
        host     = "localhost"
        port     = 5432
        user     = "postgres"
        password = "valenalt"
        dbname   = "washup"
    )

    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open("postgres", connStr)

    if err != nil {
        log.Panic("error connecting to the database: ", err)
    }

    defer db.Close()

    db.AutoMigrate(&model.User{})

    DB := Dbinstance{Db: db}
    return DB
}