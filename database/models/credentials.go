package models

type Credentials struct {
	UserId   string `gorm:"user_id"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

func (Credentials) TableName() string {
	return "credentials"
}
