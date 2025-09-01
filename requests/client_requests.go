package requests

type ClientRequest struct {
    CompanyName  string `json:"company_name" binding:"required,min=3"`
    ContactEmail string `json:"contact_email" binding:"required,email"`
}
