package user_routes

import (
	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/hashing/bcrypt"
)

func UserRoutes(router *gin.RouterGroup,DB *gorm.DB){
    userRoutes := router.Group("/user")
	UserRegistration(userRoutes,DB);
}

func UserRegistration(router *gin.RouterGroup,DB *gorm.DB){

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

