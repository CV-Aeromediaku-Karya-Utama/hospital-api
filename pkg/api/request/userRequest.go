package request

import "time"

type NewUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
}

type User struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	RoleID    string    `json:"role_id"`
}
