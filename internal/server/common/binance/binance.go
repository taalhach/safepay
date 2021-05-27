package binance

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/taalhach/safepay/pkg/items"

	"github.com/adshao/go-binance/v2"
)

func BestBidPricePoint(symbol string) (float64, error) {

	bids, err := bids(symbol)
	if err != nil {
		return 0, err
	}
	priceStr := bids[0].Price
	return strconv.ParseFloat(priceStr, 64)
}

func bids(symbol string) ([]binance.Bid, error) {
	client := binance.NewClient("", "")
	client.HTTPClient = &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.NewDepthService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.Bids) == 0 {
		return nil, items.NoBidsErr
	}

	return resp.Bids, nil
}
