package api

import (
	"errors"
	"github.com/google/uuid"
	"hospital-api/pkg/api/request"
	"strings"
	"time"
)

// UserService contains the methods of the user service
type UserService interface {
	New(user request.NewUserRequest) error
	List(page int, perPage int) (request.Users, error)
	Update(UserID uuid.UUID, request request.UpdateUserRequest) error
	UpdatePassword(UserID uuid.UUID, request request.UpdateUserPasswordRequest) error
	Delete(UserID uuid.UUID) error
	Detail(UserID uuid.UUID) (request.User, error)
}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	CreateUser(request.NewUserRequest) error
	GetUserByID(id uuid.UUID) (request.User, error)
	ListUser(page int, perPage int) (request.Users, error)
	UpdateUser(UserUD uuid.UUID, request request.UpdateUserRequest) error
	UpdateUserPassword(UserID uuid.UUID, request request.UpdateUserPasswordRequest) error
	DeleteUser(UserID uuid.UUID) error
}

type userService struct {
	storage UserRepository
}

func (u *userService) Detail(UserID uuid.UUID) (request.User, error) {
	item, err := u.storage.GetUserByID(UserID)
	if err != nil {
		return request.User{}, errors.New("user id not found")
	}
	return item, nil
}

func (u *userService) List(page int, perPage int) (request.Users, error) {
	data, err := u.storage.ListUser(page, perPage)
	if err != nil {
		return request.Users{}, err
	}
	return data, nil
}

func (u *userService) Update(UserID uuid.UUID, request request.UpdateUserRequest) error {
	if request.Name == "" {
		return errors.New("user service - name required")
	}
	if request.Username == "" {
		return errors.New("user service - Username required")
	}

	request.UpdatedAt = time.Now()

	_, err := u.storage.GetUserByID(UserID)
	if err != nil {
		return errors.New("user id not found")
	}

	err = u.storage.UpdateUser(UserID, request)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (u *userService) UpdatePassword(UserID uuid.UUID, r request.UpdateUserPasswordRequest) error {
	if r.Password == "" {
		return errors.New("user service - password required")
	}
	if r.OldPassword == "" {
		return errors.New("user service - old password required")
	}

	r.UpdatedAt = time.Now()

	user, err := u.storage.GetUserByID(UserID)
	if err != nil {
		return errors.New("user service - user id not found")
	}

	if !u.storage.CheckPasswordHash(r.OldPassword, user.Password) {
		return errors.New("user service - Old Password not match")
	}

	hash, err := u.storage.HashPassword(r.Password)
	newUser := request.UpdateUserPasswordRequest{
		Password: hash,
	}

	err = u.storage.UpdateUserPassword(UserID, newUser)
	if err != nil {
		return errors.New("update Password failed")
	}
	return nil
}

func (u *userService) Delete(UserID uuid.UUID) error {
	err := u.storage.DeleteUser(UserID)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) New(user request.NewUserRequest) error {
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
		Status:   1,
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
