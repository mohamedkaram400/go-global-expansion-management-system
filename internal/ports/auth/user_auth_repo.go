package ports

import (
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"context"
)

type UserAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	Logout(userID string) (string, error)
}
 