package account

import (
	"context"

	account "github.com/Suinar/Bank-proto/repository/account"
	common "github.com/Suinar/Bank-proto/repository/common"
	repository "github.com/kVinsom/Bank-repository-service/internal/repository/postgres_db/account"
	core "github.com/kVinsom/Bank-repository-service/pkg/core"
)

// AccountService coordinates the application use cases for its domain.
type AccountService struct {
	repository repository.IAccountRepository
}

// NewAccountService creates a ready-to-use account service.
func NewAccountService(repository repository.IAccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

// GetAll returns all records available through AccountService.
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

// GetByUser returns records matching the requested user lookup.
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

// GetById returns records matching the requested id lookup.
func (s *AccountService) GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

// Create persists a new record through AccountService.
func (s *AccountService) Create(ctx context.Context, req *account.Account) (*account.Account, error) {
	acc, err := s.repository.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

// Blocking moves the requested record to its blocked state through AccountService.
func (s *AccountService) Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.Blocking(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

// Close moves the requested account to its closed state through AccountService.
func (s *AccountService) Close(ctx context.Context, req *common.IdRequest) (*account.Account, error) {
	acc, err := s.repository.Close(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

// Update applies the requested changes through AccountService.
func (s *AccountService) Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error) {
	acc, err := s.repository.Update(ctx, req.Id, &core.AccountUpdateInput{
		Name: req.Input.Name,
	})
	if err != nil {
		return nil, err
	}

	return s.toProto(acc), nil
}

// Delete removes the requested record through AccountService.
func (s *AccountService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// toProto maps the domain model to its protobuf representation.
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

// fromProto maps a protobuf message to the domain model.
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
