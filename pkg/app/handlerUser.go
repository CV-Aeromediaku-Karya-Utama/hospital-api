package app

import (
	"github.com/gin-gonic/gin"
	"inventory-api/pkg/api/request"
	"log"
	"net/http"
)

func (s *Server) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newUser request.NewUserRequest

		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = s.userService.New(newUser)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new user created",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) ListUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		data, err := s.userService.List()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, data)
	}
}
