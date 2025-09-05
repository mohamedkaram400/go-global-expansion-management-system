package repositories

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetAllUsers(ctx context.Context, skip int, limit int) ([]entities.User, error) {
    var users []entities.User
    if err := r.DB.WithContext(ctx).
        Offset(skip).
        Limit(limit).
        Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *UserRepo) FindUserByID(ctx context.Context, userID string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.WithContext(ctx).
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} 
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) InsertUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.DB.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) UpdateUserByID(ctx context.Context, userID string, updates map[string]interface{}) (*entities.User, error) {

	user := &entities.User{}

    // Update the user
    if err := r.DB.WithContext(ctx).
        Model(user).
        Where("id = ?", userID).
        Updates(updates).Error; err != nil {
        return nil, err
    }

    // Fetch the updated record
    if err := r.DB.WithContext(ctx).
        Where("id = ?", userID).
        First(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

func (r *UserRepo) DeleteUserByID(ctx context.Context, userID string) (int, error) {
	if err := r.DB.WithContext(ctx).
		Where("id = ?", userID).
		Delete(&entities.User{}).Error; err != nil {
		return 0, err
	}

	return 1, nil
}
