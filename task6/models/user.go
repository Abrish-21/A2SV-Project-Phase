package models

type User struct {
	ID       uint   `json:"id"`                         // Unique user identifier
	Email    string `json:"email" binding:"required"`   // Email address used as username
	Password string `json:"password,omitempty"`         // Password (not exposed in JSON response)
	Role     string `json:"role"`                       // User role, e.g., "admin" or "user"
}
