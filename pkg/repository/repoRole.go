package repository

import (
	"database/sql"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
	"os"
)

func (s *storage) CreateRole(r request.NewRoleRequest) error {
	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `INSERT INTO role (name) VALUES (?);`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `INSERT INTO role (name) VALUES ($1);`
	}

	err := s.db.QueryRow(statement, r.Name).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) ListRole() ([]request.Role, error) {
	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `SELECT * FROM role`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `SELECT * FROM role`
	}

	rows, err := s.db.Query(statement)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return nil, err
	}
	defer rows.Close()

	// slice to hold data from returned rows.
	var roles []request.Role

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var role request.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (s *storage) GetRoleById(roleID int) (request.Role, error) {
	var role request.Role

	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `SELECT * FROM role WHERE id = ?`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `SELECT * FROM role WHERE id = $1`
	}

	err := s.db.QueryRow(statement, roleID).Scan(&role.ID, &role.Name)

	if err == sql.ErrNoRows {
		return request.Role{}, fmt.Errorf("unknown value : %d", roleID)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.Role{}, err
	}

	return role, nil
}

func (s *storage) UpdateRole(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error) {
	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `UPDATE role SET name = ? WHERE id = ?`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `UPDATE role SET name = $1 WHERE id = $2`
	}

	err := s.db.QueryRow(statement, role.Name, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.UpdateRoleRequest{}, err
	}

	return role, err
}

func (s *storage) DeleteRole(RoleID int) error {
	statement := ``

	if os.Getenv("DB_DRIVER") == "mysql" {
		statement = `DELETE FROM role WHERE id = ?`
	}
	if os.Getenv("DB_DRIVER") == "postgres" {
		statement = `DELETE FROM role WHERE id = $1`
	}

	err := s.db.QueryRow(statement, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
