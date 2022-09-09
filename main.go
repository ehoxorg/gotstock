package main

import (
	"net/http"

	"github.com/edihoxhalli/gotstock/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, domain.GetAll())
	})
	r.GET("/products/:code", func(c *gin.Context) {
		code := c.Param("code")
		c.JSON(http.StatusOK, domain.GetProduct(code))
	})
	r.POST("/products", func(c *gin.Context) {
		c.JSON(http.StatusCreated, domain.AddProduct(&domain.Product{}))
	})
	r.PUT("/products/:code", func(c *gin.Context) {
		code := c.Param("code")
		c.JSON(http.StatusOK, domain.UpdateProduct(&domain.Product{ProductCode: code}))
	})
	r.DELETE("/products/:code", func(c *gin.Context) {
		code := c.Param("code")
		domain.DeleteProduct(code)
		c.Status(http.StatusNoContent)
	})
	r.Run()
}
