package entities

type PacEntity struct {
	Id      int
	GroupId int
	Name    string
	Value   string
}

type PacDto struct {
	GroupName string
	PacId     int
	Name      string
	Value     string
}
