package comment

import "shop/domain"

type Repository interface {
	Create(comment *domain.Comment) error
	List(productId string) ([]*domain.Comment, error)
	Delete(id, userId string)
	AdminDelete(id string)
}
