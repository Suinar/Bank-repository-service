package data

import "Bank-repository-service/proto/repository/user"

func NewUser() *user.User {
	return &user.User{
		Id:           1,
		FirstName:    "Test",
		LastName:     "User",
		Email:        "test_user@gmail.com",
		PhoneNumber:  "+380123456789",
		PasswordHash: "DM1LVI8kjGH(f#21FH",
	}
}
