package database

import (
	"fmt"
	"log"

	"github.com/nithinsethu/bug-tracking/config"
	"github.com/nithinsethu/bug-tracking/constants"
	"github.com/nithinsethu/bug-tracking/database/models"
	"github.com/nithinsethu/bug-tracking/interfaces"
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

func (pg *PostgresDB) InitDB() {
	q := fmt.Sprintf(`DO $$ BEGIN
	CREATE TYPE role AS ENUM (
					'%v',
					'%v');
	EXCEPTION
	WHEN duplicate_object THEN null;
	END $$;`, constants.RoleAdmin, constants.RoleMember)
	pg.db.Exec(q)
	pg.AutoMigrate()
}

func (pg *PostgresDB) GetdbInstance() *gorm.DB {
	return pg.db
}

func (pg *PostgresDB) AutoMigrate() {
	tables := []interfaces.Model{&models.Organisation{}, &models.User{}}

	for _, t := range tables {
		if !pg.db.Migrator().HasTable(t.TableName()) {
			err := pg.db.AutoMigrate(t)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
