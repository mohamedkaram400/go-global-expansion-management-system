package ports

import (
	"context"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type UserAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	Logout(userID string) (string, error)
}
