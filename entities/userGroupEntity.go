package entities

type UserGroupEntity struct {
	Id      int `gorm:"primaryKey"`
	UserId  int `gorm:"index:idx_user_groups_user_group,unique"`
	GroupId int `gorm:"index:idx_user_groups_user_group,unique"`
}

type UserGroupDto struct {
	Id      int
	GroupId []int
}

func (UserGroupEntity) TableName() string {
	return "user_groups"
}
