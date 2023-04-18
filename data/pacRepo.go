package data

import (
	. "pac-sys/entities"
	. "pac-sys/share"
)

func SavePac(pac PacEntity) {
	createOrUpdate(pac, func(pac PacEntity) PacEntity {
		return PacEntity{GroupId: pac.GroupId, Name: pac.Name}
	}, func(tFrom PacEntity, tTo PacEntity) {
		tTo.Value = tFrom.Value
	})
}

func GetPacById(id int) PacEntity {
	return queryWithId(PacEntity{Id: id})
}

func GetPacByGroupId(groupIds []int) []PacDto {
	db := getDbConn()
	var pacs []PacDto
	err := db.Table("groups").
		Select("groups.name, pacs.name, pacs.id, pacs.value").
		Joins("inner join pacs on pacs.group_id=groups.id").
		Where("groups.id in ?", groupIds).Find(&pacs).Error
	if err != nil {
		CreatePanic(500, err.Error())
	}

	return pacs
}
