package request

type NewRoleRequest struct {
	Name string `json:"name"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateRoleRequest struct {
	Name string `json:"name"`
}

type DeleteRoleRequest struct {
	Name string `json:"name"`
}
