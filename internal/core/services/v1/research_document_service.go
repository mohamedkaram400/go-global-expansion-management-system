package services

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1"
)

type ResearchDocumentService struct {
	ResearchDocumentRepo ports.ResearchDocumentRepository
}

func NewResearchDocumentService(repo ports.ResearchDocumentRepository) *ResearchDocumentService {
	return &ResearchDocumentService{ResearchDocumentRepo: repo}
}

func (svc *ResearchDocumentService) UploadDocument(ctx context.Context, document *entities.Document) (*entities.Document, error) {
	// Pass document param to db layer to save it
	return svc.ResearchDocumentRepo.UploadDocument(ctx, document)
}

func (svc *ResearchDocumentService) SearchOnDocument(ctx context.Context, searchTerm string) ([]*entities.Document, error) {
	// Pass search param to db layer to search on documents
	return svc.ResearchDocumentRepo.SearchOnDocument(ctx, searchTerm)
}
