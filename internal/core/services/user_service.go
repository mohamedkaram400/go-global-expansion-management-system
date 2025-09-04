package services

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests"
)

type UserService struct {
	Repo ports.UserRepository 
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (svc *UserService) GetAllUsers(ctx context.Context, skip int, limit int) ([]entities.User, error) {
	return svc.Repo.GetAllUsers(ctx, skip, limit)
}
 
func (svc *UserService) FindUserByID(ctx context.Context, userID string) (*entities.User, error) {
	return svc.Repo.FindUserByID(ctx, userID)

}

func (svc *UserService) InsertUser(ctx context.Context, req *requests.UserRequest) (*entities.User, error) {
	user := &entities.User{
		Name: 		req.Name,
		Email: 		req.Email,
		Role: 		req.Role,
		Password: 	req.Password,
	}

	existing, _ := svc.Repo.GetByEmail(ctx, user.Email)
    if existing != nil && existing.ID != user.ID {
        return nil, errors.New("email already in use")
    }

	return svc.Repo.InsertUser(ctx, user)
}

func (svc *UserService) UpdateUserByID(ctx context.Context, userID string, newUser *entities.User) (*entities.User, error) {
	updates := map[string]interface{}{}

	if newUser.Name != "" {
		updates["name"] = newUser.Name
	}

	if newUser.Email != "" {
		updates["email"] = newUser.Email
	}

	if newUser.Role != "" {
		updates["role"] = newUser.Role
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	existing, _ := svc.Repo.GetByEmail(ctx, newUser.Email)
    if existing != nil && existing.ID != newUser.ID {
        return nil, errors.New("email already in use")
    }

	return svc.Repo.UpdateUserByID(ctx, userID, updates)
}

func (svc *UserService) DeleteUserByID(ctx context.Context, userID string) (int, error) {
	return svc.Repo.DeleteUserByID(ctx, userID)

}