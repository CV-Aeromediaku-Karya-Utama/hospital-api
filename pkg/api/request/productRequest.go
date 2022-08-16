package request

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewProductRequest struct {
	Name string `json:"name"`
}

type UpdateProductRequest struct {
	Name string `json:"name"`
}
