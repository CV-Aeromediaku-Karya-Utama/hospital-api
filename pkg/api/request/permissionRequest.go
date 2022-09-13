package request

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Permissions struct {
	Permission []Permission      `json:"permissions"`
	Pagination PaginationRequest `json:"pagination"`
}
