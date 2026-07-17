package core

// CardStatus identifies the lifecycle state of the domain entity.
type CardStatus int

const (
	CardStatusClosed CardStatus = iota
	CardStatusActive
	CardStatusBlocked
	StatusExpired
)

// Card represents the persistent domain state of a bank card.
type Card struct {
	Id        int64 `json:"id" db:"id"`
	UserId    int64 `json:"user_id" db:"user_id"`
	AccountId int64 `json:"account_id" db:"account_id"`

	Number string `json:"number" db:"number"`

	ExpiryMonth int8 `json:"expiry_month" db:"expiry_month"`
	ExpiryYear  int8 `json:"expiry_year" db:"expiry_year"`

	Status CardStatus `json:"status" db:"status"`
}

// CardCreateInput carries validated fields for a domain mutation.
type CardCreateInput struct {
	UserId    int64 `json:"user_id" db:"user_id"`
	AccountId int64 `json:"account_id" db:"account_id"`
}
