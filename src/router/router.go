package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/product/:id", controllers.GetProductById)
	r.POST("/product/add", controllers.NewProduct) // Exemplo de outro endpoint
	// r.PUT("/product/:id", controllers.UpdateProduct)
	// r.DELETE("/product/:id", controllers.DeleteProduct)

	return r
}
