package main

import (
	"github.com/Rgp-desa/app-v1/models"
	"github.com/Rgp-desa/app-v1/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicialize a configuração do banco de dados
	models.ConnectDb()

	// Crie uma instância do Gin
	router := gin.Default()

	// Configure as rotas
	routes.ConfigurarRotas(router)

	// Inicie o servidor na porta 8080
	router.Run(":3000")
}
