package models

type User struct {
	Id   string `gorm:"id;primaryKey"`
	Name string `gorm:"name;not null;unique"`
}

func (User) TableName() string {
	return "users"
}
