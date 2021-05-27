package bittrex

import (
	"github.com/taalhach/safepay/pkg/items"
	"github.com/toorop/go-bittrex"
)

func BestBidPricePoint(symbol string) (float64, error) {
	bids, err := bids(symbol)
	if err != nil {
		return 0, err
	}

	price, _ := bids[0].Rate.Float64()
	return price, nil
}

func bids(symbol string) ([]bittrex.OrderbV3, error) {
	client := bittrex.New("", "")

	bids, err := client.GetOrderBookBuySell(symbol, 1, "buy")
	if err != nil {
		return nil, err
	}

	if len(bids) == 0 {
		return nil, items.NoBidsErr
	}
	return bids, nil
}
