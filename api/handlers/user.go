package handlers

import (
	"fmt"
	"hackaton/db"
	"hackaton/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddNewUser(c *gin.Context) {
	var new_user models.User

	if err := c.ShouldBindJSON(&new_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&new_user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var buildings []models.Building
	db.DB.Find(&buildings)

	for _, b := range buildings {
		progress := models.Progress{
			UserID:     int(new_user.ID),
			BuildingID: int(b.ID),
		}
		db.DB.Create(&progress)
	}

	c.JSON(http.StatusCreated, new_user)
	fmt.Println("Received new user:", new_user)
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	result := db.DB.Delete(&models.User{}, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var buildings []models.Progress

	result = db.DB.Where("user_id = ?", userID).Find(&buildings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	result = db.DB.Delete(buildings)

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted"})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.DB.Model(&models.User{}).Where("id = ?", userID).Updates(updatedUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}
