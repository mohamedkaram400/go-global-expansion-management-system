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

func TestFeatchAllClients(t *testing.T) {
	ctx := context.Background()

	t.Run("error: while fetching clients", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			GetAllClientsFunc: func(ctx context.Context, skip int, limit int) ([]entities.Client, error) {
				return nil, errors.New("database error") 
			},
		}

		svc := services.NewClientService(mockRepo)
		clients, err := svc.GetAllClients(ctx, 0, 10)

		// Assertions
		assert.Nil(t, clients)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})


	t.Run("successfully featched", func (t *testing.T)  {
		mockRepo := &mocks.MockClientRepo{
			GetAllClientsFunc: func (ctx context.Context, skip int, limit int) ([]entities.Client, error)  {
				return []entities.Client{
					{ID: 1, CompanyName: "Test Co", ContactEmail: "a@test.com"},
					{ID: 1, CompanyName: "Another Co", ContactEmail: "b@test.com"},
				}, nil
			},
		}

		svc := services.NewClientService(mockRepo)
		clients, err := svc.GetAllClients(ctx, 0, 10)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, clients)
		assert.Equal(t, 2, len(clients)) 
		assert.Equal(t, "Test Co", clients[0].CompanyName)
	})
}

func TestFindClientByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: client id not found", func(t *testing.T) {
			mockRepo := &mocks.MockClientRepo{
				FindClientByIDFunc: func(ctx context.Context, clientID uint) (*entities.Client, error) {
					return nil, errors.New("database error") 
				},
			}

			svc := services.NewClientService(mockRepo)

			client, err := svc.FindClientByID(ctx, 1)

			// Assertions
			assert.Nil(t, client)
			assert.Error(t, err)		
			assert.EqualError(t, err, "database error")
	})

	t.Run("successfully returned", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			FindClientByIDFunc: func(ctx context.Context, clientID uint) (*entities.Client, error) {
				return &entities.Client{ID: clientID, CompanyName: "Test Co"}, nil
			},
		}

		svc := services.NewClientService(mockRepo)

		client, err := svc.FindClientByID(ctx, 1)

		// Assertions
		assert.Nil(t, err)
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, "Test Co", client.CompanyName)
	})
}

func TestInsertClient(t *testing.T) {
	ctx := context.Background()

	t.Run("error: email already exists", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.Client, error) {
				return &entities.Client{ID: 1, ContactEmail: email}, nil
			},
		}

		svc := services.NewClientService(mockRepo)
		req := &requests.ClientRequest{CompanyName: "Test", ContactEmail: "x@test.com", Password: "qazwsx@123"}

		client, err := svc.InsertClient(ctx, req)

		// Assertions
		assert.Nil(t, client)
		assert.EqualError(t, err, "email already in use")
	})

	t.Run("successfully inserted", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.Client, error) {
				return nil, nil
			},
			InsertClientFunc: func(ctx context.Context, client *entities.Client) (*entities.Client, error) {
				client.ID = 42
				return client, nil
			},
		}

		svc := services.NewClientService(mockRepo)
		req := &requests.ClientRequest{CompanyName: "NewCo", ContactEmail: "new@test.com", Password: "qazwsx@123"}

		client, err := svc.InsertClient(ctx, req)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, uint(42), client.ID)
	})
}
 
func TestUpdateClientByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: email already in use", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.Client, error) {
				return &entities.Client{ID: 2, ContactEmail: email}, nil // simulate existing user with same email
			},
			UpdateClientByIDFunc: func(ctx context.Context, clientID string, updates map[string]interface{}) (*entities.Client, error) {
				return nil, nil
			},
		}

		svc := services.NewClientService(mockRepo)
		newClient := &entities.Client{ID: 1, CompanyName: "NewCo", ContactEmail: "new@test.com"}
		client, err := svc.UpdateClientByID(ctx, "1", newClient)

		assert.Nil(t, client)
		assert.EqualError(t, err, "email already in use")
	})

	t.Run("error: database update failure", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.Client, error) {
				return nil, nil // no conflict
			},
			UpdateClientByIDFunc: func(ctx context.Context, clientID string, updates map[string]interface{}) (*entities.Client, error) {
				return nil, errors.New("database error")
			},
		}

		svc := services.NewClientService(mockRepo)
		newClient := &entities.Client{ID: 1, CompanyName: "NewCo", ContactEmail: "new@test.com"}
		client, err := svc.UpdateClientByID(ctx, "1", newClient)

		assert.Nil(t, client)
		assert.EqualError(t, err, "database error")
	})

	t.Run("successfully updated", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			GetByEmailFunc: func(ctx context.Context, email string) (*entities.Client, error) {
				return nil, nil // no existing email
			},
			UpdateClientByIDFunc: func(ctx context.Context, clientID string, updates map[string]interface{}) (*entities.Client, error) {
				return &entities.Client{
					ID:            42,
					CompanyName:   updates["company_name"].(string),
					ContactEmail:  updates["contact_email"].(string),
				}, nil
			},
		}

		svc := services.NewClientService(mockRepo)
		newClient := &entities.Client{CompanyName: "NewCo", ContactEmail: "new@test.com"}
		client, err := svc.UpdateClientByID(ctx, "1", newClient)

		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, uint(42), client.ID)
		assert.Equal(t, "NewCo", client.CompanyName)
		assert.Equal(t, "new@test.com", client.ContactEmail)
	})
}
 
func TestDeleteClientByID(t *testing.T) {
	ctx := context.Background()

	t.Run("error: client id not found", func(t *testing.T) {
			mockRepo := &mocks.MockClientRepo{
				DeleteClientByIDFunc: func(ctx context.Context, clientID string) (int, error) {
					return 0, errors.New("database error") 
				},
			}

			svc := services.NewClientService(mockRepo)
			isDeleted, err := svc.DeleteClientByID(ctx, "1")

			// Assertions
			assert.Equal(t, isDeleted, 0)
			assert.Error(t, err)		
			assert.EqualError(t, err, "database error")
	})

	t.Run("successfully deleted", func(t *testing.T) {
		mockRepo := &mocks.MockClientRepo{
			DeleteClientByIDFunc: func(ctx context.Context, clientID string) (int, error) {
				return 1, nil
			},
		}

		svc := services.NewClientService(mockRepo)
		isDeleted, err := svc.DeleteClientByID(ctx, "1")

		// Assertions
		assert.Equal(t, isDeleted, 1)
		assert.NoError(t, err)
	})
}