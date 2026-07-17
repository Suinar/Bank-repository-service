package credit

import (
	"context"
	"github.com/Suinar/Bank-repository-service/pkg/core"
)

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=interfaces_repository.go -destination=../../../mocks/repository/credit.go -package=mocks

// ICreditRepository defines the behavior required at this layer boundary.
type ICreditRepository interface {
	GetAll(ctx context.Context) ([]core.Credit, error)
	GetByUser(ctx context.Context, idUser int64) ([]core.Credit, error)
	GetById(ctx context.Context, id int64) (*core.Credit, error)
	Create(ctx context.Context, input *core.Credit) (*core.Credit, error)
	Repay(ctx context.Context, id int64, amount int64) (*core.Credit, error)
	Delete(ctx context.Context, id int64) error
}
