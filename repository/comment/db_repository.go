package comment

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

func (r *DBRepository) Create(comment *domain.Comment) error {
	return r.db.Create(&models.Comment{
		Id:        comment.Id,
		UserId:    comment.UserId,
		ProductId: comment.ProductId,
		Mention:   comment.Mention,
	}).Error
}

func (r *DBRepository) List(productId string) ([]*domain.Comment, error) {
	var commentsModels []models.Comment
	commentsDomain := make([]*domain.Comment, 0)
	err := r.db.Where(&models.Comment{ProductId: productId}).Find(commentsModels).Error
	if err != nil {
		return nil, err
	}

	for _, c := range commentsModels {
		commentsDomain = append(commentsDomain, &domain.Comment{
			Id:        c.Id,
			UserId:    c.UserId,
			ProductId: c.ProductId,
			Mention:   c.Mention,
		})
	}

	return commentsDomain, nil
}

func (r *DBRepository) Delete(id, userId string) {
	r.db.Where(&models.Comment{UserId: userId, Id: id}).Delete(&models.Comment{})
}

func (r *DBRepository) AdminDelete(id string) {
	r.db.Delete(&models.Comment{Id: id})
}
