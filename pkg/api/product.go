package api

import (
	"inventory-api/pkg/api/request"
)

// ProductService contains the methods of the user service
type ProductService interface {
	New(request request.NewProductRequest) error
	List() ([]request.Product, error)
	Update(ProductID int, request request.UpdateProductRequest) error
	Delete(ProductID int) error
	Detail(ProductID int) (request.Product, error)
}

// ProductRepository is what lets our service do db operations without knowing anything about the implementation
type ProductRepository interface {
	CreateProduct(request request.NewProductRequest) error
	GetProductByID(id int) (request.Product, error)
	ListProduct() ([]request.Product, error)
	UpdateProduct(ProductID int, request request.UpdateProductRequest) error
	DeleteProduct(ProductID int) error
}

type productService struct {
	storage ProductRepository
}

func (p productService) New(request request.NewProductRequest) error {
	err := p.storage.CreateProduct(request)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) List() ([]request.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p productService) Update(ProductID int, request request.UpdateProductRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) Delete(ProductID int) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) Detail(ProductID int) (request.Product, error) {
	//TODO implement me
	panic("implement me")
}

func NewProductService(productRepo ProductRepository) ProductService {
	return &productService{
		storage: productRepo,
	}
}
