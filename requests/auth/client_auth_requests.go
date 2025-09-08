package auth

type ClientRegisterRequest struct {
    CompanyName  string `json:"company_name" binding:"required,min=3"`
    ContactEmail string `json:"contact_email" binding:"required,email"`
    Password     string `json:"password" binding:"required,min=8"`
}

type ClientLoginRequest struct {
    CompanyName  string `json:"company_name" binding:"required,min=3"`
    Password     string `json:"password" binding:"required,min=8"`
}