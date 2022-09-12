package repository

import (
	"database/sql"
	"fmt"
	"hospital-api/pkg/api/request"
	"log"
	"strconv"
	"strings"
)

func (s *storage) CreateRole(r request.NewRoleRequest) error {
	statement := `INSERT INTO core_role (name) VALUES ($1);`

	err := s.db.QueryRow(statement, r.Name).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) ListRole(page int, perPage int) (request.Roles, error) {
	offset := (page - 1) * perPage

	statement := `SELECT id, "name", count(*) OVER() AS total_count FROM core_role ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(statement, perPage, offset)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.Roles{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var coreRoles []request.Role
	var total int
	for rows.Next() {
		var coreRole request.Role
		if err := rows.Scan(&coreRole.ID, &coreRole.Name, &total); err != nil {
			return request.Roles{}, err
		}
		coreRoles = append(coreRoles, coreRole)
	}

	pagination := request.PaginationRequest{
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}

	res := request.Roles{
		Role:       coreRoles,
		Pagination: pagination,
	}

	return res, nil
}

func (s *storage) GetRoleById(coreRoleID int) (request.Role, error) {
	var coreRole request.Role

	statement := `SELECT * FROM core_role WHERE id = $1`

	err := s.db.QueryRow(statement, coreRoleID).Scan(&coreRole.ID, &coreRole.Name)

	if err == sql.ErrNoRows {
		return request.Role{}, fmt.Errorf("unknown value : %d", coreRoleID)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.Role{}, err
	}

	return coreRole, nil
}

func (s *storage) UpdateRole(RoleID int, coreRole request.UpdateRoleRequest) (request.UpdateRoleRequest, error) {
	statement := `UPDATE core_role SET name = $1 WHERE id = $2`

	err := s.db.QueryRow(statement, coreRole.Name, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.UpdateRoleRequest{}, err
	}

	return coreRole, err
}

func (s *storage) DeleteRole(RoleID int) error {
	statement := `DELETE FROM core_role WHERE id = $1`

	err := s.db.QueryRow(statement, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) BatchDeleteRole(request request.BatchDeleteRoleRequest) error {
	statement := `DELETE FROM core_role WHERE id = ANY($1::int[])`

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
