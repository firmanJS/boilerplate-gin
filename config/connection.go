package config

import (
	"log"
	"os"

	model "github.com/firmanJS/boilerplate-gin/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	if config.GO_ENV != "production" {
		databaseURI <- config.DB_URI
	} else {
		databaseURI <- config.DB_URI
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
