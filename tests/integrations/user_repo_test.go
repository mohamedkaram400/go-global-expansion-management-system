package integrations

import (
	"context"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/adapters/repositories/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewUserRepo(db)
	ctx := context.Background()

	// Seed test data
	usersData := []*entities.User{
		{Name: "User 1", Email: "user1@gmail.com"},
		{Name: "User 2", Email: "user2@gmail.com"},
		{Name: "User 3", Email: "user3@gmail.com"},
	}
	if err := db.Create(&usersData).Error; err != nil {
		t.Fatalf("failed to seed users: %v", err)
	}

	users, err := repo.GetAllUsers(ctx, 0, 3)
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 3, len(users))
}

func TestGetUserByID(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewUserRepo(db)
	ctx := context.Background()

	// Seed one user
	userData := entities.User{Name: "Test User", Email: "test@gmail.com"}
	if err := db.Create(&userData).Error; err != nil {
		t.Fatalf("failed to seed user: %v", err)
	}

	user, err := repo.FindUserByID(ctx, userData.ID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userData.Name, user.Name)
}

func TestInsertUser(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewUserRepo(db)
	ctx := context.Background()

	user := &entities.User{
		Name:     "Ali",
		Email:    "ali@example.com",
		Password: "secret",
	}

	inserted, err := repo.InsertUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)
	assert.NotZero(t, inserted.ID)

	// verify record exists
	found, err := repo.FindUserByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Ali", found.Name)
}

func TestUpdateUser(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewUserRepo(db)
	ctx := context.Background()

	// 1. Insert a user first to make sure it exists
	user := &entities.User{
		Name:     "Ahmed",
		Email:    "ahmed@test.com",
		Password: "secret",
	}

	inserted, err := repo.InsertUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)

	updates := map[string]interface{}{
		"name":     "Ahmed update",
		"email":    "ahmed@update.com",
		"password": "secret1",
	}

	updated, err := repo.UpdateUserByID(ctx, user.ID, updates)
	assert.NoError(t, err)
	assert.NotNil(t, updated)
	assert.NotZero(t, updated.ID)

	// verify record exists
	found, err := repo.FindUserByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Ahmed update", found.Name)
}

func TestDeleteUser(t *testing.T) {
	db := conn.SetupTestDB()
	repo := repositories.NewUserRepo(db)
	ctx := context.Background()

	// 1. Insert a user first to make sure it exists
	user := &entities.User{
		Name:     "Ali",
		Email:    "ali@test.com",
		Password: "secret",
	}
	inserted, err := repo.InsertUser(ctx, user)
	assert.NoError(t, err)
	assert.NotNil(t, inserted)

	// 2. Delete the inserted user
	rowsAffected, err := repo.DeleteUserByID(ctx, inserted.ID)
	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)

	// 3. Verify it's actually deleted
	found, err := repo.FindUserByID(ctx, inserted.ID)
	assert.Error(t, err) // should return error or not found
	assert.Nil(t, found) // should not find the user
}
