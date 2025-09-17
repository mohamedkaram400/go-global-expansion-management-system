package ports

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)

type ResearchDocumentRepository interface {
	UploadDocument(ctx context.Context, project *entities.Document) (*entities.Document, error)
	SearchOnDocument(ctx context.Context, document *entities.Document) ([]entities.Document, error)
}
 