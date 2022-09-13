package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"hospital-api/pkg/api/request"
	"log"
	"net/http"
	"strconv"
)

type UserQueryParams struct {
	ID string `form:"id"`
}

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
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		perPageStr := c.DefaultQuery("per_page", "10")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		data, err := s.userService.List(page, perPage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := map[string]any{
			"status": "success",
			"data":   data,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) UserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var queryParams UserQueryParams
		if c.BindQuery(&queryParams) == nil {
			log.Println("Query String ", queryParams.ID)
		}
		id, _ := uuid.FromString(queryParams.ID)
		data, err := s.userService.Detail(uuid2.UUID(id))
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := map[string]any{
			"status": "success",
			"data":   data,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data request.UpdateUserRequest

		id := uuid.Must(uuid.FromString(c.Param("id")))
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("can't bind the value"))
			return
		}

		err = s.userService.Update(uuid2.UUID(id), Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "user updated",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) UpdateUserPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data request.UpdateUserPasswordRequest

		id := uuid.Must(uuid.FromString(c.Param("id")))
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("can't bind the value"))
			return
		}

		err = s.userService.UpdatePassword(uuid2.UUID(id), Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "user password updated",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id := uuid.Must(uuid.FromString(c.Param("id")))

		err := s.userService.Delete(uuid2.UUID(id))
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		response := map[string]any{
			"status": "success",
			"data":   "user deleted",
		}

		c.JSON(http.StatusOK, response)
	}
}
