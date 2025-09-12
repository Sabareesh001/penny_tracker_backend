package get_price

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/coins"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPrice(router *gin.RouterGroup, DB *gorm.DB) {

		router.GET("price/:coin",func(ctx *gin.Context) {
			
			UserId,exists := userId.GetUserId(ctx);

			if(!exists){return}

			coinType,contains := ctx.Params.Get("coin")

            Coin := coins.Coins{};

			fetchCoin := DB.Where("symbol=?",coinType).Find(&Coin);

            if(fetchCoin.Error!=nil){
                    response.DataInAdequate(ctx)
					return
			}

            BaseCoinPriceURL := "https://api.gold-api.com"

			symbolsRes,err := http.Get(BaseCoinPriceURL+"/price/"+coinType)

            if(err != nil){
                 response.SomethingWentWrong(ctx);
				 return;
			}

            defer symbolsRes.Body.Close();

            symbolsBody,err := io.ReadAll(symbolsRes.Body);

			if(err!=nil){
				response.SomethingWentWrong(ctx)
				return
			}


			UserCoinMapping := coins.UserCoinTracking{}

			DB.Where("user=? AND coin=?",UserId,Coin.Id).Find(&UserCoinMapping)
			
			ctx.AbortWithStatusJSON(200,gin.H{"price":json.RawMessage(symbolsBody),"holding":UserCoinMapping.Quantity})

			if(!contains){
				response.DataInAdequate(ctx)
				return;
			}
		})
}