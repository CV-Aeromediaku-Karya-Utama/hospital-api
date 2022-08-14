package app

import (
	"github.com/gin-gonic/gin"
	"inventory-api/pkg/api/request"
	"log"
	"net/http"
)

func (s *Server) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var input request.LoginInput

		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		token, err := s.authService.Login(input)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, token)
	}
}
