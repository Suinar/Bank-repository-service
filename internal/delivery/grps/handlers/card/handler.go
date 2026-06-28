package card

import (
	"context"

	service "Bank-repository-service/internal/services/card"
	card "Bank-repository-service/proto/repository/card"
	common "Bank-repository-service/proto/repository/common"
)

type CardHandler struct {
	card.UnimplementedCardRepositoryServer
	service service.ICardService
}

func NewCardHandler(service service.ICardService) *CardHandler {
	return &CardHandler{service: service}
}

func (h *CardHandler) GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error) {
	return h.service.GetAll(ctx, req)
}

func (h *CardHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error) {
	return h.service.GetByUser(ctx, req)
}

func (h *CardHandler) GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	return h.service.GetById(ctx, req)
}

func (h *CardHandler) GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error) {
	return h.service.GetByNumber(ctx, req)
}

func (h *CardHandler) Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	return h.service.Blocking(ctx, req)
}

func (h *CardHandler) Create(ctx context.Context, req *card.Card) (*card.Card, error) {
	return h.service.Create(ctx, req)
}

func (h *CardHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	return h.service.Delete(ctx, req)
}
