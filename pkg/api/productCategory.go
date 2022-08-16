package api

import (
	"errors"
	"inventory-api/pkg/api/request"
	"strings"
)

// ProductCategoryService contains the methods of the user service
type ProductCategoryService interface {
	New(request request.NewProductCategoryRequest) error
	List() ([]request.ProductCategory, error)
	Update(ProductCategoryID int, request request.UpdateProductCategoryRequest) error
	Delete(ProductCategoryID int) error
	Detail(ProductCategoryID int) (request.ProductCategory, error)
}

// ProductCategoryRepository is what lets our service do db operations without knowing anything about the implementation
type ProductCategoryRepository interface {
	CreateProductCategory(request request.NewProductCategoryRequest) error
	GetProductCategoryByID(id int) (request.ProductCategory, error)
	ListProductCategory() ([]request.ProductCategory, error)
	UpdateProductCategory(ProductCategoryID int, request request.UpdateProductCategoryRequest) error
	DeleteProductCategory(ProductCategoryID int) error
}

type productCategoryService struct {
	storage ProductCategoryRepository
}

func (p productCategoryService) New(request request.NewProductCategoryRequest) error {
	if request.Name == "" {
		return errors.New("product category service - name required")
	}
	request.Name = strings.ToLower(request.Name)
	err := p.storage.CreateProductCategory(request)
	if err != nil {
		return err
	}

	return nil
}

func (p productCategoryService) List() ([]request.ProductCategory, error) {
	data, err := p.storage.ListProductCategory()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p productCategoryService) Update(ProductCategoryID int, request request.UpdateProductCategoryRequest) error {
	if request.Name == "" {
		return errors.New("user service - name required")
	}
	_, err := p.storage.GetProductCategoryByID(ProductCategoryID)
	if err != nil {
		return err
	}

	err = p.storage.UpdateProductCategory(ProductCategoryID, request)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (p productCategoryService) Delete(ProductCategoryID int) error {
	err := p.storage.DeleteProductCategory(ProductCategoryID)
	if err != nil {
		return err
	}
	return nil
}

func (p productCategoryService) Detail(ProductCategoryID int) (request.ProductCategory, error) {
	item, err := p.storage.GetProductCategoryByID(ProductCategoryID)
	if err != nil {
		return request.ProductCategory{}, err
	}
	return item, nil
}

func NewProductCategoryService(productCategoryRepo ProductCategoryRepository) ProductCategoryService {
	return &productCategoryService{
		storage: productCategoryRepo,
	}
}
