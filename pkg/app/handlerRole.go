package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"inventory-api/pkg/api/helper"
	"inventory-api/pkg/api/request"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) CreateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newData request.NewRoleRequest

		err := c.ShouldBindJSON(&newData)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.ErrorResponse(err))
			return
		}

		err = s.roleService.New(newData)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.ErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse("New Role has been created"))
	}
}

func (s *Server) roleDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("ID not found"))
			return
		}

		data, err := s.roleService.Detail(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func (s *Server) ListRole() gin.HandlerFunc {
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

		roles, err := s.roleService.List(page, perPage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, roles)
	}
}

func (s *Server) UpdateRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var Data request.UpdateRoleRequest
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("can't bind the value"))
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("ID not found"))
			return
		}

		response, err := s.roleService.Update(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) DeleteRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, errors.New("ID not found"))
			return
		}

		err = s.roleService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, errors.New("failed to update"))
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
