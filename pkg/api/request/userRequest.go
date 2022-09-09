package request

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Username  string       `json:"username"`
	Sex       string       `json:"sex"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
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
