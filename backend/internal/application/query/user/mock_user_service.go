package user

import (
	"context"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockUserService implements a mock for the user.ServiceInterface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, firebaseUID, email, name string, profileImageURL *string) (*user.User, error) {
	args := m.Called(ctx, firebaseUID, email, name, profileImageURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) GetUser(ctx context.Context, id uuid.UUID) (*user.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) GetUserByFirebaseUID(ctx context.Context, firebaseUID string) (*user.User, error) {
	args := m.Called(ctx, firebaseUID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id uuid.UUID, email *string, name *string, profileImageURL *string) (*user.User, error) {
	args := m.Called(ctx, id, email, name, profileImageURL)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}