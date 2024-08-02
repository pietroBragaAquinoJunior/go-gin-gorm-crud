package controllers


import (
    "github.com/gin-gonic/gin"
    "github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/database"
    "github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/models"
    "net/http"
)

// retorna o id passado por enquanto
// TODO colocar gorm para buscar o produto pelo id
func GetProductById(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    result := database.DB.First(&product, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func NewProduct(c *gin.Context) {
    var product models.Product
    if err := c.BindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    result := database.DB.Create(&product)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    c.JSON(http.StatusCreated, product)
}