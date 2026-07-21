package auth

type RegisterClientResponse struct {
	ID           int    `json:"id"`
	CompanyName  string `json:"company_name"`
	ContactEmail string `json:"contact_email"`
}

type LoginClientResponse struct {
	ID           int    `json:"id"`
	CompanyName  string `json:"company_name"`
	ContactEmail string `json:"contact_email"`
	AccessToken  string `json:"access_token"`
	RefrashToken string `json:"refrash_token"`
}
