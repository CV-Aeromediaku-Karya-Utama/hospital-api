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
		role := v1.Group("/role")
		{
			role.Use(middleware.JwtTokenCheck)
			role.GET("/list", s.ListRole())
			role.GET("/detail/:id", s.roleDetail())
			role.POST("/create", s.CreateRole())
			role.PUT("/update/:id", s.UpdateRole())
			role.DELETE("/delete/:id", s.DeleteRole())
			role.POST("/batch_delete", s.BatchDeleteRole())
		}
		user := v1.Group("/user")
		{
			user.Use(middleware.JwtTokenCheck)
			user.POST("/create", s.CreateUser())
			user.GET("/list", s.ListUser())
			user.GET("/detail/:id", s.UserDetail())
			user.PUT("/update/:id", s.UpdateUser())
			user.DELETE("/delete/:id", s.DeleteUser())
		}
	}

	return router
}
