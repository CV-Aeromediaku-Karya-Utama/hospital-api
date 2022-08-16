package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "weight tracker API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) Routes() *gin.Engine {
	router := s.router

	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		v1.GET("/login", s.Login())
		role := v1.Group("/role")
		{
			//role.Use(middleware.JwtTokenCheck)
			role.GET("/list", s.ListRole())
			role.GET("/detail/:id", s.roleDetail())
			role.POST("/create", s.CreateRole())
			role.PUT("/update/:id", s.UpdateRole())
			role.DELETE("/delete/:id", s.DeleteRole())
		}
		user := v1.Group("/user")
		{
			user.POST("/create", s.CreateUser())
			user.GET("/list", s.ListUser())
			user.GET("/detail/:id", s.UserDetail())
			user.PUT("/update/:id", s.UpdateUser())
			user.DELETE("/delete/:id", s.DeleteUser())
		}
	}

	return router
}
