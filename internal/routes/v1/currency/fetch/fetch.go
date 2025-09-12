package fetch

import (
	"encoding/json"
	"io"
	"net/http"

	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCurrencies(router *gin.RouterGroup, DB *gorm.DB) {

	router.GET("/",func(ctx *gin.Context) {
		 type CurrencyData struct {
			Name string
			Currency string
			Iso2 string
			Iso3 string
			UnicodeFlag string
		 }
		 type Body struct {
			Error bool
			Message string
			Data []CurrencyData
		 }

         resp,err := http.Get("https://countriesnow.space/api/v0.1/countries/info?returns=currency,unicodeFlag,iso2,iso3")

		 if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		 }

         defer resp.Body.Close()

		 body,err := io.ReadAll(resp.Body)
		 
		 if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		 }

		 data := Body{}

		 err = json.Unmarshal(body,&data)

		 if(err!=nil){
			response.SomethingWentWrong(ctx)
			return
		 }

		 ctx.AbortWithStatusJSON(http.StatusOK,gin.H{"message":"Successfully fetched currencies","data":data.Data})

	})

}