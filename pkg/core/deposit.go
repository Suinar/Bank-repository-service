package core

// DepositStatus identifies the lifecycle state of the domain entity.
type DepositStatus int

const (
	DepositStatusClosed DepositStatus = iota
	DepositStatusActive
	DepositStatusRejected
)

// Deposit represents the persistent domain state of a bank deposit.
type Deposit struct {
	Id         int64 `json:"id" db:"id"`
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Amount       int64   `json:"amount" db:"amount"`
	InterestRate float32 `json:"interest_rate" db:"interest_rate"`

	TermMonths int8 `json:"term_months" db:"term_months"`

	Status DepositStatus `json:"status" db:"status"`
}

// DepositCreateInput carries validated fields for a domain mutation.
type DepositCreateInput struct {
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Amount int64 `json:"amount" db:"amount"`

	TermMonths int8 `json:"term_months" db:"term_months"`
}
