package data

import (
	"net/http"
	. "pac-sys/entities"
	"pac-sys/share"
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
		Select("ug.user_id, ug.group_id").
		Joins("inner join user_groups ug on users.id = ug.user_id").
		Where("users.account=?", account).Find(&userGroups).Error

	if err != nil {
		share.CreatePanic(http.StatusInternalServerError, err.Error())
	}
	if len(userGroups) == 0 {
		share.CreatePanic(http.StatusBadRequest, "Cannot find group info for this user")
	}

	token := UserTokenDto{UserId: strconv.Itoa(userGroups[0].UserId)}
	token.Groups = make([]int, len(userGroups))
	for i, v := range userGroups {
		token.Groups[i] = v.GroupId
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
