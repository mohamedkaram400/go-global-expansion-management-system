package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *entities.User) (*entities.User, error)
    FindUserByID(ctx context.Context, userID string) (*entities.User, error)
    GetAllUsers(ctx context.Context, skip int, limit int) ([]entities.User, error)
    UpdateUserByID(ctx context.Context, userID string, updates map[string]interface{}) (*entities.User, error)
    DeleteUserByID(ctx context.Context, userID string) (int, error)
    GetByEmail(ctx context.Context, ContactEmail string) (*entities.User, error)
}