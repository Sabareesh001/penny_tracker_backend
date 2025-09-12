package alter_coin_resource

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/coins"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AlterCoinResource(router *gin.RouterGroup, DB *gorm.DB) {

	router.POST("/quantity",func(ctx *gin.Context) {
		type Body struct {
			Coin_id int `json:"metal_id"`
			Quantity float32 `json:"weight"`
		}

		UserId,contains := userId.GetUserId(ctx);

		if(!contains){
            response.SomethingWentWrong(ctx);
			return
		}

		bodyContent := Body{}

		
		body,err := io.ReadAll(ctx.Request.Body)
		
		if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		}
		
        err = json.Unmarshal(body,&bodyContent)
		
		if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

		userIdInInt,err := strconv.Atoi(UserId);
		
		if(err!=nil){
			response.UnauthorizedAccess(ctx);
		}
		
		
		existingEntry := coins.UserCoinTracking{}
		
		existingRecordFetch := DB.Where("user=? AND coin = ?",UserId,bodyContent.Coin_id).Find(&existingEntry)
		
		fmt.Println(bodyContent.Quantity , bodyContent.Coin_id)
		if(existingRecordFetch.Error != nil){
			newEntry := coins.UserCoinTracking{ User: userIdInInt,Coin: bodyContent.Coin_id,Quantity:float64(bodyContent.Quantity),Status: coins.StatusEnabled}
			newEntryQuery := DB.Save(&newEntry)
			if(newEntryQuery.Error == nil){
				 response.SuccesfullyInserted(ctx)
				 return
			}else{
			     response.SomethingWentWrong(ctx)
				 return	
			}
		}

		existingEntry.Quantity = float64(bodyContent.Quantity)

        updateRecord := DB.Save(&existingEntry)

        if(updateRecord.Error!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

		response.SuccesfullyUpdated(ctx)

	})

}