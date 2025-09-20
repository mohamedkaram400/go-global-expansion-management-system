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

func (r *ResearchDocumentRepo) SearchOnDocument(ctx context.Context, searchTerm string) ([]*entities.Document, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"content": bson.M{"$regex": searchTerm, "$options": "i"}},
			{"tags": bson.M{"$in": []string{searchTerm}}},
		},
	}

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

func (r *ResearchDocumentRepo) CountResearchDocsByCountry(ctx context.Context) (map[string]int, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$country"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
	}

	cursor, err := r.MongoCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make(map[string]int)

	for cursor.Next(ctx) {
		var doc struct {
			Country string `bson:"_id"`
			Count   int    `bson:"count"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		result[doc.Country] = doc.Count
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
