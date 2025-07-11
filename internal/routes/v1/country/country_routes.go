package country_routes

import (
	"github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/country/fetch"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CountryRoutes(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	countryRoutes := router.Group("/country")
    fetch.GetCountry(countryRoutes,DB,redisClient);
}