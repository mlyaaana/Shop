package models

type Product struct {
	Id          string `gorm:"id;primaryKey"`
	Name        string `gorm:"name;not null"`
	Description string `gorm:"description;not null"`
}

func (Product) TableName() string {
	return "products"
}
