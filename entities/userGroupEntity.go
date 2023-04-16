package entities

type UserGroupEntity struct {
	Id      int
	UserId  int
	GroupId int
}

type UserGroupDto struct {
	Id      int
	GroupId []int
}
