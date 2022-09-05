package app

import (
	"github.com/gin-gonic/gin"
	"inventory-api/pkg/api/request"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) CreateProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var data request.NewProductCategoryRequest

		err := c.ShouldBindJSON(&data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.productCategoryService.New(data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new productCategory created",
		}
		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) productCategoryDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		data, err := s.productCategoryService.Detail(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, data)
	}
}

func (s *Server) ListProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		productCategories, err := s.productCategoryService.List()
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, productCategories)
	}
}

func (s *Server) UpdateProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var Data request.UpdateProductCategoryRequest
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = s.productCategoryService.Update(id, Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "productCategory updated",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) DeleteProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = s.productCategoryService.Delete(id)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}

func (s *Server) BatchDeleteProductCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var Data request.BatchDeleteProductCategoryRequest
		err := c.ShouldBindJSON(&Data)
		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = s.productCategoryService.BatchDelete(Data)
		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, nil)
	}
}
