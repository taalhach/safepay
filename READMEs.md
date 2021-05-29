# Safepay-Takehome task

This **Golang** application is REST API and fetches BTC/USD prices from several exchanges and tells which exchange will be suitable exchange to buy bitcoin and required amount of USD. 

### Supported exchanges
- [Binance](https://www.binance.com/)
- [Coinbase](https://pro.coinbase.com/)
- [Bittrex](https://global.bittrex.com/)


## Requirements
- Go 1.15.6 or later


### Frameworks used:

- Echo to build the website: https://github.com/labstack/echo

### Build 
Run make command it will take care of rest
```
make
```

### Help

``./bin/safepay --help``

### Serve api

Run following command this will serve REST api on port 4000

``./bin/safepay serve_api``

### Fetch best price
Run following curl command to get suitable exchange and required USD amount
```
curl http://localhost:4000/exchange-routing?amount=1.32
```

**Note:** You might get *context deadline exceeded* error message which means server was unable to fetch details from exchanges api in 5 seconds and cancels the request. 