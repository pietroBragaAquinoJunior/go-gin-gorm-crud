package main

import (
    "github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/database"
    "github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/router"
    "log"
)

func main() {
    database.Connect()
    r := router.SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Erro ao iniciar o servidor: %v", err)
    }
}
