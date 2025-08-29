package repositories

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"gorm.io/gorm"
)

type AuthRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

// Register a new client
func (r *AuthRepo) Register(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	if err := r.DB.WithContext(ctx).Create(client).Error; err != nil {
		return nil, err
	} 
	return client, nil
}

func (r *AuthRepo) GetClientByCompanyName(ctx context.Context, companyName string) (*entities.Client, error) {
	var client entities.Client
	if err := r.DB.WithContext(ctx).Where("company_name = ?", companyName).First(&client).Error; err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *AuthRepo) Logout(clientID string) (string, error) {

	return "User logout success", nil
}
