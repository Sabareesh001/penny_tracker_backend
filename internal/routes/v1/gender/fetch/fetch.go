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
    return "genders"
}

type CommonGenderModel interface {
	[]SelectModel | []userModels.Gender
}

func GetGender(router *gin.RouterGroup, DB *gorm.DB, redisClient *redis.Client) {

	router.GET("/",func(ctx *gin.Context) {

		format := ctx.Query("format")


        switch format {
				case "select":{
					var gender []SelectModel;
					fetchGender(&gender,ctx,DB,redisClient)
					return
				}
				default:{
					var gender []userModels.Gender;
				    fetchGender(&gender,ctx,DB,redisClient)
				}
		}

	})

}

func fetchGender[T CommonGenderModel](gender *T,ctx *gin.Context, DB *gorm.DB, redisClient *redis.Client){
					fetchGender := DB.Find(&gender);
					if(fetchGender.Error != nil){
						response.SomethingWentWrong(ctx);
						return;
					}
					ctx.AbortWithStatusJSON(200,gin.H{"message":"Succesfully Fetched Data","data":gender});
}