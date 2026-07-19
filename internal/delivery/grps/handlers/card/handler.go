package card

import (
	"context"

	card "github.com/Suinar/Bank-proto/repository/card"
	common "github.com/Suinar/Bank-proto/repository/common"
	service "github.com/kVinsom/Bank-repository-service/internal/services/card"
)

// CardHandler adapts gRPC requests to the application service contract.
type CardHandler struct {
	card.UnimplementedCardRepositoryServer
	service service.ICardService
}

// NewCardHandler creates a ready-to-use card handler.
func NewCardHandler(service service.ICardService) *CardHandler {
	return &CardHandler{service: service}
}

// GetAll returns all records available through CardHandler.
func (h *CardHandler) GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error) {
	return h.service.GetAll(ctx, req)
}

// GetByUser returns records matching the requested user lookup.
func (h *CardHandler) GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error) {
	return h.service.GetByUser(ctx, req)
}

// GetById returns records matching the requested id lookup.
func (h *CardHandler) GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	return h.service.GetById(ctx, req)
}

// GetByNumber returns records matching the requested number lookup.
func (h *CardHandler) GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error) {
	return h.service.GetByNumber(ctx, req)
}

// Blocking moves the requested record to its blocked state through CardHandler.
func (h *CardHandler) Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	return h.service.Blocking(ctx, req)
}

// Create persists a new record through CardHandler.
func (h *CardHandler) Create(ctx context.Context, req *card.Card) (*card.Card, error) {
	return h.service.Create(ctx, req)
}

// Delete removes the requested record through CardHandler.
func (h *CardHandler) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	return h.service.Delete(ctx, req)
}
