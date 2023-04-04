package credentials

import (
	"errors"
	"gorm.io/gorm"
	"shop/database/models"
	"shop/domain"
	"shop/utils"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) Create(credentials *domain.Credentials) error {
	return r.db.Create(&models.Credentials{
		UserId:   credentials.UserId,
		Username: credentials.Username,
		Password: utils.MustHashPassword(credentials.Password),
	}).Error
}

func (r *DBRepository) Get(username, password string) (string, error) {
	var creds models.Credentials
	r.db.Where(&models.Credentials{Username: username}).First(&creds)
	if utils.CheckPasswordHash(password, creds.Password) {
		return creds.UserId, nil
	}

	return "", errors.New("incorrect username or password")
}
