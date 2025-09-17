package repositories

import (
	"context"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResearchDocumentRepo struct {
	MongoCollection *mongo.Collection
}

func NewResearchDocumentRepo(mongoCollection *mongo.Collection) *ResearchDocumentRepo {
	return &ResearchDocumentRepo{MongoCollection: mongoCollection}
}

func (r *ResearchDocumentRepo) UploadDocument(ctx context.Context, document *entities.Document) (*entities.Document, error) {
	_, err := r.MongoCollection.InsertOne(context.Background(), document)

	if err != nil {
		return nil, err
	}

	return document, nil
}

func (r *ResearchDocumentRepo) SearchOnDocument(ctx context.Context, filter bson.M) ([]*entities.Document, error) {
	cursor, err := r.MongoCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var documents []*entities.Document
	if err := cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}
