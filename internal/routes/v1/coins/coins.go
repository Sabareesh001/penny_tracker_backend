package coins

import (
	// authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	fetch_coins "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/coins/fetch"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CoinRoutes(router *gin.RouterGroup, DB *gorm.DB) {
		coinRoutes := router.Group("coin",authmiddleware.AuthorizeJWT);
		fetch_coins.FetchAllCoins(coinRoutes,DB);
}