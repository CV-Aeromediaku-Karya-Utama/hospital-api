package model

import (
	"gorm.io/gorm"
	"hospital-api/pkg/api/helper"
	"time"
)

type CoreUser struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Username   string
	Sex        string
	Email      string
	Password   string
	Status     int
	Permission []CorePermission `gorm:"many2many:core_users_permissions"`
	Role       []CoreRole       `gorm:"many2many:core_users_roles"`
	gorm.Model
}

type CoreUsers struct {
	User       []CoreUser               `json:"users"`
	Pagination helper.PaginationRequest `json:"pagination"`
}

type NewCoreUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
}

type UpdateCoreUser struct {
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Sex       string    `json:"sex"`
	Email     string    `json:"email"`
	Status    int       `json:"status"`
}

type UpdateCoreUserPassword struct {
	UpdatedAt   time.Time `json:"updated_at"`
	OldPassword string    `json:"old_password"`
	Password    string    `json:"password"`
}