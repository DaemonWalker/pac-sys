package models

type PacModel struct {
	Key   string `form:"key" gorm:"primaryKey"`
	Value string `form:"value"`
}
