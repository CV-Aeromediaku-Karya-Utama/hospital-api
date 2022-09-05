package request

type ProductCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewProductCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateProductCategoryRequest struct {
	Name string `json:"name"`
}

type BatchDeleteProductCategoryRequest struct {
	ID []string `json:"id"`
}
