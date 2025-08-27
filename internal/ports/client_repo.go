package ports

import "github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"

type ClientRepository interface {
	InsertClient(emp *entities.Client) (*entities.Client, error)
    FindClientByID(ClientID string) (*entities.Client, error)
    GetAllClients(skip int, limit int) ([]*entities.Client, int, error)
    UpdateClientByID(clientID string, newClient *entities.Client) (int, error)
    DeleteClientByID(clientID string) (int, error)
}

