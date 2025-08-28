package pkg

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
)

type ClientDTO struct {
	ID           uint      `json:"id"`
	CompanyName  string    `json:"company_name"`
	ContactEmail string    `jso:"contect_email"`
}


func ConvertToClientDTO(client *entities.Client) *ClientDTO {
	return &ClientDTO{
		ID:        		 client.ID,
		CompanyName:     client.CompanyName,
		ContactEmail: 	 client.ContactEmail,
	}
}

func ConvertClientsToDTOs(clients []*entities.Client) []*ClientDTO {
	dtos := make([]*ClientDTO, len(clients))
	for i, client := range clients {
		dtos[i] = ConvertToClientDTO(client)
	}
	return dtos
}

func CheckPassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
