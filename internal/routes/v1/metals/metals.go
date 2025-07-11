package metals

import (
	get_price "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/metals/getPrice"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func MetalRoutes(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
		metalRoutes := router.Group("metal",authmiddleware.AuthorizeJWT)
		get_price.GetPrice(metalRoutes,DB,redisClient);
}