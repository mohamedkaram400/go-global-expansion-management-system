package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/services/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRebuild_Success(t *testing.T) {
	ctx := context.Background()

	// 1. Get vendors matches this project
	mockRepo := &mocks.MockMatchRepo{
		GetVendorsForProjectFunc: func(ctx context.Context, project *entities.Project) ([]*entities.Vendor, error) {
			return []*entities.Vendor{
				{ID: 1, Name: "VendorA", ServicesOffered: []byte(`["Consulting"]`), Rating: 4.5, ResponseSlaHours: 2},
			}, nil
		},
		UpsertMatchFunc: func(ctx context.Context, match *entities.Match) error {
			return nil
		},
	}

	mockProjectService := &mocks.MockProjectService{
		FindProjectByIDFunc: func(ctx context.Context, id int) (*entities.Project, error) {
			return &entities.Project{
				ID:             1,
				Country:        "KSA",
				ServicesNeeded: []byte(`["Consulting"]`),
				ClientID:       1,
			}, nil
		},
	}

	mockClientService := &mocks.MockClientService{
		FindClientByIDFunc: func(ctx context.Context, id int) (*entities.Client, error) {
			return &entities.Client{ID: 1, ContactEmail: "test@example.com"}, nil
		},
	}

	mockNotifier := &mocks.MockNotifier{
		SendMatchNotificationFunc: func(ctx context.Context, notif ports.MatchNotification) error {
			fmt.Println("Mock email sent to:", notif.To)
			return nil
		},
	}

	svc := services.NewMatchService(mockRepo, mockProjectService, mockNotifier, mockClientService)
	results, err := svc.Rebuild(ctx, 1)

	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "VendorA", results[0].VendorName)
	assert.True(t, results[0].Score > 0)
}
