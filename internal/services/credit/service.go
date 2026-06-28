package credit

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db/credit"
	core "Bank-repository-service/pkg/core"
	common "Bank-repository-service/proto/repository/common"
	credit "Bank-repository-service/proto/repository/credit"
)

type CreditService struct {
	repository repository.ICreditRepository
}

func NewCreditService(repository repository.ICreditRepository) *CreditService {
	return &CreditService{repository: repository}
}

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

func (s *CreditService) GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error) {
	c, err := s.repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

func (s *CreditService) Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error) {
	c, err := s.repository.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

func (s *CreditService) Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error) {
	c, err := s.repository.Repay(ctx, req.Id, int(req.Amount))
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

func (s *CreditService) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	entityId, err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.DeleteResponse{EntityId: entityId}, nil
}

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
