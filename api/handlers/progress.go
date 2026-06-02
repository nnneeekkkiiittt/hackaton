package handlers

import (
	"hackaton/db"
	"hackaton/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MarkTheBuildingVisited(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	var building_id int
	building_id, err = strconv.Atoi(c.Param("building_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "building not found"})
		return
	}

	var progress models.Progress
	result := db.DB.Where("user_id = ? AND building_id = ?", user_id, building_id).First(&progress)

	if result.Error != nil {
		progress = models.Progress{
			UserID:     user_id,
			BuildingID: building_id,
			Flag:       true,
		}
		db.DB.Create(&progress)
	} else {
		progress.Flag = true
		db.DB.Save(&progress)
	}

	c.JSON(http.StatusOK, gin.H{"message": "building marked as visited"})
}

func GetUserProgress(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	var progress []models.Progress
	result := db.DB.Where("user_id = ?", user_id).Find(&progress)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

func ResetUserProgress(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	result := db.DB.Where("user_id = ?", userID).Delete(&models.Progress{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "progress reset successfully",
	})
}
