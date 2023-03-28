package models

type Product struct {
	Id          string `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
}

func (Product) TableName() string {
	return "products"
}
