package api

import (
	"errors"
	"fmt"
	"inventory-api/pkg/api/request"
	"strings"
)

// UserService contains the methods of the user service
type UserService interface {
	New(user request.NewUserRequest) error
}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	CreateUser(request.NewUserRequest) error
	GetRole(RoleID int) (request.Role, error)
}

type userService struct {
	storage UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}

func (u *userService) New(user request.NewUserRequest) error {
	fmt.Println(user)
	role, err := u.storage.GetRole(user.RoleID)
	if err != nil {
		return err
	}

	// do some basic validations
	if user.Email == "" {
		return errors.New("user service - email required")
	}
	if user.Name == "" {
		return errors.New("user service - name required")
	}

	// do some basic normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	newUser := request.NewUserRequest{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
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
