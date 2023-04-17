package entities

type PacEntity struct {
	Id      int    `gorm:"primaryKey"`
	GroupId int    `gorm:"index:idx_pacs_group_name,unique"`
	Name    string `gorm:"index:idx_pacs_group_name,unique"`
	Value   string
}

type PacDto struct {
	GroupName string
	PacId     int
	Name      string
	Value     string
}

func (PacEntity) TableName() string {
	return "pacs"
}
