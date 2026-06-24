package service

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
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
	userId, err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.DeleteResponse{UserId: userId}, nil
}

func (s *CreditService) toProto(c *core.Credit) *credit.Credit {
	if c == nil {
		return nil
	}
	return &credit.Credit{
		Id:             c.Id,
		UserId:         c.UserId,
		CurrencyId:     c.CurrencyId,
		Amount:         c.Amount,
		InterestRate:   c.InterestRate,
		TermMonths:     int32(c.TermMonths),
		MonthlyPayment: c.MonthlyPayment,
		Status:         credit.CreditStatus(c.Status),
	}
}

func (s *CreditService) fromProto(c *credit.Credit) *core.Credit {
	if c == nil {
		return nil
	}
	return &core.Credit{
		Id:             c.Id,
		UserId:         c.UserId,
		CurrencyId:     c.CurrencyId,
		Amount:         c.Amount,
		InterestRate:   c.InterestRate,
		TermMonths:     int8(c.TermMonths),
		MonthlyPayment: c.MonthlyPayment,
		Status:         core.CreditStatus(c.Status),
	}
}
