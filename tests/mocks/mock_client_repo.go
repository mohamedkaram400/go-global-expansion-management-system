package mocks


import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

// MockClientRepo simulates ClientRepository behavior
type MockClientRepo struct {
	GetAllClientsFunc func(ctx context.Context, skip, limit int) ([]entities.Client, error)
	GetByEmailFunc    func(ctx context.Context, email string) (*entities.Client, error)
	InsertClientFunc  func(ctx context.Context, client *entities.Client) (*entities.Client, error)
	FindClientByIDFunc func(ctx context.Context, id uint) (*entities.Client, error)
	UpdateClientByIDFunc func(ctx context.Context, id string, updates map[string]interface{}) (*entities.Client, error)
	DeleteClientByIDFunc func(ctx context.Context, id string) (int, error)

}

func (m *MockClientRepo) GetByEmail(ctx context.Context, email string) (*entities.Client, error) {
	return m.GetByEmailFunc(ctx, email)
}

func (m *MockClientRepo) InsertClient(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	return m.InsertClientFunc(ctx, client)
}

func (m *MockClientRepo) GetAllClients(ctx context.Context, skip, limit int) ([]entities.Client, error) { 
	return m.GetAllClientsFunc(ctx, skip, limit)	
}

func (m *MockClientRepo) FindClientByID(ctx context.Context, id uint) (*entities.Client, error) {
	return m.FindClientByIDFunc(ctx, id)	
}

func (m *MockClientRepo) UpdateClientByID(ctx context.Context, id string, updates map[string]interface{}) (*entities.Client, error) {
	return m.UpdateClientByIDFunc(ctx, id, updates)	
}

func (m *MockClientRepo) DeleteClientByID(ctx context.Context, id string) (int, error) {
	return m.DeleteClientByIDFunc(ctx, id)	
}



