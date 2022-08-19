package repository

import (
	"inventory-api/pkg/api/request"
	"log"
)

func (s *storage) CreateProduct(request request.NewProductRequest, categoryID []byte) error {
	statement := `INSERT INTO inv_product (name, product_desc, product_category_id) VALUES ($1, $2, $3);`

	err := s.db.QueryRow(statement, request.Name, request.ProductDesc, categoryID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
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
