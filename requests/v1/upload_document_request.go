package requests

type UploadDocumentRequest struct {
	ProjectId uint     `json:"project_id" binding:"required"`
	Title     string   `json:"title" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Tags      []string `json:"tags"`
}