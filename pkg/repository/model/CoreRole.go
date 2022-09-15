package model

import (
	"hospital-api/pkg/api/helper"
)

type CoreRole struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Permission []CorePermission `gorm:"many2many:core_roles_permissions"`
}

type NewCoreRole struct {
	Name       string `json:"name"`
	Permission []int  `json:"permission"`
}

type CoreRoles struct {
	Roles      []CoreRole               `json:"roles"`
	Pagination helper.PaginationRequest `json:"pagination"`
}

type BatchDeleteRole struct {
	ID []int `json:"id"`
}

type UpdateRole struct {
	Name string `json:"name"`
}
