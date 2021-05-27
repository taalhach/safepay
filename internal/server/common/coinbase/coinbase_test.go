package coinbase

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taalhach/safepay/pkg/items"
)

func TestBestBidPricePoint(t *testing.T) {

	t.Run("coinbase best bid price", func(t *testing.T) {
		price, err := BestBidPricePoint(items.CoinbaseMarketBTCUSD)
		require.Nil(t, err)
		require.NotEqual(t, 0, price)
	})

}
