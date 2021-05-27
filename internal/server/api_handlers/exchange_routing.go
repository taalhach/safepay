package apihandlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/taalhach/safepay/pkg/forms"

	"github.com/taalhach/safepay/internal/server/common"

	"github.com/labstack/echo"
)

type RoutingExchangeResponse struct {
	BtcAmount float64 `json:"btc_amount"`
	UsdAmount float64 `json:"usd_amount"`
	Exchange  string  `json:"exchange"`
	Price     float64 `json:"price"`
}

func ExchangeRouting(c echo.Context) error {
	amountStr := c.QueryParam("amount")

	if strings.EqualFold(amountStr, "") {
		var resp forms.BasicResponse
		resp.Success = false

		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		var resp forms.BasicResponse
		resp.Message = fmt.Sprintf("invalid amount %s", amountStr)

		return c.JSON(http.StatusUnprocessableEntity, resp)
	}

	var ret RoutingExchangeResponse
	price, cost, exchange, err := common.MinimizedOrderCost(amount)
	common.CheckError(err)

	ret.Price = price
	ret.BtcAmount = amount
	ret.UsdAmount = cost
	ret.Exchange = exchange

	return c.JSON(http.StatusOK, ret)
}
