package user_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/authentication"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/forgot"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/insert"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/registration"
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/verification"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup,DB *gorm.DB,redisClient *redis.Client){
    userRoutes := router.Group("/user")
	registration.UserRegistration(userRoutes,DB,redisClient);
	verification.Verification(userRoutes,DB,redisClient)
	insert.Insert(userRoutes,DB,redisClient)
    authentication.Auth(userRoutes,DB,redisClient)
	forgot.Forgot(userRoutes,DB,redisClient)
}



