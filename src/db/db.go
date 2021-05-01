package db

import (
	"fmt"

	"github.com/rvramesh/tg-microwin-bot/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

// Init creates a connection to mysql database and
// migrates any new models
func Init() {
	db, err = gorm.Open(sqlite.Open("enrollment.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
	}

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Initialized DB!")

}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
