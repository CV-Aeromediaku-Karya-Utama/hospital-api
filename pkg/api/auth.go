package api

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"inventory-api/pkg/api/request"
	"time"
)

// AuthService contains the methods of the service
type AuthService interface {
	Login(input request.LoginInput) (string, error)
}

// AuthRepository is what lets our service do db operations without knowing anything about the implementation
type AuthRepository interface {
	CheckPasswordHash(password, hash string) bool
	HashPassword(password string) (string, error)
	ValidToken(t *jwt.Token, id string) bool
	GetUserByEmail(email string) (request.SingleUser, error)
	GetUserByUsername(username string) (request.SingleUser, error)
}

type authService struct {
	storage AuthRepository
}

func (a authService) Login(input request.LoginInput) (string, error) {
	singleUser := request.SingleUser{}
	username, err := a.storage.GetUserByUsername(input.Identity)
	if err != nil {
		return "Error on username", err
	}

	if username.ID != 0 {
		singleUser = request.SingleUser{
			ID:       username.ID,
			Username: username.Username,
			Email:    username.Email,
			Password: username.Password,
		}
	} else {
		return "User not found", err
	}

	if !a.storage.CheckPasswordHash(input.Identity, singleUser.Password) {
		return "Invalid password", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = singleUser.Username
	claims["user_id"] = singleUser.ID
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	t, err := token.SignedString([]byte("secret"))
	fmt.Println(t)
	return t, err
}

func NewAuthService(authRepo AuthRepository) AuthService {
	return &authService{
		storage: authRepo,
	}
}
