package repository

import (
	"inventory-api/pkg/api/request"
)

func (s *storage) CreateProduct(request request.NewProductRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) GetProductByID(id int) (request.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ListProduct() ([]request.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) UpdateProduct(ProductID int, request request.UpdateProductRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s *storage) DeleteProduct(ProductID int) error {
	//TODO implement me
	panic("implement me")
}
