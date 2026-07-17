package core

// User represents the persistent domain state of a bank user.
type User struct {
	Id int64 `json:"id" db:"id"`

	FirstName  string  `json:"first_name" db:"first_name"`
	MiddleName *string `json:"middle_name" db:"middle_name"`
	LastName   string  `json:"last_name" db:"last_name"`

	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`

	PasswordHash string `json:"password_hash" db:"password_hash"`
}

// UserCreateInput carries validated fields for a domain mutation.
type UserCreateInput struct {
	FirstName  string  `json:"first_name" db:"first_name"`
	MiddleName *string `json:"middle_name" db:"middle_name"`
	LastName   string  `json:"last_name" db:"last_name"`

	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`

	PasswordHash string `json:"password_hash" db:"password_hash"`
}

// UserUpdateInput carries validated fields for a domain mutation.
type UserUpdateInput struct {
	FirstName  *string `json:"first_name" db:"first_name"`
	MiddleName *string `json:"middle_name" db:"middle_name"`
	LastName   *string `json:"last_name" db:"last_name"`
}
