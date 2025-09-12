package fetch_coins

import (
	"fmt"
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/coins"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchAllCoins(router *gin.RouterGroup, DB *gorm.DB) {
	router.GET("/", func(ctx *gin.Context) {

		UserID,contains := userId.GetUserId(ctx);

		if(!contains){
			response.UnauthorizedAccess(ctx)
			return
		}

		type CoinsData struct {
			coins.Coins
			Status string
		}
		var rows []CoinsData
		fetchAllMetals := DB.Table("coins").
	Select("coins.*, COALESCE(uct.status, '0') as status").
	Joins("LEFT JOIN user_coin_trackings AS uct ON uct.coin = coins.id AND uct.user = "+UserID).
	Scan(&rows)
		if fetchAllMetals.Error != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		fmt.Println(rows)
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "sucessfully fetched metals", "data": rows})
	})
}
