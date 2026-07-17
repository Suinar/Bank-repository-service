package core

// AccountStatus identifies the lifecycle state of the domain entity.
type AccountStatus int

const (
	AccountStatusClosed AccountStatus = iota
	AccountStatusActive
	AccountStatusBlocked
)

// Account represents the persistent domain state of a bank account.
type Account struct {
	Id         int64 `json:"id" db:"id"`
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Name string `json:"name" db:"name"`

	Balance int64 `json:"balance" db:"balance"`

	Status AccountStatus `json:"status" db:"status"`
}

// AccountCreateInput carries validated fields for a domain mutation.
type AccountCreateInput struct {
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Name string `json:"name" db:"name"`
}

// AccountUpdateInput carries validated fields for a domain mutation.
type AccountUpdateInput struct {
	Name *string `json:"name" db:"name"`
}
