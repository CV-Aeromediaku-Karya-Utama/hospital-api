package request

import "hospital-api/pkg/repository/model"

type NewRoleRequest struct {
	Name       string `json:"name"`
	Permission []int  `json:"permission"`
}

type Role struct {
	ID         int          `json:"id"`
	Name       string       `json:"name,omitempty"`
	Permission []Permission `json:"permission,omitempty"`
}

type Roles struct {
	Roles      []model.CoreRole  `json:"roles"`
	Pagination PaginationRequest `json:"pagination"`
}

type BatchDeleteRoleRequest struct {
	ID []int `json:"id"`
}

type UpdateRoleRequest struct {
	Name string `json:"name"`
}
