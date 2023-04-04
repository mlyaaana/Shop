package domain

import "github.com/google/uuid"

type Comment struct {
	Id        string
	UserId    string
	ProductId string
	Mention   string
}

func NewComment(userId, productId, mention string) *Comment {
	return &Comment{
		Id:        uuid.New().String(),
		UserId:    userId,
		ProductId: productId,
		Mention:   mention,
	}
}
