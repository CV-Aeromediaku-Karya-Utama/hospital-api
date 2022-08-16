package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inventory-api/pkg/api"
	"log"
	"os"
)

type Server struct {
	router                 *gin.Engine
	authService            api.AuthService
	roleService            api.RoleService
	userService            api.UserService
	productCategoryService api.ProductCategoryService
	productService         api.ProductService
}

func NewServer(
	router *gin.Engine,
	authService api.AuthService,
	roleService api.RoleService,
	userService api.UserService,
	productCategoryService api.ProductCategoryService,
	productService api.ProductService,
) *Server {
	return &Server{
		router:                 router,
		authService:            authService,
		roleService:            roleService,
		userService:            userService,
		productCategoryService: productCategoryService,
		productService:         productService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	err := r.Run(port)

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
