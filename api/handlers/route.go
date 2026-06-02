package handlers

import (
	"hackaton/db"
	"hackaton/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRoutes(c *gin.Context) {
	var routes []models.Route
	if err := db.DB.Find(&routes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, routes)
}

func GetRouteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid route id"})
		return
	}

	var route models.Route
	if err := db.DB.First(&route, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
		return
	}
	c.JSON(http.StatusOK, route)
}

func GetRouteBuildings(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid route id"})
		return
	}

	var route models.Route
	if err := db.DB.Preload("Buildings", func(db *gorm.DB) *gorm.DB {
		return db.Order("id asc")
	}).First(&route, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
		return
	}

	c.JSON(http.StatusOK, route.Buildings)
}

func CreateRoute(c *gin.Context) {
	var route models.Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, route)
}

func UpdateRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid route id"})
		return
	}

	var route models.Route
	if err := db.DB.First(&route, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
		return
	}

	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&route).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, route)
}

func DeleteRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid route id"})
		return
	}

	if err := db.DB.Delete(&models.Route{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "route deleted"})
}
