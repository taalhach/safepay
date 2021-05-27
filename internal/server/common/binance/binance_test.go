package binance

import (
	"testing"

	"github.com/taalhach/safepay/pkg/items"

	"github.com/stretchr/testify/require"
)

func TestBestBidPricePoint(t *testing.T) {
	t.Run("retrieve bid current point ", func(t *testing.T) {
		price, err := BestBidPricePoint(items.BinanceMarketBTCUSDT)
		require.Nil(t, err)
		require.NotEqual(t, 0, price)
	})
}
