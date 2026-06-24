package service

import (
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db"
	core "Bank-repository-service/pkg/core"
	card "Bank-repository-service/proto/repository/card"
	common "Bank-repository-service/proto/repository/common"
)

type CardService struct {
	repository repository.ICardRepository
}

func NewCardService(repository repository.ICardRepository) *CardService {
	return &CardService{repository: repository}
}

func (s *CardService) GetAll(ctx context.Context, req *common.Empty) (*card.CardList, error) {
	cards, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &card.CardList{
		Cards: make([]*card.Card, len(cards)),
	}

	for i, c := range cards {
		res.Cards[i] = s.toProto(&c)
	}

	return res, nil
}

func (s *CardService) GetByUser(ctx context.Context, req *common.UserIdRequest) (*card.CardList, error) {
	cards, err := s.repository.GetByUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res := &card.CardList{
		Cards: make([]*card.Card, len(cards)),
	}

	for i, c := range cards {
		res.Cards[i] = s.toProto(&c)
	}

	return res, nil
}

func (s *CardService) GetById(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	c, err := s.repository.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

func (s *CardService) GetByNumber(ctx context.Context, req *card.CardNumberRequest) (*card.Card, error) {
	c, err := s.repository.GetByNumber(ctx, req.Number)
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

func (s *CardService) Blocking(ctx context.Context, req *common.IdRequest) (*card.Card, error) {
	c, err := s.repository.Blocking(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(&c), nil
}

func (s *CardService) Create(ctx context.Context, req *card.Card) (*card.Card, error) {
	c, err := s.repository.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(c), nil
}

func (s *CardService) Delete(ctx context.Context, req *common.IdRequest) (*common.DeleteResponse, error) {
	userId, err := s.repository.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.DeleteResponse{UserId: userId}, nil
}

func (s *CardService) toProto(c *core.Card) *card.Card {
	if c == nil {
		return nil
	}
	return &card.Card{
		Id:          c.Id,
		UserId:      c.UserId,
		AccountId:   c.AccountId,
		Number:      c.Number,
		ExpiryMonth: int32(c.ExpiryMonth),
		ExpiryYear:  int32(c.ExpiryYear),
		Status:      card.CardStatus(c.Status),
	}
}

func (s *CardService) fromProto(c *card.Card) *core.Card {
	if c == nil {
		return nil
	}
	return &core.Card{
		Id:          c.Id,
		UserId:      c.UserId,
		AccountId:   c.AccountId,
		Number:      c.Number,
		ExpiryMonth: int8(c.ExpiryMonth),
		ExpiryYear:  int8(c.ExpiryYear),
		Status:      core.CardStatus(c.Status),
	}
}
