package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/config/env"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/postgres"
)

func ConnectDB() {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(env.DbType, "dbname=notification_dev sslmode=disable user=user password=123456 host=localhost")

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if env.AutoMigrateDb == "true" {
		db.AutoMigrate(&model.PR{}, &model.Account{})
	}

}
