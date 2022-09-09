package app

import (
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/app/middleware"
)

func (s *Server) Routes() *gin.Engine {
	router := s.router
	router.Use(middleware.CORS())

	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		v1.POST("/login", s.Login())
		v1.POST("/register", s.CreateUser())
		user := v1.Group("/user")
		{
			user.Use(middleware.JwtTokenCheck)
			user.POST("/create", s.CreateUser())
			user.GET("/list", s.ListUser())
			user.GET("/detail", s.UserDetail())
			user.PUT("/update/:id", s.UpdateUser())
			user.DELETE("/delete/:id", s.DeleteUser())
		}
	}

	return router
}
