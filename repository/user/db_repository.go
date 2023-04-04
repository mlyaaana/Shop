package user

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

func (r *DBRepository) Create(user *domain.User) error {
	return r.db.Create(&models.User{
		Id:   user.Id,
		Name: user.Name,
	}).Error
}

func (r *DBRepository) Get(id string) (*domain.User, error) {
	user := &models.User{}
	err := r.db.Where(&models.User{Id: id}).First(user).Error
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func (r *DBRepository) List() ([]*domain.User, error) {
	var usersModels []*models.User
	usersDomain := make([]*domain.User, 0)
	err := r.db.Find(&usersModels).Error
	if err != nil {
		return nil, err
	}

	for _, u := range usersModels {
		usersDomain = append(usersDomain, &domain.User{
			Id:   u.Id,
			Name: u.Name,
		})
	}
	return usersDomain, nil
}

func (r *DBRepository) Delete(id string) {
	r.db.Delete(&models.User{Id: id})
}
