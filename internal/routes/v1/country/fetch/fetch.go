package fetch

import (
	userModels "github.com/Sabareesh001/penny_tracker_backend/internal/database/models/user"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SelectModel struct{
	Id uint `gorm:"column:id" json:"value"`
	Name string 	`gorm:"column:name" json:"label"`
}

func (SelectModel) TableName() string{
    return "countries"
}

type CommonGenderModel interface {
	[]SelectModel | []userModels.Country
}

func GetCountry(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {

	router.GET("/",func(ctx *gin.Context) {

		format := ctx.Query("format")


        switch format {
				case "select":{
					var country []SelectModel;
					fetchCountry(&country,ctx,DB,redisClient)
					return
				}
				default:{
					var  country []userModels.Country;
				    fetchCountry(&country,ctx,DB,redisClient)
				}
		}

	})

}

func fetchCountry[T CommonGenderModel](country *T,ctx *gin.Context, DB *gorm.DB, redisClient *redis.Client){
					fetchGender := DB.Find(&country);
					if(fetchGender.Error != nil){
						response.SomethingWentWrong(ctx);
						return;
					}
					ctx.AbortWithStatusJSON(200,gin.H{"message":"Succesfully Fetched Data","data":country});
}