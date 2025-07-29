package models

type User struct {
	ID       string `json:"id"`
	Email   string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}