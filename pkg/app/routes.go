package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		role := v1.Group("/role")
		{
			role.GET("/list", s.ListRole())
			role.POST("/create", s.CreateRole())
			role.PUT("/update", s.UpdateRole())
			role.DELETE("/delete", s.DeleteRole())
		}
		user := v1.Group("/user")
		{
			user.POST("/create", s.CreateUser())
		}
	}

	return router
}
