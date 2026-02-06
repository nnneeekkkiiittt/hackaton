package main

import (
	"hackaton/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//маршруты
	r.GET("/buildings", handlers.GetBuildings)

	r.GET("/building/:id", handlers.GetBuildingByID)

	//сервер
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
