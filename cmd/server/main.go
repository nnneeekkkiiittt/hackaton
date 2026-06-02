package main

import (
	"hackaton/api/handlers"
	"hackaton/db"
	"hackaton/internal/models"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	db.ConnectToDatabase()

	err := db.DB.AutoMigrate(&models.Route{}, &models.Building{}, &models.User{}, &models.Progress{})

	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	r.GET("/buildings", handlers.GetBuildings)
	r.GET("/routes", handlers.GetRoutes)
	r.GET("/route/:id", handlers.GetRouteByID)
	r.GET("/route/:id/buildings", handlers.GetRouteBuildings)

	r.GET("/building/:id/description", handlers.GetBuildingDescription)
	r.GET("/building/:id", handlers.GetBuildingByID)

	r.POST("/route", handlers.CreateRoute)
	r.POST("/building", handlers.CreateNewBuilding)

	r.PUT("/route/:id", handlers.UpdateRoute)
	r.PUT("/building/:id", handlers.UpdateTheBuilding)

	r.DELETE("/route/:id", handlers.DeleteRoute)
	r.DELETE("/building/:id", handlers.DeleteTheBuilding)

	r.PUT("/progress/:user_id/:building_id", handlers.MarkTheBuildingVisited)

	r.POST("/user", handlers.AddNewUser)

	r.GET("/users", handlers.GetAllUsers)

	r.PUT("/user/:id", handlers.UpdateUser)

	r.DELETE("/user/:id", handlers.DeleteUser)

	r.GET("/progress/:user_id", handlers.GetUserProgress)

	r.DELETE("/progress/:user_id", handlers.ResetUserProgress)

	//сервер
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
