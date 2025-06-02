package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	user_routes "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1"
	"github.com/gin-gonic/gin"
	database "github.com/Sabareesh001/penny_tracker_backend/internal/database"
)

func main(){
	err := godotenv.Load("./database.env","./main.env")
	if(err!=nil){
		fmt.Println("Error in Loading env file")
		fmt.Println(err)
		return
	}
	DB := database.Connect()
    router := gin.Default()
	router.Use();
	apiGroup := router.Group("/api")
	user_routes.UserRoutes(apiGroup.Group("/v1"),DB)
	PORT := os.Getenv("PORT")
    router.Run(":"+PORT)
}