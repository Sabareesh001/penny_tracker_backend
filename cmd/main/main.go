package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	user_routes "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1"
	"github.com/gin-gonic/gin"
	database "github.com/Sabareesh001/penny_tracker_backend/internal/database"
	redis "github.com/Sabareesh001/penny_tracker_backend/internal/redis"
)

/*

API Documentation => https://documenter.getpostman.com/view/32893888/2sB2qi7HKa

*/

func main(){
	err := godotenv.Load("./database.env","./main.env","./email.env")
	if(err!=nil){
		fmt.Println("Error in Loading env file")
		fmt.Println(err)
		return
	}
	DB := database.Connect()
	redisClient := redis.GetRedisClient(); 
    router := gin.Default()
	router.Use();
	apiGroup := router.Group("/api")
	user_routes.UserRoutes(apiGroup.Group("/v1"),DB,redisClient)
	PORT := os.Getenv("PORT")
    router.Run(":"+PORT)
}