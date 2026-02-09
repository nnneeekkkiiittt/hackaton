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

func UpdateTheBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}

	var building *models.Building

	for i, b := range mock_buildings {
		if b.ID == uint(id) {
			building = &mock_buildings[i]
			break
		}
	}

	if building == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "building not found"})
	}

	var updatedBuilding models.Building
	if err := c.ShouldBindJSON(&updatedBuilding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	building.Name = updatedBuilding.Name
	building.Description = updatedBuilding.Description
	building.Latitude = updatedBuilding.Latitude
	building.Longtitude = updatedBuilding.Longtitude
	building.VideoURL = updatedBuilding.VideoURL

	c.JSON(http.StatusOK, building)

}

func DeleteTheBuilding(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}

	for i, b := range mock_buildings {
		if b.ID == uint(id) {
			mock_buildings = append(mock_buildings[:i], mock_buildings[i+1:]...)
			c.JSON(http.StatusOK, b)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "building not found"})
}
