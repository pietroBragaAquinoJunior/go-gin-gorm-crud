package main

import (
	"github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

	// r := controllers.GetProductById()
	// r = controllers.NewProduct(r)
	// r.Run(":8080")
}