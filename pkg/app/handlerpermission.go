package app

import (
	"github.com/gin-gonic/gin"
	"hospital-api/pkg/api/helper"
	"log"
	"strconv"
)

func (s *Server) permissionDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(helper.BadResponse("ID not found"))
			return
		}

		data, err := s.permissionService.Detail(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(data))
	}
}

func (s *Server) ListPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		pageStr := c.DefaultQuery("page", "1")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.AbortWithStatusJSON(helper.BadResponse(err))
			return
		}
		perPageStr := c.DefaultQuery("per_page", "10")
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.AbortWithStatusJSON(helper.BadResponse(err))
			return
		}

		roles, err := s.permissionService.List(page, perPage)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(helper.InternalErrorResponse(err))
			return
		}

		c.JSON(helper.SuccessResponse(roles))
	}
}
