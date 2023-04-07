package data

import (
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
	var count int64
	err := db.Model(&models.PacModel{}).Where(&models.PacModel{Key: pac.Key}).Count(&count).Error

	if err != nil {
		utils.ErrorPanic(err)
	}
	if count == 0 {
		err = db.Create(&pac).Error
	} else {
		err = db.Save(&pac).Error
	}

	if err != nil {
		utils.ErrorPanic(err)
	}
}

func Get(key string) string {
	db := getDbConn()
	var pac models.PacModel
	err := db.Where(&models.PacModel{Key: key}).First(&pac).Error
	if err != nil {
		utils.ErrorPanic(err)
	}

	return pac.Value
}

func getDbConn() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		utils.ErrorPanic(err)
	}
	return db
}
