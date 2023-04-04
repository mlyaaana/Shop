package models

type Credentials struct {
	UserId   string `gorm:"user_id;not null"`
	Username string `gorm:"username;not null;unique"`
	Password string `gorm:"password;not null"`
}

func (Credentials) TableName() string {
	return "credentials"
}
