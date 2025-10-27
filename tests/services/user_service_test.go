package services

import (
	"context"
	"errors"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFeatchAllUsers(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while fetching users", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetAllUsersFunc: func(ctx context.Context, skip int, limit int) ([]*entities.User, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewUserService(mockRepo)
		users, err := svc.GetAllUsers(ctx, 0, 10)

		// Assertions
		assert.Nil(t, users)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("successfully featched", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetAllUsersFunc: func(ctx context.Context, skip int, limit int) ([]*entities.User, error) {
				return []*entities.User{
					{ID: 1, Name: "Test Co", Email: "a@test.com"},
					{ID: 1, Name: "Another Co", Email: "b@test.com"},
				}, nil
			},
		}

		svc := services.NewUserService(mockRepo)
		users, err := svc.GetAllUsers(ctx, 0, 10)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.Equal(t, 2, len(users))
		assert.Equal(t, "Test Co", users[0].Name)
	})
}

func TestFindUserByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: user id not found", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			FindUserByIDFunc: func(ctx context.Context, userID int) (*entities.User, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewUserService(mockRepo)

		user, err := svc.FindUserByID(ctx, 1)

		// Assertions
		assert.Nil(t, user)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully returned", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			FindUserByIDFunc: func(ctx context.Context, userID int) (*entities.User, error) {
				return &entities.User{ID: userID, Name: "Test Co"}, nil
			},
		}

		svc := services.NewUserService(mockRepo)

		user, err := svc.FindUserByID(ctx, 1)

		// Assertions
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test Co", user.Name)
	})
}

func TestInsertUser(t *testing.T) {
	ctx := context.Background()

	t.Run("error: email already exists", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.User, error) {
				return &entities.User{ID: 1, Email: email}, nil
			},
		}

		svc := services.NewUserService(mockRepo)
		req := &requests.UserRequest{Name: "Test", Email: "x@test.com", Password: "qazwsx@123"}

		user, err := svc.InsertUser(ctx, req)

		// Assertions
		assert.Nil(t, user)
		assert.EqualError(t, err, "email already in use")
	})

	t.Run("successfully inserted", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.User, error) {
				return nil, nil
			},
			InsertUserFunc: func(ctx context.Context, user *entities.User) (*entities.User, error) {
				user.ID = 42
				return user, nil
			},
		}

		svc := services.NewUserService(mockRepo)
		req := &requests.UserRequest{Name: "NewCo", Email: "new@test.com", Password: "qazwsx@123"}

		user, err := svc.InsertUser(ctx, req)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, 42, user.ID)
	})
}

func TestUpdateUserByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: email already in use", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.User, error) {
				return &entities.User{ID: 2, Email: email}, nil // simulate existing user with same email
			},
			UpdateUserByIDFunc: func(ctx context.Context, userID int, updates map[string]interface{}) (*entities.User, error) {
				return nil, nil
			},
		}

		svc := services.NewUserService(mockRepo)
		newUser := &entities.User{ID: 1, Name: "NewCo", Email: "new@test.com"}
		user, err := svc.UpdateUserByID(ctx, 1, newUser)

		assert.Nil(t, user)
		assert.EqualError(t, err, "email already in use")
	})

	t.Run("error: database update failure", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.User, error) {
				return nil, nil // no conflict
			},
			UpdateUserByIDFunc: func(ctx context.Context, userID int, updates map[string]interface{}) (*entities.User, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewUserService(mockRepo)
		newUser := &entities.User{ID: 1, Name: "NewCo", Email: "new@test.com"}
		user, err := svc.UpdateUserByID(ctx, 1, newUser)

		assert.Nil(t, user)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully updated", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.User, error) {
				return nil, nil // no existing email
			},
			UpdateUserByIDFunc: func(ctx context.Context, userID int, updates map[string]interface{}) (*entities.User, error) {
				return &entities.User{
					ID:    42,
					Name:  updates["name"].(string),
					Email: updates["email"].(string),
				}, nil
			},
		}

		svc := services.NewUserService(mockRepo)
		newUser := &entities.User{Name: "NewCo", Email: "new@test.com"}
		user, err := svc.UpdateUserByID(ctx, 1, newUser)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 42, user.ID)
		assert.Equal(t, "NewCo", user.Name)
		assert.Equal(t, "new@test.com", user.Email)
	})
}

func TestDeleteUserByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: user id not found", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			DeleteUserByIDFunc: func(ctx context.Context, userID int) (int, error) {
				return 0, errors.New("database error")
			},
		}

		svc := services.NewUserService(mockRepo)
		isDeleted, err := svc.DeleteUserByID(ctx, 1)

		// Assertions
		assert.Equal(t, isDeleted, 0)
		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully deleted", func(t *testing.T) {
		mockRepo := &mocks.MockUserRepo{
			DeleteUserByIDFunc: func(ctx context.Context, userID int) (int, error) {
				return 1, nil
			},
		}

		svc := services.NewUserService(mockRepo)
		isDeleted, err := svc.DeleteUserByID(ctx, 1)

		// Assertions
		assert.Equal(t, isDeleted, 1)
		assert.NoError(t, err)
	})
}
