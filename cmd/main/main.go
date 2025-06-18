package main

import (
	"fmt"
	"os"

	database "github.com/Sabareesh001/penny_tracker_backend/internal/database"
	redis "github.com/Sabareesh001/penny_tracker_backend/internal/redis"
	gender_routes "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/gender"
	user_routes "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*

API Documentation => https://documenter.getpostman.com/view/32893888/2sB2qi7HKa

*/

func main(){
	err := godotenv.Load("./database.env","./main.env","./email.env","./jwt.env")
	if(err!=nil){
		fmt.Println("Error in Loading env file")
		fmt.Println(err)
		return
	}
	DB := database.Connect()
	redisClient := redis.GetRedisClient(); 
    router := gin.Default()
	router.Use(cors.Default());
	apiGroup := router.Group("/api")
	v1:=apiGroup.Group("/v1")
	user_routes.UserRoutes(v1,DB,redisClient)
	gender_routes.GenderRoutes(v1,DB,redisClient);
	PORT := os.Getenv("PORT")
    router.Run("0.0.0.0:"+PORT)
}