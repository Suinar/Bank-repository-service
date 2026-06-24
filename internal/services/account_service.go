package service

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
	account "Bank-repository-service/proto/repository/account"
	common "Bank-repository-service/proto/repository/common"
)

type AccountService struct {
	repository repository.IAccountRepository
}

func NewAccountService(repository repository.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) GetAll(ctx context.Context, req *common.Empty) (*account.AccountList, error) {
	accounts, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &account.AccountList{
		Accounts: make([]*account.Account, len(accounts)),
	}

	for i, acc := range accounts {
		res.Accounts[i] = s.toProto(&acc)
	}

	return res, nil
}

func (s *AccountService) GetByUser(ctx context.Context, req *common.UserIdRequest) (*account.AccountList, error) {
	accounts, err := s.repository.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &account.AccountList{
		Accounts: make([]*account.Account, len(accounts)),
	}

	for i, acc := range accounts {
		res.Accounts[i] = s.toProto(&acc)
	}

	return res, nil
}

func (s *AccountService) GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

func (s *AccountService) Create(ctx context.Context, req *account.Account) (*account.Account, error) {
	acc, err := s.repository.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

func (s *AccountService) Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.Blocking(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(&acc), nil
}

func (s *AccountService) Close(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.Close(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(&acc), nil
}

func (s *AccountService) Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error) {
	var name *string
	if req.Input.Name != nil {
		name = req.Input.Name
	}

	acc, err := s.repository.Update(ctx, req.Id, &core.AccountUpdateInput{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

func (s *AccountService) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	userId, err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.DeleteResponse{UserId: userId}, nil
}

func (s *AccountService) toProto(acc *core.Account) *account.Account {
	if acc == nil {
		return nil
	}
	return &account.Account{
		Id:         acc.Id,
		UserId:     acc.UserId,
		CurrencyId: acc.CurrencyId,
		Name:       acc.Name,
		Balance:    acc.Balance,
		Status:     account.AccountStatus(acc.Status),
	}
}

func (s *AccountService) fromProto(acc *account.Account) *core.Account {
	if acc == nil {
		return nil
	}
	return &core.Account{
		Id:         acc.Id,
		UserId:     acc.UserId,
		CurrencyId: acc.CurrencyId,
		Name:       acc.Name,
		Balance:    acc.Balance,
		Status:     core.AccountStatus(acc.Status),
	}
}
