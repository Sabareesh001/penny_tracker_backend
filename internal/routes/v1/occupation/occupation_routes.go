package occupation_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/occupation/fetch"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func OccupationRoutes(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	occupationRoutes := router.Group("/occupation")
    fetch.GetOccupation(occupationRoutes,DB,redisClient);
}