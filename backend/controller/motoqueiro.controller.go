package controller

import (
	"github.com/gin-gonic/gin"
	"gestmoto/model"
	"gestmoto/service"
	"strconv"
)

func GetMotoqueiros(c *gin.Context) {
	c.JSON(200, gin.H{
		"motoqueiros": service.GetAll(),
	})
}

func GetMotoqueiro(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}

	motoqueiro, found := service.GetByID(id)
	if !found {
		c.JSON(404, gin.H{
			"message": "Motoqueiro not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"motoqueiro": motoqueiro,
	})
}

func CreateMotoqueiro(c *gin.Context) {
	var newMotoqueiro model.Motoqueiro

	if err := c.BindJSON(&newMotoqueiro); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	newMotoqueiro = service.Create(newMotoqueiro)

	c.JSON(201, gin.H{
		"motoqueiro": newMotoqueiro,
	})
}