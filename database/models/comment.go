package models

type Comment struct {
	Id        string `gorm:"id"`
	UserId    string `gorm:"userId"`
	ProductId string `gorm:"productId"`
	Mention   string `gorm:"mention"`
}

func (Comment) TableName() string {
	return "comments"
}
