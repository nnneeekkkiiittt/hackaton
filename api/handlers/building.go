package handlers

import (
	"hackaton/db"
	"hackaton/internal/models"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

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


func pkgDir() string {
	// Resolve relative to this source file so it works regardless of working directory
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)                     // .../api/handlers
	dir = filepath.Join(dir, "..", "..", "pkg")   // .../pkg (project root / pkg)
	if abs, err := filepath.Abs(dir); err == nil {
		dir = abs
	}
	if info, err := os.Stat(dir); err == nil && info.IsDir() {
		return dir
	}
	// Fallback: from current working directory
	wd, _ := os.Getwd()
	for _, rel := range []string{"pkg", filepath.Join("..", "pkg"), filepath.Join("..", "..", "pkg")} {
		try := filepath.Join(wd, rel)
		if info, err := os.Stat(try); err == nil && info.IsDir() {
			abs, _ := filepath.Abs(try)
			return abs
		}
	}
	return filepath.Join(wd, "pkg")
}


func GetBuildingDescription(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "invalid ID"})
		return
	}
	var building models.Building
	if db.DB.First(&building, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	dir := pkgDir()
	// Derive .txt filename from picture filename (e.g. al_dv_bank.jpg -> al_dv_bank.txt)
	var txtName string
	if building.Picture != "" {
		txtName = strings.TrimSuffix(building.Picture, filepath.Ext(building.Picture)) + ".txt"
	}
	if txtName != "" {
		path := filepath.Join(dir, filepath.Base(txtName))
		if data, err := os.ReadFile(path); err == nil {
			c.Data(http.StatusOK, "text/plain; charset=utf-8", data)
			return
		}
	}
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(building.Description))
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
