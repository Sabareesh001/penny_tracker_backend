package registration

import (
	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/hashing/bcrypt"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

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
			Phone : data.Phone,
			Email : data.Email,
		}
		if err!=nil {
			response.DataInAdequate(ctx)
			return
		}
		createResponse := DB.Omit("isEmailVerified","isPhoneVerified").Create(&user)
		if createResponse.Error!=nil{
			ctx.JSON(500,gin.H{"error":createResponse.Error.Error()})
			return
		}
        ctx.JSON(201,gin.H{"message":"Registration Successfull ✅"})
	})
}