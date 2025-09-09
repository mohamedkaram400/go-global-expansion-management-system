package auth

type LoginUserResponse struct {
    ID           uint   `json:"id"`
    Email        string `json:"email"`
    AccessToken  string `json:"access_token"`
    RefrashToken string `json:"refrash_token"`
}