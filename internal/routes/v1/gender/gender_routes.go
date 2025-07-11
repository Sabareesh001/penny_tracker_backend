package gender_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/gender/fetch"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GenderRoutes(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	genderRoutes := router.Group("/gender")
    fetch.GetGender(genderRoutes,DB,redisClient);
}