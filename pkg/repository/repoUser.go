package repository

import (
	"database/sql"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
	"os"
)

func (s *storage) CreateUser(request request.NewUserRequest) error {
	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `INSERT INTO user (name, username, password, sex, email, role_id) VALUES (?, ?, ?, ?, ?, ?);`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `INSERT INTO "user" (name, username, password, sex, email, role_id) VALUES ($1, $2, $3, $4, $5, $6);`
	}

	err := s.db.QueryRow(statement, request.Name, request.Username, request.Password, request.Sex, request.Email, request.RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) GetUserByEmail(email string) (request.SingleUser, error) {
	var data request.SingleUser

	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `SELECT id,username,email,password FROM user WHERE email = ?;`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `SELECT id,username,email,password FROM "user" WHERE email = $1;`
	}

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

	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `SELECT id,username,email,password FROM user WHERE username = ?;`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `SELECT id,username,email,password FROM "user" WHERE username = $1;`
	}

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
