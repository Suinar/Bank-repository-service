package user

import (
	"Bank-repository-service/pkg/core"
	"Bank-repository-service/proto/repository/common"
	"Bank-repository-service/proto/repository/user"
	"context"

	repository "Bank-repository-service/internal/repository/postgres_db/user"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll(ctx context.Context, req *common.Empty) (*user.UserList, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &user.UserList{
		Users: make([]*user.User, len(users)),
	}

	for i, u := range users {
		res.Users[i] = s.toProto(&u)
	}

	return res, nil
}

func (s *UserService) GetById(ctx context.Context, req *common.IdRequest) (*user.User, error) {
	user, err := s.repo.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return s.toProto(user), nil
}

func (s *UserService) GetByEmail(ctx context.Context, req *user.EmailRequest) (*user.User, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return s.toProto(user), nil
}

func (s *UserService) GetByPhoneNumber(ctx context.Context, req *user.PhoneNumberRequest) (*user.User, error) {
	user, err := s.repo.GetByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return nil, err
	}

	return s.toProto(user), nil
}

func (s *UserService) Create(ctx context.Context, req *user.User) (*user.User, error) {
	user, err := s.repo.Create(ctx, s.fromProto(req))
	if err != nil {
		return nil, err
	}

	return s.toProto(user), nil
}

func (s *UserService) ChangePassword(ctx context.Context, req *user.ChangePasswordRequest) (*common.Empty, error) {
	err := s.repo.ChangePassword(ctx, req.Id, req.NewPassword)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

func (s *UserService) Update(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	u, err := s.repo.Update(ctx, req.Id, &core.UserUpdateInput{
		FirstName:  req.Input.FirstName,
		MiddleName: req.Input.MiddleName,
		LastName:   req.Input.LastName,
	})

	if err != nil {
		return nil, err
	}

	return s.toProto(u), nil
}

func (s *UserService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	err := s.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

func (s *UserService) toProto(input *core.User) *user.User {
	if input == nil {
		return nil
	}

	return &user.User{
		Id:           input.Id,
		FirstName:    input.FirstName,
		MiddleName:   input.MiddleName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		PasswordHash: input.PasswordHash,
	}
}

func (s *UserService) fromProto(input *user.User) *core.User {
	if input == nil {
		return nil
	}

	return &core.User{
		Id:           input.Id,
		FirstName:    input.FirstName,
		MiddleName:   input.MiddleName,
		LastName:     input.LastName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		PasswordHash: input.PasswordHash,
	}
}
