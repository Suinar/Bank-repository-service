package fixture

import common "Bank-repository-service/proto/repository/common"

func NewEmpty() *common.Empty {
	return &common.Empty{}
}

func NewIdRequest() *common.IdRequest {
	return &common.IdRequest{
		Id: TestId,
	}
}

func NewUserIdRequest() *common.UserIdRequest {
	return &common.UserIdRequest{
		UserId: TestId,
	}
}

func NewDeleteResponse() *common.DeleteResponse {
	return &common.DeleteResponse{
		EntityId: TestId,
	}
}

func NewAmountRequest(amount int64) *common.AmountRequest {
	return &common.AmountRequest{
		Id:     TestId,
		Amount: amount,
	}
}

func String(v string) *string { return &v }

func Int32(v int32) *int32 { return &v }
