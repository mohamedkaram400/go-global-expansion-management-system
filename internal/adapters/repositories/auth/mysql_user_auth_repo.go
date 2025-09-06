package repositories

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"gorm.io/gorm"
)

type UserAuthRepo struct {
	DB *gorm.DB
}

func NewUserAuthRepo(db *gorm.DB) *UserAuthRepo {
	return &UserAuthRepo{DB: db}
}

// Register a new User
func (r *UserAuthRepo) Register(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.DB.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	} 
	return user, nil
}

func (r *UserAuthRepo) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserAuthRepo) Logout(userID string) (string, error) {

	return "User logout success", nil
}
