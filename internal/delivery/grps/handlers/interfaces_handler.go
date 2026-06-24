package handler

import (
	"context"

	account "Bank-repository-service/proto/repository/account"
	card "Bank-repository-service/proto/repository/card"
	common "Bank-repository-service/proto/repository/common"
	credit "Bank-repository-service/proto/repository/credit"
	currency "Bank-repository-service/proto/repository/currency"
	deposit "Bank-repository-service/proto/repository/deposit"
	user "Bank-repository-service/proto/repository/user"
)

type IAccountHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*account.AccountList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*account.AccountList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Create(ctx context.Context, req *account.Account) (*account.Account, error)
	Blocking(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Close(ctx context.Context, req *common.IdRequest) (*account.Account, error)
	Update(ctx context.Context, req *account.UpdateAccountRequest) (*account.Account, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}

type ICardHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error)
	GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error)
	Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error)
	Create(ctx context.Context, req *card.Card) (*card.Card, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}

type ICreditHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*credit.CreditList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*credit.CreditList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*credit.Credit, error)
	Create(ctx context.Context, req *credit.Credit) (*credit.Credit, error)
	Repay(ctx context.Context, req *common.AmountRequest) (*credit.Credit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}

type ICurrencyHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*currency.CurrencyList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*currency.Currency, error)
	GetByIso(ctx context.Context, req *currency.IsoCodeRequest) (*currency.Currency, error)
	GetBySymbol(ctx context.Context, req *currency.SymbolRequest) (*currency.Currency, error)
	Create(ctx context.Context, req *currency.Currency) (*currency.Currency, error)
	Update(ctx context.Context, req *currency.UpdateCurrencyRequest) (*currency.Currency, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}

type IDepositHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*deposit.DepositList, error)
	GetByUser(ctx context.Context, req *common.UserIdRequest) (*deposit.DepositList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*deposit.Deposit, error)
	Create(ctx context.Context, req *deposit.Deposit) (*deposit.Deposit, error)
	Replenish(ctx context.Context, req *common.AmountRequest) (*deposit.Deposit, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error)
}

type IUserHandler interface {
	GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error)
	GetById(ctx context.Context, req *common.IdRequest) (*user.User, error)
	GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error)
	GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error)
	Create(ctx context.Context, req *user.User) (*user.User, error)
	ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (*common.Empty, error)
	Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error)
	Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error)
}
