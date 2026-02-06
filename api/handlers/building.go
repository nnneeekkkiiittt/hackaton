package handlers

import (
	"hackaton/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var mock_buildings = []models.Building{
	{ID: 1, Name: "Coliseum", Description: "Tremendous arena", Latitude: 41.89002, Longtitude: 12.4925, VideoURL: "http://example.com/video1"},
	{ID: 2, Name: "Eiffel Tower", Description: "The renowned tower in Paris", Latitude: 48.858, Longtitude: 2.2945, VideoURL: "http://example.com/video2"},
}

func GetBuildings(c *gin.Context) {
	c.JSON(http.StatusOK, mock_buildings)
}

func GetBuildingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}
	for _, b := range mock_buildings {
		if b.ID == uint(id) {
			c.JSON(http.StatusOK, b)
			return
		}
	}

	c.JSON(404, gin.H{"message": "Building not found"})
}

func CreateNewBuilding(c *gin.Context) {
	var newBuilding models.Building

	if err := c.ShouldBindJSON(&newBuilding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mock_buildings = append(mock_buildings, newBuilding)

	c.JSON(http.StatusCreated, newBuilding)
}
