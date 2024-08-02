package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/models"
	"net/http"
)

// func GetProductById() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/product/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		c.String(200, id)
// 	})
// 	return r
// }

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, id)
}



// func NewProduct(r *gin.Engine) *gin.Engine {
// 	r.POST("/product/add", func(c *gin.Context) {
// 		var product models.Product
// 		c.BindJSON(&product)
// 		c.JSON(200, product)
// 	})
// 	return r
// }


func NewProduct(c *gin.Context)  {
	var product models.Product
	c.BindJSON(&product)
	c.JSON(http.StatusCreated, product)
}