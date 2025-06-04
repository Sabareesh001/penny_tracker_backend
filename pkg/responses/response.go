package response

import "github.com/gin-gonic/gin"

func SomethingWentWrong(ctx *gin.Context){
	ctx.JSON(500,gin.H{"error":"Something Went Wrong 😖"})
}

func DataInAdequate(ctx *gin.Context){
	ctx.JSON(400,gin.H{"error":"Data Inadequate 🗑️"})
}

func NoMatchingRecords(ctx *gin.Context){
	ctx.JSON(400,gin.H{"error":"No Matching Records 🥹"})
}
