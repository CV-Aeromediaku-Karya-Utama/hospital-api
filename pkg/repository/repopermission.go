package repository

import (
	"database/sql"
	"fmt"
	"hospital-api/pkg/api/request"
	"log"
)

func (s *storage) GetPermissionById(PermissionID int) (request.Permission, error) {
	var corePermission request.Permission

	statement := `SELECT * FROM core_permissions WHERE id = $1`

	err := s.db.QueryRow(statement, PermissionID).Scan(&corePermission.ID, &corePermission.Name)

	if err == sql.ErrNoRows {
		return request.Permission{}, fmt.Errorf("unknown value : %d", PermissionID)
	}

	if err != nil {
		log.Printf("this was the error: %v", err.Error())
		return request.Permission{}, err
	}

	return corePermission, nil
}

func (s *storage) ListPermission(page int, perPage int) (request.Permissions, error) {
	offset := (page - 1) * perPage

	statement := `SELECT id, "name", count(*) OVER() AS total_count FROM core_permissions ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(statement, perPage, offset)

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.Permissions{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var corePermissions []request.Permission
	var total int
	for rows.Next() {
		var corePermission request.Permission
		if err := rows.Scan(&corePermission.ID, &corePermission.Name, &total); err != nil {
			return request.Permissions{}, err
		}
		corePermissions = append(corePermissions, corePermission)
	}

	pagination := request.PaginationRequest{
		Page:    page,
		PerPage: perPage,
		Total:   total,
	}

	res := request.Permissions{
		Permission: corePermissions,
		Pagination: pagination,
	}

	return res, nil
}
