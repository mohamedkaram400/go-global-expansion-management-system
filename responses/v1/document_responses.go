package responses

import (
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
)


type DocumentResponse struct {
	ProjectId 	   int  	`json:"project_id" `
	Title  		   string 	`json:"title"`
	Content 	   string	`json:"content"`
	Tags		   []string	`json:"tags"`
}


func FormatDocument(document *entities.Document) DocumentResponse {
    return DocumentResponse{
        ProjectId: document.ProjectId,
        Title:     document.Title,
        Content:   document.Content,
        Tags:      document.Tags,
    }
}

func FormatDocuments(documents []*entities.Document) []DocumentResponse {
	responses := make([]DocumentResponse, 0, len(documents))
	for _, v := range documents {
		responses = append(responses, FormatDocument(v))
	}
	return responses
}