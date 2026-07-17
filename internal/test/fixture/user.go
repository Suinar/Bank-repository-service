package fixture

import (
	user "github.com/Suinar/Bank-proto/repository/user"
	"github.com/Suinar/Bank-repository-service/pkg/core"
)

// NewUserCore creates a ready-to-use user core.
func NewUserCore() core.User {
	return core.User{
		Id:          TestId,
		FirstName:   TestFirstName,
		MiddleName:  StringPointer(TestMidlName),
		LastName:    TestLastName,
		Email:       TestEmail,
		PhoneNumber: TestPhoneNumber,
	}
}

// NewUserCoreInputIdAndEmailAndPhoneNumber creates a ready-to-use user core input id and email and phone number.
func NewUserCoreInputIdAndEmailAndPhoneNumber(id int64, email string, phoneNumber string) core.User {
	return core.User{
		Id:          id,
		FirstName:   TestFirstName,
		MiddleName:  StringPointer(TestMidlName),
		LastName:    TestLastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}

// NewUserProto creates a ready-to-use user proto.
func NewUserProto() *user.User {
	return &user.User{
		Id:          TestId,
		FirstName:   TestFirstName,
		MiddleName:  StringPointer(TestMidlName),
		LastName:    TestLastName,
		Email:       TestEmail,
		PhoneNumber: TestPhoneNumber,
	}
}

// NewUserUpdateInputCore creates a ready-to-use user update input core.
func NewUserUpdateInputCore() *core.UserUpdateInput {
	return &core.UserUpdateInput{
		FirstName:  StringPointer(TestFirstName),
		MiddleName: StringPointer(TestMidlName),
		LastName:   StringPointer(TestLastName),
	}
}

// NewUserListProto creates a ready-to-use user list proto.
func NewUserListProto() *user.UserList {
	return &user.UserList{
		Users: []*user.User{
			NewUserProto(),
			NewUserProto(),
		},
	}
}

// NewEmailRequestProto creates a ready-to-use email request proto.
func NewEmailRequestProto() *user.EmailRequest {
	return &user.EmailRequest{
		Email: TestEmail,
	}
}

// NewPhoneNumberRequestProto creates a ready-to-use phone number request proto.
func NewPhoneNumberRequestProto() *user.PhoneNumberRequest {
	return &user.PhoneNumberRequest{
		PhoneNumber: TestPhoneNumber,
	}
}

// NewChangePasswordRequestProto creates a ready-to-use change password request proto.
func NewChangePasswordRequestProto() *user.ChangePasswordRequest {
	return &user.ChangePasswordRequest{
		Id:          1,
		NewPassword: TestPassword,
	}
}

// NewUserUpdateInputProto creates a ready-to-use user update input proto.
func NewUserUpdateInputProto() *user.UserUpdateInput {
	return &user.UserUpdateInput{
		FirstName:  StringPointer(TestFirstName),
		MiddleName: StringPointer(TestMidlName),
		LastName:   StringPointer(TestLastName),
	}
}

// NewUpdateUserRequestProto creates a ready-to-use update user request proto.
func NewUpdateUserRequestProto() *user.UpdateUserRequest {
	return &user.UpdateUserRequest{
		Id:    1,
		Input: NewUserUpdateInputProto(),
	}
}
