package payment

import (
	add_payment "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/payment/add"
	fetch_payment "github.com/Sabareesh001/penny_tracker_backend/internal/routes/v1/payment/fetch"
	authmiddleware "github.com/Sabareesh001/penny_tracker_backend/pkg/jwt/auth-middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func PaymentRoutes(router *gin.RouterGroup, DB *gorm.DB) {
     paymentRoutes := router.Group("payment",authmiddleware.AuthorizeJWT);
	 add_payment.AddPayment(paymentRoutes,DB)
	 fetch_payment.FetchPayments(paymentRoutes,DB);
}