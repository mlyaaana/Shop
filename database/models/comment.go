package models

type Comment struct {
	Id        string `gorm:"id;primaryKey"`
	UserId    string `gorm:"userId;not null"`
	ProductId string `gorm:"productId;not null"`
	Mention   string `gorm:"mention;not null"`
}

func (Comment) TableName() string {
	return "comments"
}
