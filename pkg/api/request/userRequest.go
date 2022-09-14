package request

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	Sex       string       `json:"sex"`
	Email     string       `json:"email"`
	Password  string       `json:"password,omitempty"`
	Status    int          `json:"status"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Users struct {
	User       []User            `json:"users"`
	Pagination PaginationRequest `json:"pagination"`
}

type NewUserRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

type UpdateUserRequest struct {
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
}

type UpdateUserPasswordRequest struct {
	UpdatedAt   time.Time `json:"updated_at"`
	OldPassword string    `json:"old_password"`
	Password    string    `json:"password"`
}
