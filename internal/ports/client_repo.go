package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
)

type ClientRepository interface {
    InsertClient(ctx context.Context, client *entities.Client) (*entities.Client, error)
    FindClientByID(ctx context.Context, clientID string) (*entities.Client, error)
    GetAllClients(ctx context.Context, skip int, limit int) ([]entities.Client, error)
    UpdateClientByID(ctx context.Context, clientID string, updates map[string]interface{}) (*entities.Client, error)
    DeleteClientByID(ctx context.Context, clientID string) (int, error)
}