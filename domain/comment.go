package domain

import "github.com/google/uuid"

type Comment struct {
	Id        string
	UserId    string
	ProductId string
	Mention   string
}

func NewComment(userid, productid, mention string) *Comment {
	return &Comment{
		Id:        uuid.New().String(),
		UserId:    userid,
		ProductId: productid,
		Mention:   mention,
	}
}
