package items

import "errors"

const (
	Binance              = "Binance"
	Coinbase             = "Coinbase"
	Bittrex              = "Bittrex"
	BinanceMarketBTCUSDT = "BTCUSDT"
	CoinbaseMarketBTCUSD = "BTC-USD"
	BittrexMarketBTCUSD  = "BTC-USD"
)

var NoBidsErr = errors.New("order book has no bids")
