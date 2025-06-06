package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*-_"`
	Username string `json:"username" binding:"required,min=3,max=100"`
}
