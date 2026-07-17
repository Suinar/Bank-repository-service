package card

import (
	"context"
	"github.com/Suinar/Bank-proto/repository/card"
	"github.com/Suinar/Bank-proto/repository/common"
)

// ICardHandler defines the behavior required at this layer boundary.
type ICardHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error)
	GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error)
	Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error)
	Create(ctx context.Context, req *card.Card) (*card.Card, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
