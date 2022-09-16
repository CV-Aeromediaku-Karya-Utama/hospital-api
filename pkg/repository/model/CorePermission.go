package model

import "hospital-api/pkg/api/helper"

type CorePermission struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type CorePermissions struct {
	Permission []CorePermission         `json:"permissions"`
	Pagination helper.PaginationRequest `json:"pagination"`
}
