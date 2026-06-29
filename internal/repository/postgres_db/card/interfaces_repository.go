package card

import (
	"Bank-repository-service/pkg/core"
	"context"
)

//go:generate mockgen -source=interfaces_repository.go -destination=../../../mocks/reposit/card.go -package=mocks

type ICardRepository interface {
	GetAll(ctx context.Context) ([]core.Card, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Card, error)
	GetById(ctx context.Context, id int64) (*core.Card, error)
	GetByNumber(ctx context.Context, number string) (*core.Card, error)
	Blocking(ctx context.Context, id int64) (*core.Card, error)
	Create(ctx context.Context, input *core.Card) (*core.Card, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
