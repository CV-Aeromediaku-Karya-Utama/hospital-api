package api

import (
	"errors"
	"fmt"
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
	CreateProduct(request request.Product) error
	GetProductByID(id int) (request.Product, error)
	GetProductCategoryByID(id int) (request.ProductCategory, error)
	ListProduct() ([]request.Product, error)
	UpdateProduct(ProductID int, request request.UpdateProductRequest) error
	DeleteProduct(ProductID int) error
}

type productService struct {
	storage ProductRepository
}

func (p productService) New(r request.NewProductRequest) error {
	var categories []request.ProductCategory

	for i := 0; i < len(r.ProductCategoryID); i++ {
		id := r.ProductCategoryID[i]
		item, err := p.storage.GetProductCategoryByID(id)
		fmt.Println("LOOPING", i, item)
		if err != nil {
			return err
		}
		categories = append(categories, item)
	}
	fmt.Println("CATEGORIES", categories)

	product := new(request.Product)
	product.Name = r.Name
	product.ProductDesc = r.ProductDesc
	product.ProductCategoryID = categories

	err := p.storage.CreateProduct(*product)
	if err != nil {
		return err
	}

	return nil
}

func (p productService) List() ([]request.Product, error) {
	data, err := p.storage.ListProduct()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p productService) Update(ProductID int, request request.UpdateProductRequest) error {
	if request.Name == "" {
		return errors.New("user service - name required")
	}
	_, err := p.storage.GetProductByID(ProductID)
	if err != nil {
		return err
	}

	err = p.storage.UpdateProduct(ProductID, request)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (p productService) Delete(ProductID int) error {
	err := p.storage.DeleteProduct(ProductID)
	if err != nil {
		return err
	}
	return nil
}

func (p productService) Detail(ProductID int) (request.Product, error) {
	item, err := p.storage.GetProductByID(ProductID)
	if err != nil {
		return request.Product{}, err
	}
	return item, nil
}

func NewProductService(productRepo ProductRepository) ProductService {
	return &productService{
		storage: productRepo,
	}
}
