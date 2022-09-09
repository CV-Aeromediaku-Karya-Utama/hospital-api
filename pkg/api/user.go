package api

import (
	"errors"
	"hospital-api/pkg/api/request"
	"strings"
	"time"
)

// UserService contains the methods of the user service
type UserService interface {
	New(user request.NewUserRequest) error
	List() ([]request.User, error)
	Update(UserID int, request request.UpdateUserRequest) error
	Delete(UserID int) error
	Detail(UserID int) (request.SingleUser, error)
}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	HashPassword(password string) (string, error)
	CreateUser(request.NewUserRequest) error
	GetRoleById(RoleID int) (request.Role, error)
	GetUserByID(id int) (request.SingleUser, error)
	ListUser() ([]request.User, error)
	UpdateUser(UserUD int, role request.UpdateUserRequest) error
	DeleteUser(UserID int) error
}

type userService struct {
	storage UserRepository
}

func (u *userService) Detail(UserID int) (request.SingleUser, error) {
	item, err := u.storage.GetUserByID(UserID)
	if err != nil {
		return request.SingleUser{}, errors.New("user id not found")
	}
	return item, nil
}

func (u *userService) List() ([]request.User, error) {
	data, err := u.storage.ListUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userService) Update(UserID int, request request.UpdateUserRequest) error {
	if request.Name == "" {
		return errors.New("user service - name required")
	}
	if request.Username == "" {
		return errors.New("user service - Username required")
	}
	if request.Sex == "" {
		return errors.New("user service - Sex required")
	}
	if request.Email == "" {
		return errors.New("user service - Email required")
	}
	if request.RoleID == 0 {
		return errors.New("user service - RoleID required")
	}
	request.UpdatedAt = time.Now()

	_, err := u.storage.GetUserByID(UserID)
	if err != nil {
		return errors.New("user id not found")
	}

	_, err = u.storage.GetRoleById(request.RoleID)
	if err != nil {
		return errors.New("role id not found")
	}

	err = u.storage.UpdateUser(UserID, request)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (u *userService) Delete(UserID int) error {
	err := u.storage.DeleteUser(UserID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) New(user request.NewUserRequest) error {
	role, err := u.storage.GetRoleById(user.RoleID)
	if err != nil {
		return err
	}

	if user.Email == "" {
		return errors.New("user service - email required")
	}
	if user.Name == "" {
		return errors.New("user service - name required")
	}

	// do some basic normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	hash, err := u.storage.HashPassword(user.Password)
	if err != nil {
		return err
	}

	newUser := request.NewUserRequest{
		Name:     user.Name,
		Username: user.Username,
		Password: hash,
		Sex:      user.Sex,
		Email:    user.Email,
		RoleID:   role.ID,
	}

	err = u.storage.CreateUser(newUser)

	if err != nil {
		return err
	}

	return nil
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}
