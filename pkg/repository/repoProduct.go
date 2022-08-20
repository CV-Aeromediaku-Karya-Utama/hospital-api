package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
)

func (s *storage) CreateProduct(r request.Product, categories []request.ProductCategory) error {

	statement := `INSERT INTO inv_product (name, product_desc, product_category_id) VALUES ($1, $2, $3);`

	uye, _ := json.Marshal(categories)
	err := s.db.QueryRow(statement, r.Name, r.ProductDesc, uye).Err() //pq.Array()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetProductByID(id int) (request.Product, error) {
	var item request.Product
	//var test []int
	statement := `SELECT * FROM inv_product WHERE id = $1`

	err := s.db.QueryRow(statement, id).Scan(&item.ID, &item.Name, &item.ProductDesc, item.ProductCategoryID)

	if err == sql.ErrNoRows {
		return request.Product{}, fmt.Errorf("unknown value : %d", id)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.Product{}, err
	}

	return item, nil
}

func (s *storage) ListProduct() ([]request.Product, error) {
	statement := `SELECT * FROM inv_product`

	rows, err := s.db.Query(statement)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var data []request.Product

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item request.Product
		if err := rows.Scan(&item.ID, &item.Name); err != nil {
			return data, err
		}
		data = append(data, item)
	}

	return data, nil
}

func (s *storage) UpdateProduct(ProductID int, request request.UpdateProductRequest) error {
	statement := `UPDATE inv_product SET name = $1 WHERE id = $2`

	err := s.db.QueryRow(statement, request.Name, ProductID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return err
}

func (s *storage) DeleteProduct(ProductID int) error {
	statement := `DELETE FROM inv_product WHERE id = $1`

	err := s.db.QueryRow(statement, ProductID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
