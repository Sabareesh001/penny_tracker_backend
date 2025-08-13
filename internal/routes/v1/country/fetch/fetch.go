package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCountry(router *gin.RouterGroup, DB *gorm.DB) {

	router.GET("/",func(ctx *gin.Context) {

		
		type CountryData struct {
			Name string `json:"name"`
			Currency string `json:"currency"`
			UnicodeFlag string `json:"unicodeFlag"`
			Iso3       string `json:"iso3"`
		}

		type Body struct {
			Err bool `json:"err"`
			Message string `json:"message"`
			Data []CountryData `json:"data"`
		}

		resp,err := http.Get("https://countriesnow.space/api/v0.1/countries/info?returns=currency,unicodeFlag,iso3")
		if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

        defer resp.Body.Close()

		body,err := io.ReadAll(resp.Body)

		parsedBody  := Body{}

		json.Unmarshal(body,&parsedBody)

		if(err!=nil || parsedBody.Err){
			response.SomethingWentWrong(ctx);
			return
		}
		
		ctx.AbortWithStatusJSON(200,gin.H{"message":"successfully fetched data","data":parsedBody.Data })
		

	})

}