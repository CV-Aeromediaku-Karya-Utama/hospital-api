package request

type Product struct {
	ID                int               `json:"id"`
	Name              string            `json:"name"`
	ProductDesc       string            `json:"product_desc"`
	ProductCategoryID []ProductCategory `json:"product_category"`
}

type NewProductRequest struct {
	Name              string `json:"name"`
	ProductDesc       string `json:"product_desc"`
	ProductCategoryID []int  `json:"product_category_id"`
}

type UpdateProductRequest struct {
	Name              string `json:"name"`
	ProductDesc       string `json:"product_desc"`
	ProductCategoryID []int  `json:"product_category_id"`
}

type UpdateCategoryProductRequest struct {
	ProductCategoryID []int `json:"product_category_id"`
}
