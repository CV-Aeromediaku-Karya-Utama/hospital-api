package repository

import (
	"database/sql"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
)

func (s *storage) CreateUser(request request.NewUserRequest) error {
	newUserStatement := `
		INSERT INTO user (name, username, password, sex, email, role_id) 
		VALUES (?, ?, ?, ?, ?, ?);
		`

	err := s.db.QueryRow(newUserStatement, request.Name, request.Username, request.Password, request.Sex, request.Email, request.RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetUserByEmail(email string) (request.SingleUser, error) {
	var data request.SingleUser
	err := s.db.QueryRow("SELECT id,username,email,password FROM user WHERE email = ?", email).Scan(&data.ID, &data.Username, &data.Email, &data.Password)

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
	err := s.db.QueryRow("SELECT id,username,email,password FROM user WHERE username = ?", username).Scan(&data.ID, &data.Username, &data.Email, &data.Password)

	if err == sql.ErrNoRows {
		return request.SingleUser{}, fmt.Errorf("unknown value : %s", username)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.SingleUser{}, err
	}

	return data, nil
}
