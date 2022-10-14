package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipbproject/IPB-Vote/src/database"
	"github.com/ipbproject/IPB-Vote/src/models"
)

func Insert(c *gin.Context) {
	var candidato models.Candidato
	if err := c.ShouldBindJSON(&candidato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaNomeDoCandidatos(&candidato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Create(&candidato)
	c.JSON(http.StatusOK, candidato)
}

func Get(c *gin.Context) {
	var candidatos []models.Candidato
	database.DB.Find(&candidatos)
	c.JSON(http.StatusOK, candidatos)
}

func Update(c *gin.Context) {
	var candidato models.Candidato
	if err := c.ShouldBindJSON(&candidato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidaNomeDoCandidatos(&candidato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	database.DB.Save(&candidato)
	c.JSON(http.StatusOK, candidato)
}

func Delete(c *gin.Context) {
	var candidato models.Candidato
	if err := c.ShouldBindJSON(&candidato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Delete(&candidato)
	c.JSON(http.StatusOK, candidato)
}
