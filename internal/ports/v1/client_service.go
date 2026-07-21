package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
)

type ClientService interface {
	GetAllClients(ctx context.Context, skip int, limit int) ([]*entities.Client, error)
	FindClientByID(ctx context.Context, clientID int) (*entities.Client, error)
	InsertClient(ctx context.Context, req *requests.ClientRequest) (*entities.Client, error)
	UpdateClientByID(ctx context.Context, clientID int, newClient *entities.Client) (*entities.Client, error)
	DeleteClientByID(ctx context.Context, clientID int) (int, error)
}
