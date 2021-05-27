package bittrex

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taalhach/safepay/pkg/items"
)

func TestBestBidPricePoint(t *testing.T) {
	t.Run("bittrex best bid", func(t *testing.T) {

		price, err := BestBidPricePoint(items.BittrexMarketBTCUSD)
		require.Nil(t, err)
		require.NotEqual(t, 0, price)
	})
}
