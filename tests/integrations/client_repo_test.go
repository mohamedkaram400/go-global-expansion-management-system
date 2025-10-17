package integrations

import (
	"context"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/stretchr/testify/assert"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
)

func TestGetAllClients(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewClientRepo(db)
	ctx := context.Background()

	// Seed test data
	clientsData := []entities.Client{
		{CompanyName: "Client 1", ContactEmail: "client1@gmail.com"},
		{CompanyName: "Client 2", ContactEmail: "client2@gmail.com"},
		{CompanyName: "Client 3", ContactEmail: "client3@gmail.com"},
	}
	if err := db.Create(&clientsData).Error; err != nil {
		t.Fatalf("failed to seed clients: %v", err)
	}

	clients, err := repo.GetAllClients(ctx, 0, 3)
	assert.NoError(t, err)
	assert.NotNil(t, clients)
	assert.Equal(t, 3, len(clients))
}

func TestGetClientByID(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewClientRepo(db)
	ctx := context.Background()

	// Seed one client
	clientData := entities.Client{CompanyName: "Test Client", ContactEmail: "test@gmail.com"}
	if err := db.Create(&clientData).Error; err != nil {
		t.Fatalf("failed to seed client: %v", err)
	}

	client, err := repo.FindClientByID(ctx, clientData.ID)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, clientData.CompanyName, client.CompanyName)
}

func TestInsertClient(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewClientRepo(db)
	ctx := context.Background()

	client := &entities.Client{
		CompanyName:  "Test Company",
		ContactEmail: "test@example.com",
		Password:     "secret",
	}

	inserted, err := repo.InsertClient(ctx, client)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)
	assert.NotZero(t, inserted.ID)

	// verify record exists
	found, err := repo.FindClientByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Test Company", found.CompanyName)
}

func TestUpdateClient(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewClientRepo(db)
	ctx := context.Background()

	// 1. Insert a client first to make sure it exists
	client := &entities.Client{
		CompanyName:  "Delete Test Company",
		ContactEmail: "delete@test.com",
		Password:     "secret",
	}
	inserted, err := repo.InsertClient(ctx, client)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)

	updates := map[string]interface{}{
		"company_name":  "Updated Company",
		"contact_email": "test@update.com",
		"password":      "secret1",
	}

	updated, err := repo.UpdateClientByID(ctx, client.ID, updates)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.NotZero(t, updated.ID)

	// verify record exists
	found, err := repo.FindClientByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Company", found.CompanyName)
}

func TestDeleteClient(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewClientRepo(db)
	ctx := context.Background()

	// 1. Insert a client first to make sure it exists
	client := &entities.Client{
		CompanyName:  "Delete Test Company",
		ContactEmail: "delete@test.com",
		Password:     "secret",
	}
	inserted, err := repo.InsertClient(ctx, client)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)

	// 2. Delete the inserted client
	rowsAffected, err := repo.DeleteClientByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)

	// 3. Verify it's actually deleted
	found, err := repo.FindClientByID(ctx, inserted.ID)
	assert.Error(t, err) // should return error or not found
	assert.Nil(t, found) // should not find the client
}
