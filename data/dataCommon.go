package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pac-sys/models"
	"pac-sys/utils"
)

func Migrate() {
	db := getDbConn()
	db.AutoMigrate(&models.PacModel{})
}

func createOrUpdate[T any](t T, keySelector func(T) T, copy func(tFrom T, tTo T)) {
	var entity T

	db := getDbConn()

	err := db.Model(&entity).Where(keySelector(t)).First(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		utils.CreatePanic(500, err.Error())
	}

	if err == nil {
		copy(t, entity)
		err = db.Save(&entity).Error
	} else {
		entity = t
		result := db.Create(&entity)
		if result.Error != nil {
			err = result.Error
		}
	}
	if err != nil {
		utils.CreatePanic(500, err.Error())
	}
}

func queryWithId[T any](condition T) T {
	db := getDbConn()
	var t T
	err := db.Where(&condition).First(&t).Error
	if err != nil {
		utils.ErrorPanic(err)
	}

	return t
}

func query[T any](condition T) []T {
	var ts []T

	db := getDbConn()
	err := db.Where(&condition).Find(&ts).Error
	if err != nil {
		utils.CreatePanic(500, err.Error())
	}

	return ts
}

func getDbConn() *gorm.DB {
	connStr := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		utils.ErrorPanic(err)
	}
	return db
}