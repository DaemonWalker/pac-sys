package entities

type UserEntity struct {
	Id       int
	SsoId    int
	Account  string
	Password string
}

type UserTokenDto struct {
	UserId string
	Groups []string
}
