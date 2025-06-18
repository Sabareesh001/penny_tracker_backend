package gender_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/gender/fetch"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GenderRoutes(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	genderRoutes := router.Group("/gender",authmiddleware.AuthorizeJWT)
    fetch.GetGender(genderRoutes,DB,redisClient);
}