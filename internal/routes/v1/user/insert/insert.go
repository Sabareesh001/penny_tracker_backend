package insert

import (
	email "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/user/insert/email"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Insert(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
	insertRoutes := router.Group("/add",authmiddleware.AuthorizeJWT)
	email.InsertEmail(insertRoutes, DB, redisClient)
}
