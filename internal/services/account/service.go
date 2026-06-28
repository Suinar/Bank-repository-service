package account

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db/account"
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

	return s.toProto(acc), nil
}

func (s *AccountService) Close(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.Close(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

func (s *AccountService) Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error) {
	acc, err := s.repository.Update(ctx, req.Id, &core.AccountUpdateInput{
		Name: req.Input.Name,
	})
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

func (s *AccountService) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	entityId, err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.DeleteResponse{EntityId: entityId}, nil
}

func (s *AccountService) toProto(input *core.Account) *account.Account {
	if input == nil {
		return nil
	}

	return &account.Account{
		Id:         input.Id,
		UserId:     input.UserId,
		CurrencyId: input.CurrencyId,
		Name:       input.Name,
		Balance:    input.Balance,
		Status:     account.AccountStatus(input.Status),
	}
}

func (s *AccountService) fromProto(input *account.Account) *core.Account {
	if input == nil {
		return nil
	}

	return &core.Account{
		Id:         input.Id,
		UserId:     input.UserId,
		CurrencyId: input.CurrencyId,
		Name:       input.Name,
		Balance:    input.Balance,
		Status:     core.AccountStatus(input.Status),
	}
}
