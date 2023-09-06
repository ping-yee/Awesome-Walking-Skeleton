package resolvers

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/taimoor99/three-tier-golang/app/entities"
)

type MockedUserModels struct {
	mock.Mock
}

func (s *MockedUserModels) CreateUser(ctx context.Context, user entities.UsersCreateInput) (string, error) {
	args := s.Called(ctx, user)
	return args.Get(1).(string), args.Error(1)
}

func (s *MockedUserModels) GetAllUsers(ctx context.Context, limit, offset int64) ([]entities.Users, error) {
	args := s.Called(ctx, limit, offset)
	return args.Get(1).([]entities.Users), args.Error(1)
}

func (s *MockedUserModels) FindUserByID(ctx context.Context, id string) (*entities.Users, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(*entities.Users), args.Error(1)
}