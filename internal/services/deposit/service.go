package deposit

import (
	"context"
	"github.com/Suinar/Bank-proto/repository/common"
	"github.com/Suinar/Bank-proto/repository/deposit"
	"github.com/kVinsom/Bank-repository-service/pkg/core"

	repository "github.com/kVinsom/Bank-repository-service/internal/repositories/postgres/deposit"
)

// DepositService coordinates the application use cases for its domain.
type DepositService struct {
	repo repository.IDepositRepository
}

// NewDepositService creates a ready-to-use deposit service.
func NewDepositService(repo repository.IDepositRepository) *DepositService {
	return &DepositService{repo: repo}
}

// GetAll returns all records available through DepositService.
func (s *DepositService) GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error) {
	deposits, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &deposit.DepositList{
		Deposits: make([]*deposit.Deposit, len(deposits)),
	}

	for i, d := range deposits {
		res.Deposits[i] = s.toProto(&d)
	}

	return res, nil
}

// GetByUser returns records matching the requested user lookup.
func (s *DepositService) GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error) {
	deposits, err := s.repo.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &deposit.DepositList{
		Deposits: make([]*deposit.Deposit, len(deposits)),
	}

	for i, d := range deposits {
		res.Deposits[i] = s.toProto(&d)
	}

	return res, nil
}

// GetById returns records matching the requested id lookup.
func (s *DepositService) GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error) {
	d, err := s.repo.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(d), nil
}

// Create persists a new record through DepositService.
func (s *DepositService) Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error) {
	d, err := s.repo.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(d), nil
}

// Replenish adds funds to the requested deposit through DepositService.
func (s *DepositService) Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error) {
	d, err := s.repo.Replenish(ctx, req.Id, req.Amount)
	if err != nil {
		return nil, err
	}

	return s.toProto(d), nil
}

// Delete removes the requested record through DepositService.
func (s *DepositService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	err := s.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// toProto maps the domain model to its protobuf representation.
func (s *DepositService) toProto(input *core.Deposit) *deposit.Deposit {
	if input == nil {
		return nil
	}

	return &deposit.Deposit{
		Id:           input.Id,
		UserId:       input.UserId,
		CurrencyId:   input.CurrencyId,
		Amount:       input.Amount,
		InterestRate: input.InterestRate,
		TermMonths:   int32(input.TermMonths),
		Status:       deposit.DepositStatus(input.Status),
	}
}

// fromProto maps a protobuf message to the domain model.
func (s *DepositService) fromProto(input *deposit.Deposit) *core.Deposit {
	if input == nil {
		return nil
	}

	return &core.Deposit{
		Id:           input.Id,
		UserId:       input.UserId,
		CurrencyId:   input.CurrencyId,
		Amount:       input.Amount,
		InterestRate: input.InterestRate,
		TermMonths:   int8(input.TermMonths),
		Status:       core.DepositStatus(input.Status),
	}
}
