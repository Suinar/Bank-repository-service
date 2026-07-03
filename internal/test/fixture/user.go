package fixture

import (
	"Bank-repository-service/pkg/core"
	user "Bank-repository-service/proto/repository/user"
)

func NewUserCore(firstName string, mildName string, lastName string) core.User {
	return core.User{
		Id:           TestId,
		FirstName:    firstName,
		MiddleName:   StringPointer(mildName),
		LastName:     lastName,
		Email:        TestEmail,
		PhoneNumber:  TestPhoneNumber,
		PasswordHash: TestPasswordHash,
	}
}

func NewUserProto(firstName string, mildName string, lastName string) *user.User {
	return &user.User{
		Id:           TestId,
		FirstName:    firstName,
		MiddleName:   StringPointer(mildName),
		LastName:     lastName,
		Email:        TestEmail,
		PhoneNumber:  TestPhoneNumber,
		PasswordHash: TestPasswordHash,
	}
}

func NewUserUpdateInputCore(firstName string, mildName string, lastName string) *core.UserUpdateInput {
	return &core.UserUpdateInput{
		FirstName:  StringPointer(firstName),
		MiddleName: StringPointer(mildName),
		LastName:   StringPointer(lastName),
	}
}

func NewUserListProto(
	firstNameFirst string, firstNameSecond string,
	mildNameFirst string, mildNameSecond string,
	lastNameFirst string, lastNameSecond string) *user.UserList {
	return &user.UserList{
		Users: []*user.User{
			NewUserProto(firstNameFirst, mildNameFirst, lastNameFirst),
			NewUserProto(firstNameSecond, mildNameSecond, lastNameSecond),
		},
	}
}

func NewEmailRequestProto() *user.EmailRequest {
	return &user.EmailRequest{
		Email: TestEmail,
	}
}

func NewPhoneNumberRequestProto() *user.PhoneNumberRequest {
	return &user.PhoneNumberRequest{
		PhoneNumber: TestPhoneNumber,
	}
}

func NewChangePasswordRequestProto() *user.ChangePasswordRequest {
	return &user.ChangePasswordRequest{
		Id:          1,
		NewPassword: TestPassword,
	}
}

func NewUserUpdateInputProto(firstName string, mildName string, lastName string) *user.UserUpdateInput {
	return &user.UserUpdateInput{
		FirstName:  StringPointer(firstName),
		MiddleName: StringPointer(mildName),
		LastName:   StringPointer(lastName),
	}
}

func NewUpdateUserRequestProto(firstName string, mildName string, lastName string) *user.UpdateUserRequest {
	return &user.UpdateUserRequest{
		Id:    1,
		Input: NewUserUpdateInputProto(firstName, mildName, lastName),
	}
}
