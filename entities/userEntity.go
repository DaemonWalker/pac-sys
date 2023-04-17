package entities

type UserEntity struct {
	Id       int    `gorm:"primaryKey"`
	SsoId    string `gorm:"index:idx_users_sso_id,unique"`
	Account  string `gorm:"index:idx_users_account,unique"`
	Password string
}

type UserTokenDto struct {
	UserId string
	Groups []string
}

func (UserEntity) TableName() string {
	return "users"
}
