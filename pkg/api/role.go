package api

import (
	"context"
	"errors"
	"hospital-api/pkg/api/request"
	"strings"
)

// RoleService contains the methods of the user service
type RoleService interface {
	New(r request.NewRoleRequest) error
	List(page int, perPage int) (request.Roles, error)
	Update(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error)
	Delete(RoleID int) error
	BatchDelete(request request.BatchDeleteRoleRequest) error
	Detail(RoleID int) (request.Role, error)
}

// RoleRepository is what lets our service do db operations without knowing anything about the implementation
type RoleRepository interface {
	CreateRole(ctx context.Context, r request.NewRoleRequest) error
	GetRoleById(RoleID int) (request.Role, error)
	ListRole(ctx context.Context, page int, perPage int) (request.Roles, error)
	UpdateRole(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error)
	DeleteRole(RoleID int) error
	BatchDeleteRole(request request.BatchDeleteRoleRequest) error
}

type roleService struct {
	storage RoleRepository
}

func (s *roleService) Update(RoleID int, r request.UpdateRoleRequest) (request.UpdateRoleRequest, error) {
	if r.Name == "" {
		return r, errors.New("role service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	_, err := s.storage.GetRoleById(RoleID)
	if err != nil {
		return request.UpdateRoleRequest{}, err
	}

	role, err := s.storage.UpdateRole(RoleID, r)
	if err != nil {
		return request.UpdateRoleRequest{}, err
	}

	return role, nil
}

func (s *roleService) Delete(RoleID int) error {
	err := s.storage.DeleteRole(RoleID)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleService) BatchDelete(request request.BatchDeleteRoleRequest) error {
	err := s.storage.BatchDeleteRole(request)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleService) Detail(RoleID int) (request.Role, error) {
	item, err := s.storage.GetRoleById(RoleID)
	if err != nil {
		return request.Role{}, errors.New("role id not found")
	}
	return item, nil
}

func (s *roleService) List(page int, perPage int) (request.Roles, error) {
	ctx := context.TODO()
	roles, err := s.storage.ListRole(ctx, page, perPage)
	if err != nil {
		return request.Roles{}, err
	}
	return roles, nil
}

func (s *roleService) New(r request.NewRoleRequest) error {
	if r.Name == "" {
		return errors.New("role service - name required")
	}
	if r.Permission == nil {
		return errors.New("role service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	ctx := context.Background()
	err := s.storage.CreateRole(ctx, r)

	if err != nil {
		return err
	}

	return nil
}

func NewRoleService(roleRepo RoleRepository) RoleService {
	return &roleService{
		storage: roleRepo,
	}
}
