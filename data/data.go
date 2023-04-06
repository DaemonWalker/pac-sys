package data

import (
	"log"
	"pac-sys/models"
	"pac-sys/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Migrate() {
	db := getDbConn()
	db.AutoMigrate(&models.PacModel{})
}

func Save(pac models.PacModel) {
	db := getDbConn()
	db.Create(&pac)
}

func Get(key string) string {
	db := getDbConn()
	var pac models.PacModel
	db.Where(&models.PacModel{Key: key}).First(&pac)
	return pac.Value
}

func getDbConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		utils.CreatePanic(500, "数据库挂了")
	}
	return db
}
