// routes/routes.go
package routes

import (
	"github.com/Rgp-desa/app-v1/controllers" // Importe o pacote de controladores aqui
	"github.com/gin-gonic/gin"
)

// ConfigurarRotas configura todas as rotas da API
func ConfigurarRotas(router *gin.Engine) {
	// Rota para obter um ticket pelo ID
	router.GET("/tickets/:id", controllers.FindTicketByID)

	// Rota para criar um novo ticket
	router.POST("/tickets", controllers.CreateTicket)

	// // Rota para atualizar um ticket existente
	router.PUT("/tickets/:id", controllers.UpdatedTicketInput)

	// // Rota para excluir um ticket pelo ID
	router.DELETE("/tickets/:id", controllers.DeleteTicket)
}
