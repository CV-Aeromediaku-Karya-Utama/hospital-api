package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api"
	"log"
)

type Server struct {
	router      *gin.Engine
	authService api.AuthService
	roleService api.RoleService
	userService api.UserService
}

func NewServer(
	router *gin.Engine,
	authService api.AuthService,
	roleService api.RoleService,
	userService api.UserService,
) *Server {
	return &Server{
		router:      router,
		authService: authService,
		roleService: roleService,
		userService: userService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	//port := os.Getenv("APP_PORT")
	//if port == "" {
	//	port = "9000"
	//}
	port := "8080"
	server := fmt.Sprintf(":%s", port)
	err := r.Run(server)

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
