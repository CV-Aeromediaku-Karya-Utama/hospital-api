package repository

import (
	"context"
	"database/sql"
	"fmt"
	"hospital-api/pkg/api/request"
	"hospital-api/pkg/repository/model"
	"log"
	"strconv"
	"strings"
)

func (s *storage) CreateRole(ctx context.Context, r request.NewRoleRequest) error {
	var ID int
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repo err - %v", err)
	}
	defer tx.Rollback()
	err = tx.QueryRow("INSERT INTO core_roles (name) VALUES ($1) RETURNING id;", r.Name).Scan(&ID)
	if err != nil {
		return fmt.Errorf("repo err - %v", err)
	}
	for _, v := range r.Permission {
		_, err = tx.ExecContext(ctx, "INSERT INTO core_roles_permissions (core_role_id, core_permission_id) VALUES ($1, $2);", ID, v)
		if err != nil {
			return fmt.Errorf("repo err - %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repo err - %v", err)
	}

	return nil
}

func (s *storage) ListRole(ctx context.Context, page int, perPage int) (request.Roles, error) {
	offset := (page - 1) * perPage
	var roles []model.CoreRole
	s.gorm.Preload("Permission").Find(&roles).Select("*")
	res := request.Roles{
		Roles: roles,
		Pagination: request.PaginationRequest{
			Page:    page,
			PerPage: perPage,
			Total:   offset,
		},
	}
	return res, nil
}

func (s *storage) GetRoleById(coreRoleID int) (request.Role, error) {
	var coreRole request.Role

	statement := `SELECT * FROM core_roles WHERE id = $1`

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
	statement := `UPDATE core_roles SET name = $1 WHERE id = $2`

	err := s.db.QueryRow(statement, coreRole.Name, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.UpdateRoleRequest{}, err
	}

	return coreRole, err
}

func (s *storage) DeleteRole(RoleID int) error {
	statement := `DELETE FROM core_roles WHERE id = $1`

	err := s.db.QueryRow(statement, RoleID).Err()

	if err != nil {
		log.Printf("this was the error: %v", err)
		return err
	}

	return nil
}

func (s *storage) BatchDeleteRole(request request.BatchDeleteRoleRequest) error {
	statement := `DELETE FROM core_roles WHERE id = ANY($1::int[])`

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
