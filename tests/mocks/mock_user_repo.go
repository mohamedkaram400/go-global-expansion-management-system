package mocks


import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

// MockUserRepo simulates UserRepository behavior
type MockUserRepo struct {
	GetAllUsersFunc func(ctx context.Context, skip, limit int) ([]entities.User, error)
	GetByEmailFunc    func(ctx context.Context, email string) (*entities.User, error)
	InsertUserFunc  func(ctx context.Context, user *entities.User) (*entities.User, error)
	FindUserByIDFunc func(ctx context.Context, id int) (*entities.User, error)
	UpdateUserByIDFunc func(ctx context.Context, id int, updates map[string]interface{}) (*entities.User, error)
	DeleteUserByIDFunc func(ctx context.Context, id int) (int, error)

}

func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	return m.GetByEmailFunc(ctx, email)
}

func (m *MockUserRepo) InsertUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	return m.InsertUserFunc(ctx, user)
}

func (m *MockUserRepo) GetAllUsers(ctx context.Context, skip, limit int) ([]entities.User, error) { 
	return m.GetAllUsersFunc(ctx, skip, limit)	
}

func (m *MockUserRepo) FindUserByID(ctx context.Context, id int) (*entities.User, error) {
	return m.FindUserByIDFunc(ctx, id)	
}

func (m *MockUserRepo) UpdateUserByID(ctx context.Context, id int, updates map[string]interface{}) (*entities.User, error) {
	return m.UpdateUserByIDFunc(ctx, id, updates)	
}

func (m *MockUserRepo) DeleteUserByID(ctx context.Context, id int) (int, error) {
	return m.DeleteUserByIDFunc(ctx, id)	
}



