package api

import (
	"errors"
	"hospital-api/pkg/api/request"
)

// PermissionService contains the methods of the user service
type PermissionService interface {
	List(page int, perPage int) (request.Permissions, error)
	Detail(PermissionID int) (request.Permission, error)
}

// PermissionRepository is what lets our service do db operations without knowing anything about the implementation
type PermissionRepository interface {
	GetPermissionById(PermissionID int) (request.Permission, error)
	ListPermission(page int, perPage int) (request.Permissions, error)
}

type permissionService struct {
	storage PermissionRepository
}

func (p permissionService) List(page int, perPage int) (request.Permissions, error) {
	data, err := p.storage.ListPermission(page, perPage)
	if err != nil {
		return request.Permissions{}, err
	}
	return data, nil
}

func (p permissionService) Detail(PermissionID int) (request.Permission, error) {
	item, err := p.storage.GetPermissionById(PermissionID)
	if err != nil {
		return request.Permission{}, errors.New("permission id not found")
	}
	return item, nil
}

func NewPermissionService(permissionRepo PermissionRepository) PermissionService {
	return &permissionService{
		storage: permissionRepo,
	}
}
