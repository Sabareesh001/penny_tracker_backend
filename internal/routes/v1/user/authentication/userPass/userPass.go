package userpass

import (
	"os"
	"time"

	userModel "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	jwtAuth "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserPassValidation(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client){
    router.POST("/userpass",func(ctx *gin.Context) {
		type Body struct {
			Username  string `form:"username" binding:"required"`
			Password  string `form:"password" binding:"required"`
		}
        body := Body{}
        parseError := ctx.ShouldBind(&body)
     
        if parseError!=nil {
			response.DataInAdequate(ctx)
		}

		model := userModel.User{}

        getPassword :=  DB.Where("username=?",body.Username).First(&model)

        if(getPassword.Error != nil){
			if(getPassword.Error.Error()=="record not found"){
				response.NoSuchUserExist(ctx)
				return
			}
			response.SomethingWentWrong(ctx)
			return
		}

		passwordError := bcrypt.CompareHashAndPassword([]byte(model.Password),[]byte(body.Password))

        if(passwordError != nil){
			ctx.JSON(400,gin.H{"error":"Invalid Password 🚫"})
			return
		}
        claims := jwt.MapClaims{
				"userId": model.Id,
				"expiry": time.Now().Add(time.Hour*1).Unix(),
			}
        token,err := jwtAuth.AssignJWT(claims,[]byte(os.Getenv("JWT_SECRET")))

        if err!=nil {
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.JSON(202,gin.H{"message":"Successfully validated credentials ✅","token":token})

	})
}

