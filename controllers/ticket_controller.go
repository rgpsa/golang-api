package controllers

import (
	"net/http"
	"time"

	"github.com/Rgp-desa/app-v1/models"
	"github.com/gin-gonic/gin"
)

func FindTicketByID(c *gin.Context) {
	ticket := models.Ticket{}
	if err := models.DB.Db.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"ticket": &ticket})
}

func CreateTicket(c *gin.Context) {
	ticket := models.Ticket{}
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Db.Create(&ticket)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Faield to create ticket"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"ticket": &ticket})
}

func UpdatedTicketInput(c *gin.Context) {
	ticket := models.Ticket{}

	if err := models.DB.Db.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	ticket.UpdatedAt = time.Now()
	updatedInput := models.Ticket{}

	if err := c.ShouldBindJSON(&updatedInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Db.Model(ticket).Updates(&updatedInput)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Faield to update ticket"})
		return
	}

	// Atualize o objeto ticket com os novos valores
	// ticket = updatedInput
	// Recupere o registro atualizado do banco de dados
	models.DB.Db.Where("id = ?", c.Param("id")).First(&ticket)

	c.JSON(http.StatusOK, gin.H{"ticket": ticket})

}

func DeleteTicket(c *gin.Context) {
	ticket := models.Ticket{}

	if err := models.DB.Db.Where("id = ?", c.Param("id")).First(&ticket).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}

	if err := models.DB.Db.Delete(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o ticket"})
		return
	}

	c.JSON(http.StatusNoContent, nil)

}
