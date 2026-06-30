package fixture

import "Bank-repository-service/proto/repository/user"

func NewUser(firstName string, mildName string, lastName string) *user.User {
	return &user.User{
		Id:           TestId,
		FirstName:    firstName,
		MiddleName:   String(mildName),
		LastName:     lastName,
		Email:        TestEmail,
		PhoneNumber:  TestPhoneNumber,
		PasswordHash: TestPasswordHash,
	}
}

func NewUserList(
	firstNameFirst string, firstNameSecond string,
	mildNameFirst string, mildNameSecond string,
	lastNameFirst string, lastNameSecond string) *user.UserList {
	return &user.UserList{
		Users: []*user.User{
			NewUser(firstNameFirst, mildNameFirst, lastNameFirst),
			NewUser(firstNameSecond, mildNameSecond, lastNameSecond),
		},
	}
}

func NewEmailRequest() *user.EmailRequest {
	return &user.EmailRequest{
		Email: TestEmail,
	}
}

func NewPhoneNumberRequest() *user.PhoneNumberRequest {
	return &user.PhoneNumberRequest{
		PhoneNumber: TestPhoneNumber,
	}
}

func NewChangePasswordRequest() *user.ChangePasswordRequest {
	return &user.ChangePasswordRequest{
		Id:          1,
		NewPassword: TestPassword,
	}
}

func NewUserUpdateInput(firstName string, mildName string, lastName string) *user.UserUpdateInput {
	return &user.UserUpdateInput{
		FirstName:  String(firstName),
		MiddleName: String(mildName),
		LastName:   String(lastName),
	}
}

func NewUpdateUserRequest(firstName string, mildName string, lastName string) *user.UpdateUserRequest {
	return &user.UpdateUserRequest{
		Id:    1,
		Input: NewUserUpdateInput(firstName, mildName, lastName),
	}
}
