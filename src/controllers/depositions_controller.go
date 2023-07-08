package controllers

import (
	"alura-backend-7/src/database"
	"alura-backend-7/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllDepositions(c *gin.Context) {
	c.JSON(http.StatusOK, database.GetAllDepositions())
}

func GetDepositionById(c *gin.Context) {
	targetId, _ := strconv.Atoi(c.Params.ByName("id"))
	c.JSON(http.StatusOK, database.GetDepositionById(targetId))
}

func CreateDeposition(c *gin.Context) {
	var newDeposition models.Deposition
	err := c.BindJSON(&newDeposition)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deposition := database.CreateNewDeposition(newDeposition)

	c.JSON(http.StatusOK, deposition)
}

func UpdateDeposition(c *gin.Context) {
	//get id by path
	targetId, _ := strconv.Atoi(c.Params.ByName("id"))
	var newDeposition models.Deposition
	err := c.BindJSON(&newDeposition)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = database.UpdateDeposition(targetId, newDeposition)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deposition updated"})
}

func DeleteDeposition(c *gin.Context) {
	//get id by path
	targetId, _ := strconv.Atoi(c.Params.ByName("id"))
	err := database.DeleteDeposition(targetId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deposition deleted"})
}
