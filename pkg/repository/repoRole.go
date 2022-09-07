package repository

import (
	"database/sql"
	"fmt"
	"inventory-api/pkg/api/request"
	"log"
	"strconv"
	"strings"
)

func (s *storage) CreateRole(r request.NewRoleRequest) error {
	statement := `INSERT INTO role (name) VALUES ($1);`

	err := s.db.QueryRow(statement, r.Name).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) ListRole(page int, perPage int) (request.Roles, error) {
	offset := (page - 1) * perPage

	statement := `SELECT id, "name", count(*) OVER() AS total_count FROM role ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(statement, perPage, offset)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.Roles{}, err
	}
	defer rows.Close()

	var roles []request.Role
	var total int
	for rows.Next() {
		var role request.Role
		if err := rows.Scan(&role.ID, &role.Name, &total); err != nil {
			return request.Roles{}, err
		}
		roles = append(roles, role)
	}

	pagination := request.PaginationRequest{
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}

	res := request.Roles{
		Role:       roles,
		Pagination: pagination,
	}

	return res, nil
}

func (s *storage) GetRoleById(roleID int) (request.Role, error) {
	var role request.Role

	statement := `SELECT * FROM role WHERE id = $1`

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
	statement := `UPDATE role SET name = $1 WHERE id = $2`

	err := s.db.QueryRow(statement, role.Name, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.UpdateRoleRequest{}, err
	}

	return role, err
}

func (s *storage) DeleteRole(RoleID int) error {
	statement := `DELETE FROM role WHERE id = $1`

	err := s.db.QueryRow(statement, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) BatchDeleteRole(request request.BatchDeleteRoleRequest) error {
	statement := `DELETE FROM role WHERE id = ANY($1::int[])`

	var ids []string
	for _, s := range request.ID {
		ids = append(ids, strconv.Itoa(s))
	}

	param := "{" + strings.Join(ids, ",") + "}"

	err := s.db.QueryRow(statement, param).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}
