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

func (s *storage) GetRole(roleID int) (request.Role, error) {
	var role request.Role
	err := s.db.QueryRow("SELECT * FROM role WHERE id = ?", roleID).Scan(&role.ID, &role.Name)

	if err == sql.ErrNoRows {
		return request.Role{}, fmt.Errorf("unknown value : %d", roleID)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.Role{}, err
	}

	return role, nil
}
