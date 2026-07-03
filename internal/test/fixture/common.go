package fixture

import (
	common "Bank-repository-service/proto/repository/common"
)

func NewEmptyProto() *common.Empty {
	return &common.Empty{}
}

func NewIdRequestProto() *common.IdRequest {
	return &common.IdRequest{
		Id: TestId,
	}
}

func NewUserIdRequestProto() *common.UserIdRequest {
	return &common.UserIdRequest{
		UserId: TestId,
	}
}

func NewDeleteResponseProto() *common.DeleteResponse {
	return &common.DeleteResponse{
		EntityId: TestId,
	}
}

func NewAmountRequestProto(amount int64) *common.AmountRequest {
	return &common.AmountRequest{
		Id:     TestId,
		Amount: amount,
	}
}
