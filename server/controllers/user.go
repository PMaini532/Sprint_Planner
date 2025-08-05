package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/PMaini532/Sprint_Planner/database"
	"github.com/PMaini532/Sprint_Planner/models"
)

func GetAllUsers(c *gin.Context){
	var users []models.User;
	if err := database.DB.Select("id","name","email","role").Find(&users).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK,users);
}

func GetAdmins(c *gin.Context){
	var admins []models.User;
	if err := database.DB.Select("id","name","email","role").Where("role = ?", "admin").Find(&admins).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve admins"})
		return
	}
	c.JSON(http.StatusOK,admins);
}

func GetDevelopers(c *gin.Context){
	var developers []models.User;
	if err := database.DB.Select("id","name","email","role").Where("role = ?", "developer").Find(&developers).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve admins"})
		return
	}
	c.JSON(http.StatusOK,developers);
}