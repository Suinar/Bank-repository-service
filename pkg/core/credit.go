package core

// CreditStatus identifies the lifecycle state of the domain entity.
type CreditStatus int

const (
	CreditStatusClosed CreditStatus = iota
	CreditStatusActive
	CreditStatusRejected
)

// Credit represents the persistent domain state of a bank credit.
type Credit struct {
	Id         int64 `json:"id" db:"id"`
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Amount       int64   `json:"amount" db:"amount"`
	InterestRate float32 `json:"interest_rate" db:"interest_rate"`

	TermMonths     int8  `json:"term_months" db:"term_months"`
	MonthlyPayment int64 `json:"monthly_payment" db:"monthly_payment"`

	Status CreditStatus `json:"status" db:"status"`
}

// CreditCreateInput carries validated fields for a domain mutation.
type CreditCreateInput struct {
	UserId     int64 `json:"user_id" db:"user_id"`
	CurrencyId int64 `json:"currency_id" db:"currency_id"`

	Amount     int64 `json:"amount" db:"amount"`
	TermMonths int8  `json:"term_months" db:"term_months"`
}
