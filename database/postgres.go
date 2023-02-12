package database

import (
	"log"

	"github.com/nithinsethu/bug-tracking/config"
	"github.com/nithinsethu/bug-tracking/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB() *PostgresDB {
	db, err := gorm.Open(postgres.Open(config.PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgres connection successful...")
	return &PostgresDB{db: db}
}

func (pg *PostgresDB) AutoMigrate() {
	if !pg.db.Migrator().HasTable("organisations") {
		err := pg.db.AutoMigrate(&models.Organisation{})
		if err != nil {
			log.Fatal(err)
		}
	}
}
