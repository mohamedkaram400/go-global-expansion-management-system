package responses

import (
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
)


type UserResponse struct {
	ID           uint      `json:"id"` 
	Name  		 string    `json:"name"`
	Email 		 string    `json:"contact_email"`
	Role     	 string    `json:"role"`
}

func FormatUser(user *entities.User) UserResponse {
	return UserResponse{
		ID:                  user.ID,
		Name:                user.Name,
		Email:               user.Email,
		Role:                user.Role,
	}
}

func FormatUsers(users []entities.User) []UserResponse {
	responses := make([]UserResponse, 0, len(users))
	for _, v := range users {
		responses = append(responses, FormatUser(&v))
	}
	return responses
}