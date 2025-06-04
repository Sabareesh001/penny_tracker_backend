package authmiddleware

import (
	// "fmt"
	"fmt"
	"os"
	"strings"

	jwtAuth "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthorizeJWT(ctx *gin.Context){
	SECRET_KEY := os.Getenv("JWT_SECRET")
    fmt.Println(SECRET_KEY)
	tokenStr := ctx.GetHeader("Authorization")
	tokenStr = strings.Replace(tokenStr,"Bearer","",1)
	tokenStr = strings.Trim(tokenStr," ")
    token,err := jwtAuth.ValidateJwt(tokenStr,SECRET_KEY)
	
	if(err!=nil){
		response.UnauthorizedAccess(ctx)
	}
    
	claims := token.Claims.(jwt.MapClaims)

	ctx.Set("userId",claims["userId"].(float64))
}