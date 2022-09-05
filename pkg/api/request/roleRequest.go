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

type UpdateRoleRequest struct {
	Name string `json:"name"`
}
