package card

import (
	"context"
	"github.com/kVinsom/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interfaces.go -destination=../../../mocks/repository/card.go -package=mocks

// ICardRepository defines the behavior required at this layer boundary.
type ICardRepository interface {
	GetAll(ctx context.Context) ([]core.Card, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Card, error)
	GetById(ctx context.Context, id int64) (*core.Card, error)
	GetByNumber(ctx context.Context, number string) (*core.Card, error)
	Blocking(ctx context.Context, id int64) (*core.Card, error)
	Create(ctx context.Context, input *core.Card) (*core.Card, error)
	Delete(ctx context.Context, id int64) error
}
