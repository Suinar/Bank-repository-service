package credit

import (
	"context"

	common "github.com/Suinar/Bank-proto/repository/common"
	credit "github.com/Suinar/Bank-proto/repository/credit"
	repository "github.com/Suinar/Bank-repository-service/internal/repository/postgres_db/credit"
	core "github.com/Suinar/Bank-repository-service/pkg/core"
)

// CreditService coordinates the application use cases for its domain.
type CreditService struct {
	repository repository.ICreditRepository
}

// NewCreditService creates a ready-to-use credit service.
func NewCreditService(repository repository.ICreditRepository) *CreditService {
	return &CreditService{repository: repository}
}

// GetAll returns all records available through CreditService.
func (s *CreditService) GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error) {
	credits, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &credit.CreditList{
		Credits: make([]*credit.Credit, len(credits)),
	}

	for i, c := range credits {
		res.Credits[i] = s.toProto(&c)
	}

	return res, nil
}

// GetByUser returns records matching the requested user lookup.
func (s *CreditService) GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error) {
	credits, err := s.repository.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &credit.CreditList{
		Credits: make([]*credit.Credit, len(credits)),
	}

	for i, c := range credits {
		res.Credits[i] = s.toProto(&c)
	}

	return res, nil
}

// GetById returns records matching the requested id lookup.
func (s *CreditService) GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error) {
	c, err := s.repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

// Create persists a new record through CreditService.
func (s *CreditService) Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error) {
	c, err := s.repository.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

// Repay applies a repayment to the requested credit through CreditService.
func (s *CreditService) Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error) {
	c, err := s.repository.Repay(ctx, req.Id, req.Amount)
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

// Delete removes the requested record through CreditService.
func (s *CreditService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// toProto maps the domain model to its protobuf representation.
func (s *CreditService) toProto(input *core.Credit) *credit.Credit {
	if input == nil {
		return nil
	}

	return &credit.Credit{
		Id:             input.Id,
		UserId:         input.UserId,
		CurrencyId:     input.CurrencyId,
		Amount:         input.Amount,
		InterestRate:   input.InterestRate,
		TermMonths:     int32(input.TermMonths),
		MonthlyPayment: input.MonthlyPayment,
		Status:         credit.CreditStatus(input.Status),
	}
}

// fromProto maps a protobuf message to the domain model.
func (s *CreditService) fromProto(input *credit.Credit) *core.Credit {
	if input == nil {
		return nil
	}

	return &core.Credit{
		Id:             input.Id,
		UserId:         input.UserId,
		CurrencyId:     input.CurrencyId,
		Amount:         input.Amount,
		InterestRate:   input.InterestRate,
		TermMonths:     int8(input.TermMonths),
		MonthlyPayment: input.MonthlyPayment,
		Status:         core.CreditStatus(input.Status),
	}
}
