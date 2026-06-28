package card

import (
	"Bank-repository-service/proto/repository/card"
	"Bank-repository-service/proto/repository/common"
	"context"
)

type ICardService interface {
	GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error)
	GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error)
	Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error)
	Create(ctx context.Context, req *card.Card) (*card.Card, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}
