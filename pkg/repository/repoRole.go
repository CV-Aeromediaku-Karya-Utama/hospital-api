package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"hospital-api/pkg/api/request"
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
	err = tx.QueryRow("INSERT INTO core_role (name) VALUES ($1) RETURNING id;", r.Name).Scan(&ID)
	if err != nil {
		return fmt.Errorf("repo err - %v", err)
	}
	for _, v := range r.Permission {
		_, err = tx.ExecContext(ctx, "INSERT INTO core_role_permission (role_id, permission_id) VALUES ($1, $2);", ID, v)
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

	tx, _ := s.db.BeginTx(ctx, nil)
	defer tx.Rollback()
	statement := `SELECT role_id, array_agg(distinct permission_id) as groups, count(*) OVER() AS total_count 
				FROM core_role_permission GROUP BY role_id ORDER BY role_id DESC LIMIT $1 OFFSET $2`
	rows, err := tx.QueryContext(ctx, statement, perPage, offset)
	if err != nil {
		log.Printf("this was the error: %v", err)
		return request.Roles{}, err
	}
	defer rows.Close()

	var coreRoles []request.Role
	var totalCount int
	var roleId int
	var permissionIds pq.Int32Array
	var coreRole request.Role
	var corePermission request.Permission
	var corePermissions []request.Permission

	for rows.Next() {
		if err := rows.Scan(&roleId, &permissionIds, &totalCount); err != nil {
			return request.Roles{}, err
		}
		log.Print(roleId, permissionIds)
	}
	rows.NextResultSet()

	for rows.Next() {
		// Get Role Detail
		if err := tx.QueryRowContext(ctx, `SELECT * FROM core_role WHERE id = $1`, roleId).Scan(&coreRole.ID, &coreRole.Name); err != nil {
			return request.Roles{}, err
		}
		log.Println(coreRoles)

		// Get Permission Detail
		for _, v := range permissionIds {
			if err = tx.QueryRowContext(ctx, `SELECT * FROM core_permission WHERE id = $1`, v).Scan(&corePermission.ID, &corePermission.Name); err != nil {
				return request.Roles{}, err
			}
			corePermissions = append(corePermissions, corePermission)
		}

		coreRoles = append(coreRoles, request.Role{
			ID:         coreRole.ID,
			Name:       coreRole.Name,
			Permission: corePermissions,
		})
	}

	res := request.Roles{
		Role: coreRoles,
		Pagination: request.PaginationRequest{
			Page:    page,
			PerPage: perPage,
			Total:   totalCount,
		},
	}

	err = tx.Commit()
	if err != nil {
		return request.Roles{}, errors.New("FAILED TO COMMIT")
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
