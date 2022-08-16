package request

import (
	"database/sql"
	"time"
)

type NewUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
}

type User struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	Sex       string       `json:"sex"`
	Email     string       `json:"email"`
	Role      Role         `json:"role_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type UpdateUserRequest struct {
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	RoleID    int       `json:"role_id"`
}

type SingleUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}
