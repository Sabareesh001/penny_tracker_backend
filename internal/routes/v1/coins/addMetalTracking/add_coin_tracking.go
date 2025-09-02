package add_coin_tracking

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/coins"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddCoinTracking(router *gin.RouterGroup, DB *gorm.DB) {
	router.POST("/tracking/:coinId/:status", func(ctx *gin.Context) {

		StatusCode := http.StatusOK

		UserID, contains := userId.GetUserId(ctx)
		if !contains {
			response.UnauthorizedAccess(ctx)
			return
		}

		CoinID, contains := ctx.Params.Get("coinId")
		Status, contains := ctx.Params.Get("status")

		if !contains {
			response.DataInAdequate(ctx)
			return
		}

		/*Update status if record exits else create new one*/

		record := coins.UserCoinTracking{}

		fmt.Println(CoinID)

		DB.Where("user =? AND coin=?", UserID, CoinID).Find(&record)

		if record.Id == 0 {
			numCoinId, err := strconv.Atoi(CoinID)
			numUserId, err := strconv.Atoi(UserID)
			if err != nil {
				response.SomethingWentWrong(ctx)
				return
			}
			record.Coin = numCoinId
			record.User = numUserId
			StatusCode = http.StatusCreated
		}

		if Status == "add" {
			record.Status = coins.StatusEnabled
		} else if Status == "remove" {
			record.Status = coins.StatusDisabled
		} else {
			response.DataInAdequate(ctx)
			return
		}

		insertRecord := DB.Save(&record)

		if insertRecord.Error != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		var Message string

		if Status == "add" {
			Message = "Succesfully Added Tracking"
		} else {
			Message = "Succesfully Removed Tracking"
		}

		ctx.AbortWithStatusJSON(StatusCode, gin.H{"message": Message})

	})
}
