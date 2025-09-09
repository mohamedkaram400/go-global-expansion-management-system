package auth

type UserLoginRequest struct {
    Email        string `json:"email" binding:"required,min=3"`
    Password     string `json:"password" binding:"required,min=8"`
} 