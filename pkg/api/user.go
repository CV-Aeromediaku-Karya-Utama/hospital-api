package api

import (
	"errors"
	"hospital-api/pkg/repository/model"
	"strings"
	"time"
)

// UserService contains the methods of the user service
type UserService interface {
	New(user model.NewCoreUser) error
	List(page int, perPage int) (model.CoreUsers, error)
	Update(UserID int, request model.UpdateCoreUser) error
	UpdatePassword(UserID int, request model.UpdateCoreUserPassword) error
	Delete(UserID int) error
	Detail(UserID int) (model.CoreUser, error)
	AssignPermission(UserID int, request model.CoreUser) error
	AssignRole(UserID int, request model.CoreUser) error
}

// UserRepository is what lets our service do db operations without knowing anything about the implementation
type UserRepository interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	CreateUser(model.NewCoreUser) error
	GetUserByID(id int) (model.CoreUser, error)
	ListUser(page int, perPage int) (model.CoreUsers, error)
	UpdateUser(UserUD int, request model.UpdateCoreUser) error
	UpdateUserPassword(UserID int, request model.UpdateCoreUserPassword) error
	DeleteUser(UserID int) error
	AssignPermission(UserID int, request model.CoreUser) error
	AssignRole(UserID int, request model.CoreUser) error
}

type userService struct {
	storage UserRepository
}

func (u *userService) Detail(UserID int) (model.CoreUser, error) {
	item, err := u.storage.GetUserByID(UserID)
	if err != nil {
		return model.CoreUser{}, errors.New("user id not found")
	}
	return item, nil
}

func (u *userService) List(page int, perPage int) (model.CoreUsers, error) {
	data, err := u.storage.ListUser(page, perPage)
	if err != nil {
		return model.CoreUsers{}, err
	}
	return data, nil
}

func (u *userService) Update(UserID int, request model.UpdateCoreUser) error {
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

func (u *userService) UpdatePassword(UserID int, r model.UpdateCoreUserPassword) error {
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
	newUser := model.UpdateCoreUserPassword{
		Password: hash,
	}

	err = u.storage.UpdateUserPassword(UserID, newUser)
	if err != nil {
		return errors.New("update Password failed")
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

func (u *userService) New(user model.NewCoreUser) error {
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

	newUser := model.NewCoreUser{
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

func (u *userService) AssignPermission(UserID int, request model.CoreUser) error {
	err := u.storage.AssignPermission(UserID, request)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) AssignRole(UserID int, request model.CoreUser) error {
	err := u.storage.AssignRole(UserID, request)
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
