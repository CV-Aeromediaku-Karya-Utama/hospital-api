package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api"
	"log"
	"os"
)

type Server struct {
	router            *gin.Engine
	authService       api.AuthService
	roleService       api.RoleService
	userService       api.UserService
	permissionService api.PermissionService
}

func NewServer(
	router *gin.Engine,
	authService api.AuthService,
	roleService api.RoleService,
	userService api.UserService,
	permissionService api.PermissionService,
) *Server {
	return &Server{
		router:            router,
		authService:       authService,
		roleService:       roleService,
		userService:       userService,
		permissionService: permissionService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9000"
	}
	server := fmt.Sprintf(":%s", port)
	err := r.Run(server)

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
