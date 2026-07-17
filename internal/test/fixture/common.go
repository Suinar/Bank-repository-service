package fixture

import (
	common "github.com/Suinar/Bank-proto/repository/common"
)

// NewEmptyProto creates a ready-to-use empty proto.
func NewEmptyProto() *common.Empty {
	return &common.Empty{}
}

// NewIdRequestProto creates a ready-to-use id request proto.
func NewIdRequestProto() *common.IdRequest {
	return &common.IdRequest{
		Id: TestId,
	}
}

// NewUserIdRequestProto creates a ready-to-use user id request proto.
func NewUserIdRequestProto() *common.UserIdRequest {
	return &common.UserIdRequest{
		UserId: TestId,
	}
}

// NewAmountRequestProto creates a ready-to-use amount request proto.
func NewAmountRequestProto(amount int64) *common.AmountRequest {
	return &common.AmountRequest{
		Id:     TestId,
		Amount: amount,
	}
}
