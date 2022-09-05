package request

type PaginationRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"per-page"`
	Total   int `json:"total"`
}
