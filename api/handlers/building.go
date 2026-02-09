package handlers

import (
	"hackaton/db"
	"hackaton/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var mock_buildings = []models.Building{
	{ID: 1, Name: "Coliseum", Description: "Tremendous arena", Latitude: 41.89002, Longitude: 12.4925, VideoURL: "http://example.com/video1"},
	{ID: 2, Name: "Eiffel Tower", Description: "The renowned tower in Paris", Latitude: 48.858, Longitude: 2.2945, VideoURL: "http://example.com/video2"},
}

func GetBuildings(c *gin.Context) {
	var buildings []models.Building

	result := db.DB.Find(&buildings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, buildings)
}

func GetBuildingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}

	var building models.Building

	result := db.DB.First(&building, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, building)
}

func CreateNewBuilding(c *gin.Context) {
	var newBuilding models.Building

	if err := c.ShouldBindJSON(&newBuilding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&newBuilding)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBuilding)
}

func UpdateTheBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}

	var building models.Building
	result := db.DB.First(&building, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not found"})
		return
	}

	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&building)

	c.JSON(http.StatusOK, building)

}

func DeleteTheBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}

	var building models.Building

	result := db.DB.First(&building, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	db.DB.Delete(&building)

	c.JSON(http.StatusOK, building)
}
