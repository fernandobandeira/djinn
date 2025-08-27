package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"testing"
	"time"

	"github.com/fernandobandeira/djinn/backend/internal/domain/user"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/generated"
	"github.com/fernandobandeira/djinn/backend/internal/infrastructure/persistence/postgres/mocks"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_Create(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))

	tests := []struct {
		name       string
		user       *user.User
		setupMock  func(*mocks.MockQuerier)
		wantErr    bool
		errMessage string
	}{
		{
			name: "successful creation",
			user: &user.User{
				FirebaseUID:     "firebase123",
				Email:           "test@example.com",
				Name:            "Test User",
				ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
			},
			setupMock: func(mq *mocks.MockQuerier) {
				expectedParams := generated.CreateUserParams{
					FirebaseUid:     "firebase123",
					Email:           "test@example.com",
					Name:            "Test User",
					ProfileImageUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
				}
				
				now := time.Now()
				userID := uuid.New()
				
				mq.On("CreateUser", mock.Anything, expectedParams).Return(generated.User{
					ID:              pgtype.UUID{Bytes: userID, Valid: true},
					FirebaseUid:     "firebase123",
					Email:           "test@example.com",
					Name:            "Test User",
					ProfileImageUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
					CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "database error",
			user: &user.User{
				FirebaseUID: "firebase789",
				Email:       "test3@example.com",
				Name:        "Test User 3",
			},
			setupMock: func(mq *mocks.MockQuerier) {
				mq.On("CreateUser", mock.Anything, mock.Anything).Return(generated.User{}, sql.ErrConnDone)
			},
			wantErr:    true,
			errMessage: "failed to create user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQueries := &mocks.MockQuerier{}
			repo := NewUserRepository(mockQueries, logger)
			
			tt.setupMock(mockQueries)
			
			err := repo.Create(context.Background(), tt.user)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, uuid.Nil, tt.user.ID)
				assert.NotZero(t, tt.user.CreatedAt)
				assert.NotZero(t, tt.user.UpdatedAt)
			}
			
			mockQueries.AssertExpectations(t)
		})
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))
	userID := uuid.New()

	tests := []struct {
		name      string
		userID    uuid.UUID
		setupMock func(*mocks.MockQuerier)
		wantErr   bool
		isNotFound bool
	}{
		{
			name:   "successful retrieval",
			userID: userID,
			setupMock: func(mq *mocks.MockQuerier) {
				now := time.Now()
				
				mq.On("GetUser", mock.Anything, pgtype.UUID{Bytes: userID, Valid: true}).Return(generated.User{
					ID:              pgtype.UUID{Bytes: userID, Valid: true},
					FirebaseUid:     "firebase123",
					Email:           "test@example.com",
					Name:            "Test User",
					ProfileImageUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
					CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
					UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
				}, nil)
			},
			wantErr: false,
		},
		{
			name:   "user not found",
			userID: userID,
			setupMock: func(mq *mocks.MockQuerier) {
				mq.On("GetUser", mock.Anything, mock.Anything).Return(generated.User{}, pgx.ErrNoRows)
			},
			wantErr:    true,
			isNotFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQueries := &mocks.MockQuerier{}
			repo := NewUserRepository(mockQueries, logger)
			
			tt.setupMock(mockQueries)
			
			result, err := repo.GetByID(context.Background(), tt.userID)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.isNotFound {
					assert.ErrorIs(t, err, user.ErrUserNotFound)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, "firebase123", result.FirebaseUID)
				assert.Equal(t, "test@example.com", result.Email)
				assert.Equal(t, "Test User", result.Name)
			}
			
			mockQueries.AssertExpectations(t)
		})
	}
}

func TestUserRepository_GetByFirebaseUID(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(nil, nil))

	tests := []struct {
		name        string
		firebaseUID string
		setupMock   func(*mocks.MockQuerier)
		wantErr     bool
		isNotFound  bool
	}{
		{
			name:        "successful retrieval",
			firebaseUID: "firebase123",
			setupMock: func(mq *mocks.MockQuerier) {
				userID := uuid.New()
				now := time.Now()
				
				mq.On("GetUserByFirebaseUID", mock.Anything, "firebase123").Return(generated.User{
					ID:          pgtype.UUID{Bytes: userID, Valid: true},
					FirebaseUid: "firebase123",
					Email:       "test@example.com",
					Name:        "Test User",
					CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
					UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				}, nil)
			},
			wantErr: false,
		},
		{
			name:        "user not found",
			firebaseUID: "nonexistent",
			setupMock: func(mq *mocks.MockQuerier) {
				mq.On("GetUserByFirebaseUID", mock.Anything, "nonexistent").Return(generated.User{}, pgx.ErrNoRows)
			},
			wantErr:    true,
			isNotFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQueries := &mocks.MockQuerier{}
			repo := NewUserRepository(mockQueries, logger)
			
			tt.setupMock(mockQueries)
			
			result, err := repo.GetByFirebaseUID(context.Background(), tt.firebaseUID)
			
			if tt.wantErr {
				assert.Error(t, err)
				if tt.isNotFound {
					assert.ErrorIs(t, err, user.ErrUserNotFound)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.firebaseUID, result.FirebaseUID)
			}
			
			mockQueries.AssertExpectations(t)
		})
	}
}

func TestUserRepository_toDomainUser(t *testing.T) {
	repo := &UserRepository{}
	userID := uuid.New()
	now := time.Now()

	tests := []struct {
		name   string
		dbUser *generated.User
		want   *user.User
	}{
		{
			name: "with profile image",
			dbUser: &generated.User{
				ID:              pgtype.UUID{Bytes: userID, Valid: true},
				FirebaseUid:     "firebase123",
				Email:           "test@example.com",
				Name:            "Test User",
				ProfileImageUrl: pgtype.Text{String: "https://example.com/avatar.jpg", Valid: true},
				CreatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:       pgtype.Timestamptz{Time: now, Valid: true},
			},
			want: &user.User{
				ID:              userID,
				FirebaseUID:     "firebase123",
				Email:           "test@example.com",
				Name:            "Test User",
				ProfileImageURL: stringPtr("https://example.com/avatar.jpg"),
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		},
		{
			name: "without profile image",
			dbUser: &generated.User{
				ID:          pgtype.UUID{Bytes: userID, Valid: true},
				FirebaseUid: "firebase456",
				Email:       "test2@example.com",
				Name:        "Test User 2",
				CreatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
				UpdatedAt:   pgtype.Timestamptz{Time: now, Valid: true},
			},
			want: &user.User{
				ID:              userID,
				FirebaseUID:     "firebase456",
				Email:           "test2@example.com",
				Name:            "Test User 2",
				ProfileImageURL: nil,
				CreatedAt:       now,
				UpdatedAt:       now,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := repo.toDomainUser(tt.dbUser)
			
			assert.Equal(t, tt.want.ID, result.ID)
			assert.Equal(t, tt.want.FirebaseUID, result.FirebaseUID)
			assert.Equal(t, tt.want.Email, result.Email)
			assert.Equal(t, tt.want.Name, result.Name)
			assert.Equal(t, tt.want.CreatedAt.Unix(), result.CreatedAt.Unix())
			assert.Equal(t, tt.want.UpdatedAt.Unix(), result.UpdatedAt.Unix())
			
			if tt.want.ProfileImageURL != nil {
				require.NotNil(t, result.ProfileImageURL)
				assert.Equal(t, *tt.want.ProfileImageURL, *result.ProfileImageURL)
			} else {
				assert.Nil(t, result.ProfileImageURL)
			}
		})
	}
}

// Helper function
func stringPtr(s string) *string {
	return &s
}