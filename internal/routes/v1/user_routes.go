package user_routes

import (
	"context"
	"fmt"
	"strconv"
	"time"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/hashing/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
    userRoutes := router.Group("/user")
	UserRegistration(userRoutes,DB,redisClient);
	Verification(userRoutes,DB,redisClient)
}

func UserRegistration(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){

	router.POST("/register",func(ctx *gin.Context) {
        type UserRegister struct{
			FirstName string `form:"firstName" binding:"required"`
			LastName string  `form:"lastName" binding:"required"`
			Phone string `form:"phone"`
			Email string `form:"email"`
			Country int `form:"country" binding:"required"`
			Age int `form:"age" binding:"required"`
			Username string `form:"username" binding:"required"`
			Password string `form:"password" binding:"required"`
			Occupation int `form:"occupation" binding:"required"`
            Gender int `form:"gender" binding:"required"`
		}
		var data UserRegister
		err := ctx.ShouldBind(&data)
		user := userModel.User{
			First_Name: data.FirstName,
			Last_Name: data.LastName,
            Username: data.Username,
			Password: bcrypt.BcryptGetHash(data.Password),
			Occupation: data.Occupation,
			Gender: data.Gender,
			Country: data.Country,                
			Age : data.Age,
		}
		if err!=nil {
			ctx.String(400,"Data inadequate")
			return
		}
		createResponse := DB.Create(&user)
		if createResponse.Error!=nil{
			panic(createResponse.Error)
		}
        ctx.String(201,"Registration Successfull ✅")
	})
}

func Verification(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
	verificationRoutes := router.Group("/verify")
    MobileVerification(verificationRoutes,DB,redisClient)
}

func MobileVerification(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
    mobileRoutes := router.Group("/mobile")
	RequestOtp(mobileRoutes,DB,redisClient)
}

func RequestOtp(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
      
		router.GET("/requestOtp",func(ctx *gin.Context) {

			
			type MobileData struct{
				Mobile string `json:"mobile"`
				UserId int `json:"userId"`
			}
			data  := MobileData{}
			err := ctx.ShouldBindBodyWithJSON(&data)
			if err!=nil {
				panic(err)
			}
			
			var existingList []userModel.User;
			DB.Where("phone = ?",data.Mobile).Find(&existingList)
			
			if(len(existingList)>0){
				ctx.String(400,"Mobile Number Already Used")
				return
			}
            

			fmt.Println(data);
				mobileCtx := context.Background()
				redisClient.Set(mobileCtx,data.Mobile+"user:"+strconv.Itoa(data.UserId),900,30*time.Second)
			})
}