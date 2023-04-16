package data

import (
	. "pac-sys/entities"
	"pac-sys/utils"
	"strconv"
)

func SaveUser(user UserEntity) {
	createOrUpdate(user, func(u UserEntity) UserEntity {
		return UserEntity{Account: u.Account}
	}, func(tf UserEntity, tt UserEntity) {
		tt.SsoId = tf.SsoId
	})
}

func GetUserInfoForLogin(account string) UserTokenDto {
	db := getDbConn()

	var userGroups []UserGroupEntity
	err := db.Table("users").
		Select("usergroups.user_id, userGroups.group_id").
		Joins("inner join usergroups on users.id = usergroups.user_id").
		Where("users.account=?", account).Find(&userGroups).Error

	if err != nil || len(userGroups) == 0 {
		utils.CreatePanic(500, err.Error())
	}

	token := UserTokenDto{UserId: strconv.Itoa(userGroups[0].Id)}
	token.Groups = make([]string, len(userGroups))
	for i, v := range userGroups {
		token.Groups[i] = strconv.Itoa(v.GroupId)
	}
	return token
}

func SaveGroup(group GroupEntity) {
	createOrUpdate(group, func(g GroupEntity) GroupEntity {
		return GroupEntity{Name: group.Name}
	}, func(tFrom GroupEntity, tTo GroupEntity) {})
}

func SaveUserGroup(entity UserGroupEntity) {
	createOrUpdate(entity,
		func(ug UserGroupEntity) UserGroupEntity {
			return UserGroupEntity{UserId: ug.UserId, GroupId: ug.GroupId}
		}, func(tFrom UserGroupEntity, tTo UserGroupEntity) {})
}
