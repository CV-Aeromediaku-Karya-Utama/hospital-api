package repository

import (
	"database/sql"
	"fmt"
	"hospital-api/pkg/api/request"
	"log"
)

func (s *storage) CreateUser(request request.NewUserRequest) error {
	statement := `INSERT INTO "user" (name, username, password, sex, email, role_id) VALUES ($1, $2, $3, $4, $5, $6);`

	err := s.db.QueryRow(statement, request.Name, request.Username, request.Password, request.Sex, request.Email, request.RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetUserByID(id int) (request.SingleUser, error) {
	var data request.SingleUser

	statement := `SELECT id,username,email,password,name,sex,role_id FROM "user" WHERE id = $1;`

	err := s.db.QueryRow(statement, id).Scan(&data.ID, &data.Username, &data.Email, &data.Password, &data.Name, &data.Sex, &data.RoleID)

	if err == sql.ErrNoRows {
		return request.SingleUser{}, fmt.Errorf("unknown value : %s", id)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.SingleUser{}, err
	}

	return data, nil
}

func (s *storage) GetUserByEmail(email string) (request.SingleUser, error) {
	var data request.SingleUser

	statement := `SELECT id,username,email,password FROM "user" WHERE email = $1;`

	err := s.db.QueryRow(statement, email).Scan(&data.ID, &data.Username, &data.Email, &data.Password)

	if err == sql.ErrNoRows {
		return request.SingleUser{}, fmt.Errorf("unknown value : %s", email)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.SingleUser{}, err
	}

	return data, nil
}

func (s *storage) GetUserByUsername(username string) (request.SingleUser, error) {
	var data request.SingleUser

	statement := `SELECT id,username,email,password FROM "user" WHERE username = $1;`

	err := s.db.QueryRow(statement, username).Scan(&data.ID, &data.Username, &data.Email, &data.Password)

	if err == sql.ErrNoRows {
		return request.SingleUser{}, fmt.Errorf("unknown value : %s", username)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.SingleUser{}, err
	}

	return data, nil
}

func (s *storage) ListUser() ([]request.User, error) {
	statement := `SELECT a.id, a.created_at, a.updated_at, a.name, a.username, a.sex, a.email, b.id, b.name
					 FROM "user" a
					 INNER JOIN role b
					 ON a.role_id = b.id`

	rows, err := s.db.Query(statement)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var data []request.User

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item request.User
		var role request.Role

		if err := rows.Scan(
			&item.ID,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Name,
			&item.Username,
			&item.Sex,
			&item.Email,
			&role.ID,
			&role.Name,
		); err != nil {
			return data, err
		}
		data = append(data, request.User{
			ID:       item.ID,
			Name:     item.Name,
			Username: item.Username,
			Sex:      item.Sex,
			Email:    item.Email,
			Role: request.Role{
				ID:   role.ID,
				Name: role.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return data, nil
}

func (s *storage) UpdateUser(UserID int, r request.UpdateUserRequest) error {
	statement := `UPDATE "user" SET name = $1, username = $2, sex = $3, email = $4, role_id = $5, updated_at = $6 WHERE id = $7`

	err := s.db.QueryRow(statement, r.Name, r.Username, r.Sex, r.Email, r.RoleID, r.UpdatedAt, UserID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) DeleteUser(UserID int) error {
	statement := `DELETE FROM "user" WHERE id = $1`

	err := s.db.QueryRow(statement, UserID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
