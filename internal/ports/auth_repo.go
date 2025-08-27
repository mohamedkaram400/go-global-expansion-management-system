package ports

import (
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"context"
)

type AuthRepository interface {
	Register(ctx context.Context, client *entities.Client) (*entities.Client, error)
	Login(ctx context.Context, email, password string) (*entities.Client, error)
	Logout(clientID string) (string, error)
}
