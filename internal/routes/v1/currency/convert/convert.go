package convert

import (
	"encoding/json"
	"io"
	"net/http"
	response "github.com/Sabareesh001/penny_tracker_backend/pkg/responses"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConvertCurrency(router *gin.RouterGroup, DB *gorm.DB) {
	router.GET("/convert", func(ctx *gin.Context) {
		from, containsFrom := ctx.GetQuery("from")
		to, containsTo := ctx.GetQuery("to")

		if !containsFrom || !containsTo {
			response.SomethingWentWrong(ctx)
			return
		}

		url := "https://www.revolut.com/api/exchange/quote?amount=1&country=GB&fromCurrency=" + from + "&isRecipientAmount=false&toCurrency=" + to


		// Create a new request so we can set the User-Agent header
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		// Set User-Agent to Mozilla
		req.Header.Set("Accept-Language", "en-US,en;q=0.5")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			response.SomethingWentWrong(ctx)
			return
		}
		defer resp.Body.Close()

		type Rate struct {
			From      string  `json:"from"`
			To        string  `json:"to"`
			Rate      float64 `json:"rate"`
			Timestamp int64   `json:"timestamp"`
		}

		type Body struct {
			Rate Rate `json:"rate"`
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		data := Body{}

		err = json.Unmarshal(body, &data)
		if err != nil {
			response.SomethingWentWrong(ctx)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "Successfully converted currency",
			"data":    data.Rate,
		})
	})
}
