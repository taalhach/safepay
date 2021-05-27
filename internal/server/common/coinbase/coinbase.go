package coinbase

import (
	"net/http"
	"strconv"
	"time"

	"github.com/taalhach/safepay/pkg/items"

	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func BestBidPricePoint(symbol string) (float64, error) {

	bids, err := bids(symbol)
	if err != nil {
		return 0, err
	}

	priceStr := bids[0].Price
	return strconv.ParseFloat(priceStr, 64)
}

func bids(symbol string) ([]coinbasepro.BookEntry, error) {
	client := coinbasepro.NewClient()

	client.HTTPClient = &http.Client{
		Timeout: 3 * time.Second,
	}

	book, err := client.GetBook(symbol, 1)
	if err != nil {
		return nil, err
	}

	if len(book.Bids) == 0 {
		return nil, items.NoBidsErr

	}
	return book.Bids, nil
}
