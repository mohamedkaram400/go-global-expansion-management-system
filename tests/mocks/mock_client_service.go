package mocks

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1"
)

type MockClientService struct {
	GetAllClientsFunc    func(ctx context.Context, skip, limit int) ([]*entities.Client, error)
	InsertClientFunc     func(ctx context.Context, client *requests.ClientRequest) (*entities.Client, error)
	FindClientByIDFunc   func(ctx context.Context, id int) (*entities.Client, error)
	UpdateClientByIDFunc func(ctx context.Context, id int, updates *entities.Client) (*entities.Client, error)
	DeleteClientByIDFunc func(ctx context.Context, id int) (int, error)
}

func (m *MockClientService) InsertClient(ctx context.Context, client *requests.ClientRequest) (*entities.Client, error) {
	return m.InsertClientFunc(ctx, client)
}

func (m *MockClientService) GetAllClients(ctx context.Context, skip, limit int) ([]*entities.Client, error) {
	return m.GetAllClientsFunc(ctx, skip, limit)
}

func (m *MockClientService) FindClientByID(ctx context.Context, clientID int) (*entities.Client, error) {
	return m.FindClientByIDFunc(ctx, clientID)
}

func (m *MockClientService) UpdateClientByID(ctx context.Context, clientID int, updates *entities.Client) (*entities.Client, error) {
	return m.UpdateClientByIDFunc(ctx, clientID, updates)
}

func (m *MockClientService) DeleteClientByID(ctx context.Context, clientID int) (int, error) {
	return m.DeleteClientByIDFunc(ctx, clientID)
}
