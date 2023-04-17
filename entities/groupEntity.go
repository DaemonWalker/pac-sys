package entities

type GroupEntity struct {
	Id   int    `gorm:"primaryKey"`
	Name string `gorm:"index:idx_groups_name,unique"`
}

func (GroupEntity) TableName() string {
	return "groups"
}
