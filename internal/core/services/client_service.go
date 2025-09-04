package services

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports"
	"github.com/mohamedkaram400/go-global-expansion-management-system/pkg"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests"
)

type ClientService struct {
	Repo ports.ClientRepository 
}

func NewClientService(repo ports.ClientRepository) *ClientService {
	return &ClientService{Repo: repo}
}

func (svc *ClientService) GetAllClients(ctx context.Context, skip int, limit int) ([]entities.Client, error) {
	return svc.Repo.GetAllClients(ctx, skip, limit)
}
 
func (svc *ClientService) FindClientByID(ctx context.Context, clientID string) (*entities.Client, error) {
	return svc.Repo.FindClientByID(ctx, clientID)

}

func (svc *ClientService) InsertClient(ctx context.Context, req *requests.ClientRequest) (*entities.Client, error) {
	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	client := &entities.Client{
		CompanyName:  req.CompanyName,
		ContactEmail: req.ContactEmail,
		Password: 	  hashedPassword,
	}

	existing, _ := svc.Repo.GetByEmail(ctx, client.ContactEmail)
    if existing != nil && existing.ID != client.ID {
        return nil, errors.New("email already in use")
    }

	return svc.Repo.InsertClient(ctx, client)
}

func (svc *ClientService) UpdateClientByID(ctx context.Context, clientID string, newClient *entities.Client) (*entities.Client, error) {
	updates := map[string]interface{}{}

	if newClient.ContactEmail != "" {
		updates["contact_email"] = newClient.ContactEmail
	}

	if newClient.CompanyName != "" {
		updates["company_name"] = newClient.CompanyName
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	existing, _ := svc.Repo.GetByEmail(ctx, newClient.ContactEmail)
    if existing != nil && existing.ID != newClient.ID {
        return nil, errors.New("email already in use")
    }

	return svc.Repo.UpdateClientByID(ctx, clientID, updates)
}

func (svc *ClientService) DeleteClientByID(ctx context.Context, clientID string) (int, error) {
	return svc.Repo.DeleteClientByID(ctx, clientID)

}