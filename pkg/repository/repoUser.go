package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"hospital-api/pkg/api/request"
	"log"
)

func (s *storage) CreateUser(request request.NewUserRequest) error {
	statement := `INSERT INTO core_user (name, username, password, sex, email, status) VALUES ($1, $2, $3, $4, $5, $6);`

	err := s.db.QueryRow(statement, request.Name, request.Username, request.Password, request.Sex, request.Email, request.Status).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetUserByID(id uuid.UUID) (request.User, error) {
	var data request.User

	statement := `SELECT id,username,email,password,name,sex,status FROM core_user WHERE id = $1;`

	err := s.db.QueryRow(statement, id).Scan(&data.ID, &data.Username, &data.Email, &data.Password, &data.Name, &data.Sex, &data.Status)

	if err == sql.ErrNoRows {
		return request.User{}, fmt.Errorf("unknown value : %d", id)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.User{}, err
	}

	return data, nil
}

func (s *storage) GetUserByEmail(email string) (request.User, error) {
	var data request.User

	statement := `SELECT id,username,email,password,status FROM core_user WHERE email = $1;`

	err := s.db.QueryRow(statement, email).Scan(&data.ID, &data.Username, &data.Email, &data.Password, &data.Status)

	if err == sql.ErrNoRows {
		return request.User{}, fmt.Errorf("unknown value : %s", email)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.User{}, err
	}

	return data, nil
}

func (s *storage) GetUserByUsername(username string) (request.User, error) {
	var data request.User

	statement := `SELECT id,username,email,password,status FROM core_user WHERE username = $1;`

	err := s.db.QueryRow(statement, username).Scan(&data.ID, &data.Username, &data.Email, &data.Password, &data.Status)

	if err == sql.ErrNoRows {
		return request.User{}, fmt.Errorf("unknown value : %s", username)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.User{}, err
	}

	return data, nil
}

func (s *storage) ListUser(page int, perPage int) (request.Users, error) {
	offset := (page - 1) * perPage
	statement := `SELECT id,name,username,sex,email,status,created_at,updated_at,count(*) OVER() AS total_count FROM core_user ORDER BY id DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(statement, perPage, offset)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.Users{}, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var data []request.User
	var total int
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item request.User

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Username,
			&item.Sex,
			&item.Email,
			&item.Status,
			&item.CreatedAt,
			&item.UpdatedAt,
			&total,
		); err != nil {
			return request.Users{}, err
		}
		data = append(data, request.User{
			ID:        item.ID,
			Name:      item.Name,
			Username:  item.Username,
			Sex:       item.Sex,
			Email:     item.Email,
			Status:    item.Status,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	pagination := request.PaginationRequest{
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}
	res := request.Users{
		User:       data,
		Pagination: pagination,
	}
	return res, nil
}

func (s *storage) UpdateUser(UserID uuid.UUID, r request.UpdateUserRequest) error {
	statement := `UPDATE core_user SET name = $1, username = $2, sex = $3, email = $4,  updated_at = $5 WHERE id = $6`

	err := s.db.QueryRow(statement, r.Name, r.Username, r.Sex, r.Email, r.UpdatedAt, UserID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) DeleteUser(UserID uuid.UUID) error {
	statement := `DELETE FROM core_user WHERE id = $1`

	err := s.db.QueryRow(statement, UserID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
