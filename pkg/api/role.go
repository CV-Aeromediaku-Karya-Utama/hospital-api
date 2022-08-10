package api

import (
	"errors"
	"strings"
	"weight-tracker-api/pkg/api/request"
)

// RoleService contains the methods of the user service
type RoleService interface {
	New(r request.NewRoleRequest) error
	List() ([]request.Role, error)
	Update(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error)
	Delete(RoleID int) error
}

// RoleRepository is what lets our service do db operations without knowing anything about the implementation
type RoleRepository interface {
	CreateRole(request.NewRoleRequest) error
	ListRole() ([]request.Role, error)
	UpdateRole(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error)
	DeleteRole(RoleID int) error
}

type roleService struct {
	storage RoleRepository
}

func (s *roleService) Update(RoleID int, role request.UpdateRoleRequest) (request.UpdateRoleRequest, error) {
	role, err := s.storage.UpdateRole(RoleID, role)
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

func (s *roleService) List() ([]request.Role, error) {
	role, err := s.storage.ListRole()
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *roleService) New(r request.NewRoleRequest) error {
	if r.Name == "" {
		return errors.New("role service - name required")
	}
	r.Name = strings.ToUpper(r.Name)

	err := s.storage.CreateRole(r)

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
