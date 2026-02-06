package handlers

import (
	"github.com/gin-gonic/gin"
	"hackaton/internal/models"
)

var mock_buildings = []models.Building{
	{ID: 1, Name: "Coliseum", Description: "Tremendous arena", Latitude: 41.89002, Longtitude: 12.4925, VideoURL: "http://example.com/video1"}, 
	{ID: 2, Name: "Eiffel Tower", Description: "The renowned tower in Paris", Latitude: 48.858, Longtitude: 2.2945, VideoURL: "http://example.com/video2"}
}

func GetBuildings(c *gin.Context) {
	c.JSON(200, mock_buildings)
}

func GetBuildingByID(c *gin.Context) {
	id := int(c.Param("id"))
	for _, b := range mock_buildings {
		if b.ID == id {
			c.JSON(200, b)
			return
		}
	}

	c.JSON(404, gin.H{"message": "Building not found"})
}

func CreateNewBuilding(c *gin.Context, ) {

}

