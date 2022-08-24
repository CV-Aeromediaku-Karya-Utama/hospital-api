package request

import "mime/multipart"

type Product struct {
	ID                int               `json:"id"`
	Name              string            `json:"name"`
	ProductDesc       string            `json:"product_desc"`
	ProductPhoto      multipart.File    `json:"product_photo,omitempty"`
	ProductPhotoUrl   string            `json:"product_photo_url"`
	ProductCategoryID []ProductCategory `json:"product_category"`
}

type NewProductRequest struct {
	Name              string         `json:"name"`
	ProductDesc       string         `json:"product_desc"`
	ProductPhoto      multipart.File `json:"product_photo"`
	ProductPhotoUrl   string         `json:"product_photo_url"`
	ProductCategoryID []int          `json:"product_category_id"`
}

type UpdateProductRequest struct {
	Name              string         `json:"name"`
	ProductDesc       string         `json:"product_desc"`
	ProductPhoto      multipart.File `json:"product_photo"`
	ProductPhotoUrl   string         `json:"product_photo_url"`
	ProductCategoryID []int          `json:"product_category_id"`
}
