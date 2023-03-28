package models

type User struct {
	Id   string `gorm:"id"`
	Name string `gorm:"name"`
}

func (User) TableName() string {
	return "users"
}
