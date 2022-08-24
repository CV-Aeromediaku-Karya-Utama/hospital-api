package api

import (
	"encoding/json"
	"errors"
	"inventory-api/pkg/api/helper"
	"inventory-api/pkg/api/request"
)

// ProductService contains the methods of the user service
type ProductService interface {
	New(request request.NewProductRequest) error
	List() ([]request.Product, error)
	ListByCategory(CategoryID int) ([]request.Product, error)
	Update(ProductID int, request request.UpdateProductRequest) error
	UpdateCategory(ProductID int, CategoryID request.UpdateProductRequest) error
	Delete(ProductID int) error
	Detail(ProductID int) (request.Product, error)
}

// ProductRepository is what lets our service do db operations without knowing anything about the implementation
type ProductRepository interface {
	CreateProduct(request request.Product, categories []request.ProductCategory) error
	GetProductByID(ProductID int) (request.Product, error)
	GetProductByCategory(Category string) ([]request.Product, error)
	GetProductCategoryByID(id int) (request.ProductCategory, error)
	ListProduct() ([]request.Product, error)
	UpdateProduct(ProductID int, request request.UpdateProductRequest) error
	UpdateCategoryByProductID(ProductID int, CategoryID []request.ProductCategory) error
	DeleteProduct(ProductID int) error
}

type productService struct {
	storage ProductRepository
}

func (p productService) New(r request.NewProductRequest) error {
	var categories []request.ProductCategory

	uniqueSlice := helper.UniqueInt(r.ProductCategoryID)
	for i := 0; i < len(uniqueSlice); i++ {
		id := uniqueSlice[i]
		item, err := p.storage.GetProductCategoryByID(id)
		if err != nil {
			return err
		}
		categories = append(categories, item)
	}

	product := new(request.Product)
	product.Name = r.Name
	product.ProductDesc = r.ProductDesc

	err := p.storage.CreateProduct(*product, categories)
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

func (p productService) ListByCategory(CategoryID int) ([]request.Product, error) {
	category, err := p.storage.GetProductCategoryByID(CategoryID)
	if err != nil {
		return nil, err
	}

	// Convert struct to string
	b, err := json.Marshal(category)
	c := "[" + string(b) + "]"
	data, err := p.storage.GetProductByCategory(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p productService) Update(ProductID int, request request.UpdateProductRequest) error {
	if request.Name == "" {
		return errors.New("product service - name required")
	}
	if request.ProductDesc == "" {
		return errors.New("product service - product desc required")
	}

	err := p.storage.UpdateProduct(ProductID, request)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (p productService) UpdateCategory(ProductID int, CategoryID request.UpdateProductRequest) error {
	var categories []request.ProductCategory

	uniqueSlice := helper.UniqueInt(CategoryID.ProductCategoryID)
	for i := 0; i < len(uniqueSlice); i++ {
		id := uniqueSlice[i]
		item, err := p.storage.GetProductCategoryByID(id)
		if err != nil {
			return err
		}
		categories = append(categories, item)
	}

	err := p.storage.UpdateCategoryByProductID(ProductID, categories)
	if err != nil {
		return err
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
