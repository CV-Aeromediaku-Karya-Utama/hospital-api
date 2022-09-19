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
		//Use(middleware.CheckRole("ADMIN"))
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
			user.PUT("/update_password/:id", s.UpdateUserPassword())
			user.PUT("/assign_permission/:id", s.AssignUserPermission())
			user.PUT("/assign_role/:id", s.AssignUserRole())
			user.DELETE("/delete/:id", s.DeleteUser())
		}
		permission := v1.Group("/permission")
		{
			permission.Use(middleware.JwtTokenCheck)
			permission.GET("/list", s.ListPermission())
			permission.GET("/detail/:id", s.permissionDetail())
		}
		role := v1.Group("/role")
		{
			role.Use(middleware.JwtTokenCheck)
			role.POST("/create", s.CreateRole())
			role.GET("/list", s.ListRole())
			role.GET("/detail/:id", s.roleDetail())
			role.PUT("/update/:id", s.UpdateRole())
			user.PUT("/assign_permission/:id", s.AssignRolePermission())
			role.DELETE("/delete/:id", s.DeleteRole())
		}
	}

	return router
}
