package common

import (
	"math"

	"github.com/taalhach/safepay/internal/server/common/bittrex"

	"github.com/taalhach/safepay/internal/server/common/binance"
	"github.com/taalhach/safepay/internal/server/common/coinbase"
	"github.com/taalhach/safepay/pkg/items"
)

func MinimizedOrderCost(amount float64) (price float64, cost float64, exchange string, err error) {

	binancePrice, err := binance.BestBidPricePoint(items.BinanceMarketBTCUSDT)
	if err != nil {
		return 0, 0, "", err
	}

	coinbasePrice, err := coinbase.BestBidPricePoint(items.CoinbaseMarketBTCUSD)
	if err != nil {
		return 0, 0, "", err
	}

	bittrexPrice, err := bittrex.BestBidPricePoint(items.BittrexMarketBTCUSD)
	if err != nil {
		return 0, 0, "", err
	}

	price = math.Min(binancePrice, coinbasePrice)

	price = math.Min(price, binancePrice)

	switch price {
	case binancePrice:
		exchange = items.Binance
	case coinbasePrice:
		exchange = items.Coinbase
	case bittrexPrice:

	}

	return price, math.Round(price * amount), exchange, nil
}
