package coins

import (
	// authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	add_coin_tracking "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/coins/addCoinTracking"
	alter_coin_resource "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/coins/alterCoinResource"
	fetch_coins "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/coins/fetch"
	get_price "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/coins/getPrice"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CoinRoutes(router *gin.RouterGroup, DB *gorm.DB) {
		coinRoutes := router.Group("coin",authmiddleware.AuthorizeJWT);
		fetch_coins.FetchAllCoins(coinRoutes,DB);
		add_coin_tracking.AddCoinTracking(coinRoutes,DB);
		alter_coin_resource.AlterCoinResource(coinRoutes,DB);
		get_price.GetPrice(coinRoutes,DB);
}