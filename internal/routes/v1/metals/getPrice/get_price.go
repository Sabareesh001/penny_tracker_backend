package get_price

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sabareesh001/penny_tracker_backend/internal/database/models/metals"
	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func GetPrice(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {
		router.GET("price/:metal",func(ctx *gin.Context) {
			
			UserId,exists := userId.GetUserId(ctx);

			if(!exists){return}

			metalType,contains := ctx.Params.Get("metal")

            Metal := metals.Metals{};

			fetchMetal := DB.Where("symbol=?",metalType).Find(&Metal);

            if(fetchMetal.Error!=nil){
                    response.DataInAdequate(ctx)
					return
			}

            BaseMetalPriceURL := "https://api.gold-api.com"

			symbolsRes,err := http.Get(BaseMetalPriceURL+"/price/"+metalType)

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


			UserMetalMapping := metals.UserMetalTracking{}

			fetchHolding := DB.Where("user=? AND metal=?",UserId,Metal.Id).Find(&UserMetalMapping)

			if(fetchHolding.Error != nil){
				response.SomethingWentWrong(ctx)
				return
			}

			ctx.AbortWithStatusJSON(200,gin.H{"price":json.RawMessage(symbolsBody),"holding":UserMetalMapping.Weight})

			if(!contains){
				response.DataInAdequate(ctx)
				return;
			}
			fmt.Println(metalType);
		})
}