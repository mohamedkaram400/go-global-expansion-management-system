package entities

type Document struct {
	ID        uint    `json:"id" bson:"_id,omitempty"`
	ProjectId uint     `json:"project_id" bson:"project_id"`
	Title     string   `json:"title" bson:"title"`
	Content   string   `json:"content" bson:"content"`
	Tags      []string `json:"tags" bson:"tags"`
}