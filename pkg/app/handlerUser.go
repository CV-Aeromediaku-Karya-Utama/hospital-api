package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api/helper"
	"hospital-api/pkg/repository/model"
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

		var newUser model.NewCoreUser

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
			c.AbortWithStatusJSON(helper.ErrorResponse(err))
			return
		}
		perPageStr := c.DefaultQuery("per_page", "10")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			c.AbortWithStatusJSON(helper.ErrorResponse(err))
			return
		}

		data, err := s.userService.List(page, perPage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(helper.SuccessResponse(data))
	}
}

func (s *Server) UserDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var queryParams UserQueryParams
		if c.BindQuery(&queryParams) == nil {
			log.Println("Query String ", queryParams.ID)
		}
		//id, _ := uuid.FromString(queryParams.ID)
		id, _ := strconv.Atoi(queryParams.ID)
		data, err := s.userService.Detail(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(helper.SuccessResponse(data))
	}
}

func (s *Server) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data model.UpdateCoreUser

		//id := uuid.Must(uuid.FromString(c.Param("id")))
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("can't bind the value"))
			return
		}

		err = s.userService.Update(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		c.JSON(helper.SuccessResponse("user updated"))
	}
}

func (s *Server) UpdateUserPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data model.UpdateCoreUserPassword

		//id := uuid.Must(uuid.FromString(c.Param("id")))
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("can't bind the value"))
			return
		}

		err = s.userService.UpdatePassword(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		c.JSON(helper.SuccessResponse("user password updated"))
	}
}

func (s *Server) AssignUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data model.CoreUser

		//id := uuid.Must(uuid.FromString(c.Param("id")))
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.ShouldBindJSON(&Data)
		if err != nil {
			c.AbortWithStatusJSON(helper.ErrorResponse(err))
			return
		}

		err = s.userService.AssignRole(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(helper.SuccessResponse("user password updated"))
	}
}

func (s *Server) AssignUserPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		var Data model.CoreUser

		//id := uuid.Must(uuid.FromString(c.Param("id")))
		id, _ := strconv.Atoi(c.Param("id"))

		err := c.ShouldBindJSON(&Data)
		if err != nil {
			c.AbortWithStatusJSON(helper.ErrorResponse(err))
			return
		}

		err = s.userService.AssignPermission(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(helper.SuccessResponse("user password updated"))
	}
}

func (s *Server) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		//id := uuid.Must(uuid.FromString(c.Param("id")))
		id, _ := strconv.Atoi(c.Param("id"))

		err := s.userService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		c.JSON(helper.SuccessResponse("user deleted"))
	}
}
