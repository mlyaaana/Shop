package models

type Comment struct {
	Id        string `gorm:"id"`
	UserId    string `gorm:"userid"`
	ProductId string `gorm:"productid"`
	Mention   string `gorm:"mention"`
}

func (Comment) TableName() string {
	return "comments"
}
