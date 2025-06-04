package response

import "github.com/gin-gonic/gin"

func SomethingWentWrong(ctx *gin.Context){
	ctx.AbortWithStatusJSON(500,gin.H{"error":"Something Went Wrong 😖"})
}

func DataInAdequate(ctx *gin.Context){
	ctx.AbortWithStatusJSON(400,gin.H{"error":"Data Inadequate 🗑️"})
}

func NoMatchingRecords(ctx *gin.Context){
	ctx.AbortWithStatusJSON(400,gin.H{"error":"No Matching Records 🥹"})
}

func NoSuchUserExist(ctx *gin.Context){
	ctx.AbortWithStatusJSON(400,gin.H{"error":"No Such User Exist 🚫"})
}

func UnauthorizedAccess(ctx *gin.Context){
	ctx.AbortWithStatusJSON(401,gin.H{"error":"Unauthorized Access 🚫"})
}