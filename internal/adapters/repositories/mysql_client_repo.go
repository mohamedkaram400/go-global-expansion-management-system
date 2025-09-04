package repositories

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"gorm.io/gorm"
)

type ClientRepo struct {
	DB *gorm.DB
}

func NewClientRepo(db *gorm.DB) *ClientRepo {
	return &ClientRepo{DB: db}
}

func (r *ClientRepo) GetAllClients(ctx context.Context, skip int, limit int) ([]entities.Client, error) {
    var clients []entities.Client
    if err := r.DB.WithContext(ctx).
        Offset(skip).
        Limit(limit).
        Find(&clients).Error; err != nil {
        return nil, err
    }
    return clients, nil
}

func (r *ClientRepo) FindClientByID(ctx context.Context, clientID string) (*entities.Client, error) {
	var client entities.Client
	if err := r.DB.WithContext(ctx).
		Where("id = ?", clientID).
		First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepo) GetByEmail(ctx context.Context, ContactEmail string) (*entities.Client, error) {
	var client entities.Client
	if err := r.DB.WithContext(ctx).
		Where("contact_email = ?", ContactEmail).
		First(&client).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &client, nil
}

func (r *ClientRepo) InsertClient(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	if err := r.DB.WithContext(ctx).Create(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}

func (r *ClientRepo) UpdateClientByID(ctx context.Context, clientID string, updates map[string]interface{}) (*entities.Client, error) {

	client := &entities.Client{}

    // Update the client
    if err := r.DB.WithContext(ctx).
        Model(client).
        Where("id = ?", clientID).
        Updates(updates).Error; err != nil {
        return nil, err
    }

    // Fetch the updated record
    if err := r.DB.WithContext(ctx).
        Where("id = ?", clientID).
        First(client).Error; err != nil {
        return nil, err
    }

    return client, nil
}

func (r *ClientRepo) DeleteClientByID(ctx context.Context, clientID string) (int, error) {
	if err := r.DB.WithContext(ctx).
		Where("id = ?", clientID).
		Delete(&entities.Client{}).Error; err != nil {
		return 0, err
	}

	return 1, nil
}
