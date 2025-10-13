package integrations

import (
	"context"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/stretchr/testify/assert"
)

func TestGetAllClients(t *testing.T) {
	repo := repositories.NewClientRepo(TestDB)
	ctx := context.Background()

	clients, err := repo.GetAllClients(ctx, 0, 10)
	assert.NoError(t, err)
	assert.NotNil(t, clients)
	assert.Equal(t, 5, len(clients))
}

func TestGetClientByID(t *testing.T) {
	repo := repositories.NewClientRepo(TestDB)
	ctx := context.Background()

	client, err := repo.FindClientByID(ctx, 5)
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestInsertClient(t *testing.T) {
	repo := repositories.NewClientRepo(TestDB)
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
	repo := repositories.NewClientRepo(TestDB)
	ctx := context.Background()

	updates := map[string]interface{}{
		"company_name":  "Updated Company",
		"contact_email": "test@update.com",
		"password":      "secret1",
	}

	inserted, err := repo.UpdateClientByID(ctx, "4", updates)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)
	assert.NotZero(t, inserted.ID)

	// verify record exists
	found, err := repo.FindClientByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Company", found.CompanyName)
}

func TestDeleteClient(t *testing.T) {
	repo := repositories.NewClientRepo(TestDB)
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
	rowsAffected, err := repo.DeleteClientByID(ctx, "6")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), rowsAffected) 

	// 3. Verify it's actually deleted
	found, err := repo.FindClientByID(ctx, inserted.ID)
	assert.Error(t, err)                // should return error or not found
	assert.Nil(t, found)                // should not find the client
}