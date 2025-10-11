package alter_income

import (
	"encoding/json"
	"io"

	"github.com/Sabareesh001/penny_tracker_backend/pkg/contextKeys/userId"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AlterIncome(router *gin.RouterGroup, DB *gorm.DB) {

	router.PATCH("/", func(ctx *gin.Context) {
		UserId, contains := userId.GetUserId(ctx)

        type Body struct {
			Income float32 `json:"income"`
			Currency string `json:"currency"`
		}

		if(!contains){
			response.SomethingWentWrong(ctx)
			return
		}

		defer ctx.Request.Body.Close()

		bodyContent,err := io.ReadAll(ctx.Request.Body)

		if(err!=nil){
			response.DataInAdequate(ctx)
			return
		}
 
		var body Body

		err = json.Unmarshal(bodyContent,&body)
        
		if(err!=nil){
			response.DataInAdequate(ctx)
			return
		}



		updateRow := DB.Table("users").Where("id=?",UserId).Update("monthly_income",body.Income).Update("monthly_income_currency",body.Currency)

		if(updateRow.Error!=nil){
			response.SomethingWentWrong(ctx)
			return
		}

		response.SuccesfullyUpdated(ctx)

	})

}
