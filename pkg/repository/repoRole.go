package repository

import (
	"database/sql"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
)

func (s *storage) CreateRole(r request.NewRoleRequest) error {
	statement := `
		INSERT INTO role (name) 
		VALUES (?);
		`

	err := s.db.QueryRow(statement, r.Name).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) ListRole() ([]request.Role, error) {
	rows, err := s.db.Query("SELECT * FROM role")
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

func (s *storage) UpdateRole(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error) {
	statement := `
		UPDATE role SET name = ? WHERE id = ?
		`

	err := s.db.QueryRow(statement, role.Name, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.UpdateRoleRequest{}, err
	}

	return role, err
}

func (s *storage) DeleteRole(RoleID int) error {
	statement := `
		DELETE FROM role WHERE id = ?
		`

	err := s.db.QueryRow(statement, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
