package repository

import (
	"database/sql"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
)

func (s *storage) CreateProductCategory(request request.NewProductCategoryRequest) error {
	statement := `INSERT INTO inv_product_category (name) VALUES ($1);`

	err := s.db.QueryRow(statement, request.Name).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetProductCategoryByID(id int) (request.ProductCategory, error) {
	var item request.ProductCategory

	statement := `SELECT * FROM inv_product_category WHERE id = $1`

	err := s.db.QueryRow(statement, id).Scan(&item.ID, &item.Name)

	if err == sql.ErrNoRows {
		return request.ProductCategory{}, fmt.Errorf("unknown category id : %d", id)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.ProductCategory{}, err
	}

	return item, nil
}

func (s *storage) ListProductCategory() ([]request.ProductCategory, error) {
	statement := `SELECT * FROM inv_product_category`

	rows, err := s.db.Query(statement)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var data []request.ProductCategory

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item request.ProductCategory
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return data, err
		}
		data = append(data, item)
	}

	return data, nil
}

func (s *storage) UpdateProductCategory(ProductCategoryID int, request request.UpdateProductCategoryRequest) error {
	statement := `UPDATE inv_product_category SET name = $1 WHERE id = $2`

	err := s.db.QueryRow(statement, request.Name, ProductCategoryID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return err
}

func (s *storage) DeleteProductCategory(ProductCategoryID int) error {
	statement := `DELETE FROM inv_product_category WHERE id = $1`

	err := s.db.QueryRow(statement, ProductCategoryID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
