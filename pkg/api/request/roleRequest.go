package request

type NewRoleRequest struct {
	Name string `json:"name"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Roles struct {
	Pagination PaginationRequest `json:"pagination"`
	Role       []Role            `json:"data"`
}

type BatchDeleteRoleRequest struct {
	ID []int `json:"id"`
}

type UpdateRoleRequest struct {
	Name string `json:"name"`
}
