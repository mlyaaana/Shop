package product

import (
	"gorm.io/gorm"
	"shop/database/models"
	"shop/domain"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) Create(product *domain.Product) error {
	return r.db.Create(&models.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
	}).Error
}

func (r *DBRepository) Get(id string) (*domain.Product, error) {
	product := &models.Product{}
	err := r.db.Where(&models.Product{Id: id}).First(product).Error
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
	}, nil
}

func (r *DBRepository) List() ([]*domain.Product, error) {
	var productsModels []models.Product
	productsDomain := make([]*domain.Product, 0)
	err := r.db.Find(productsModels).Error
	if err != nil {
		return nil, err
	}

	for _, t := range productsModels {
		productsDomain = append(productsDomain, &domain.Product{
			Id:          t.Id,
			Name:        t.Name,
			Description: t.Description,
		})
	}

	return productsDomain, nil
}

func (r *DBRepository) Delete(id string) {
	r.db.Delete(&models.Product{Id: id})
}
