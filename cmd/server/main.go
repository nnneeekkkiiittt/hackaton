package main

import (
	"hackaton/api/handlers"
	"hackaton/db"
	"hackaton/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectToDatabase()

	err := db.DB.AutoMigrate(&models.Building{}, &models.User{}, &models.Progress{})

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	//маршруты
	r.GET("/buildings", handlers.GetBuildings)

	r.GET("/building/:id", handlers.GetBuildingByID)

	r.POST("/building", handlers.CreateNewBuilding)

	r.PUT("/building/:id", handlers.UpdateTheBuilding)

	r.DELETE("/building/:id", handlers.DeleteTheBuilding)

	r.PUT("/progress/:user_id/:building_id", handlers.MarkTheBuildingVisited)

	r.POST("/user", handlers.AddNewUser)

	r.GET("/users", handlers.GetAllUsers)

	r.PUT("/user/:id", handlers.UpdateUser)

	r.DELETE("/user/:id", handlers.DeleteUser)

	//сервер
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
