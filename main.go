package main

import (
	"net/http"

	"github.com/edihoxhalli/gotstock/db"
	"github.com/edihoxhalli/gotstock/domain"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		p, sts, err := domain.GetAll()
		if err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(sts, gin.H{"status": sts, "error": err.Error()})
		} else {
			c.JSON(http.StatusOK, p)
		}
	})

	r.GET("/products/:code", func(c *gin.Context) {
		code := c.Param("code")
		p, sts, err := domain.GetProduct(code)
		if err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(sts, gin.H{"status": sts, "error": err.Error()})
		} else {
			c.JSON(http.StatusOK, p)
		}

	})

	r.POST("/products", func(c *gin.Context) {
		var requestBody db.Product
		if err := c.BindJSON(&requestBody); err != nil {
			panic(err)
		}
		res, sts, err := domain.AddProduct(&requestBody)
		if err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(sts, gin.H{"status": sts, "message": err.Error()})
		} else {
			c.JSON(sts, res)
		}
	})

	r.PUT("/products/:code", func(c *gin.Context) {
		code := c.Param("code")
		var requestBody db.Product
		requestBody.ProductCode = code
		if err := c.BindJSON(&requestBody); err != nil {
			panic(err)
		}
		res, sts, err := domain.UpdateProduct(&requestBody, code)
		if err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(sts, gin.H{"status": sts, "message": err.Error()})
		} else {
			c.JSON(sts, res)
		}
	})

	r.DELETE("/products/:code", func(c *gin.Context) {
		code := c.Param("code")
		sts, err := domain.DeleteProduct(code)
		if err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(sts, gin.H{"status": sts, "message": err.Error()})
		} else {
			c.Status(http.StatusNoContent)
		}
	})

	r.Run()
}
