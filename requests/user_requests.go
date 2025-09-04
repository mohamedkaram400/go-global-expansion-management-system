package requests

type UserRequest struct {
    Name  string `json:"name"  binding:"required,min=3"`
    Email string `json:"email" binding:"required,email"`
    Role string  `json:"role" binding:"required,in:Admin"`
	Password     string `json:"password"   	  binding:"password;not null"`
}
