package fixture

import (
	"github.com/Suinar/Bank-repository-service/pkg/core"
	user "github.com/Suinar/Bank-proto/repository/user"
)

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

func NewUserUpdateInputCore() *core.UserUpdateInput {
	return &core.UserUpdateInput{
		FirstName:  StringPointer(TestFirstName),
		MiddleName: StringPointer(TestMidlName),
		LastName:   StringPointer(TestLastName),
	}
}

func NewUserListProto() *user.UserList {
	return &user.UserList{
		Users: []*user.User{
			NewUserProto(),
			NewUserProto(),
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

func NewUserUpdateInputProto() *user.UserUpdateInput {
	return &user.UserUpdateInput{
		FirstName:  StringPointer(TestFirstName),
		MiddleName: StringPointer(TestMidlName),
		LastName:   StringPointer(TestLastName),
	}
}

func NewUpdateUserRequestProto() *user.UpdateUserRequest {
	return &user.UpdateUserRequest{
		Id:    1,
		Input: NewUserUpdateInputProto(),
	}
}



