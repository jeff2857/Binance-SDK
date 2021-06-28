package quotation

import (
	"log"

	"github.com/jeff2857/Binance-SDK/internal/client"
	"github.com/jeff2857/Binance-SDK/pkg/constants/methods"
)

var bClient = &client.Client{}

const (
	EP_PING          = "/api/v3/ping"
	EP_TIME          = "/api/v3/time"
	EP_EXCHANGE_INFO = "/api/v3/exchangeInfo"
	EP_DEPTH         = "/api/v3/depth"
	EP_TRADES        = "/api/v3/trades"
	EP_HIS_TRADES    = "/api/v3/historicalTrades"
	EP_AGG_TRADES    = "/api/v3/aggTrades"
	EP_KLINES        = "/api/v3/klines"
	EP_AVG_PRICE     = "/api/v3/avgPrice"
	EP_TICKER_24HR   = "/api/v3/ticker/24hr"
	EP_TICKER_PRICE  = "/api/v3/ticker/price"
	EP_TICKER_BOOK   = "/api/v3/ticker/bookTicker"
)

var depthLimits = [8]int{5, 10, 20, 50, 100, 500, 1000, 5000}

// ping api
func PingAPI() ([]byte, error) {
	res, err := bClient.Request(methods.GET, EP_PING, nil, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get server time
func GetServerTime() ([]byte, error) {
	res, err := bClient.Request(methods.GET, EP_TIME, nil, false, false)
	if err != nil {
		log.Fatal(nil, err)
		return nil, err
	}

	return res, nil
}

// get exchange info and rule
func GetExchangeInfo(symbols []string) ([]byte, error) {
	urlPath := EP_EXCHANGE_INFO
	if symbols != nil {
		if len(symbols) == 1 {
			urlPath += "?symbol=" + symbols[0]
		} else {
			symbolAry := "["
			for _, s := range symbols {
				symbolAry += "\"" + s + "\","
			}
			symbolAry = symbolAry[:len(symbolAry)-1] + "]"
			urlPath += "?symbol=" + symbolAry
		}
	}

	res, err := bClient.Request(methods.GET, urlPath, nil, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get depth
func GetDepth(symbol string, limit int) ([]byte, error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["limit"] = 100
	for _, l := range depthLimits {
		if l == limit {
			params["limit"] = limit
		}
	}

	res, err := bClient.Request(methods.GET, EP_DEPTH, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get trades
func GetTrades(symbol string, limit int) ([]byte, error) {
	if limit <= 0 || limit > 1000 {
		limit = 500
	}
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["limit"] = limit

	res, err := bClient.Request(methods.GET, EP_TRADES, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get historical trades
func GetHistoricalTrades(symbol string, limit int, fromId int64) ([]byte, error) {
	if limit <= 0 || limit > 1000 {
		limit = 500
	}
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["limit"] = limit
	if fromId > 0 {
		params["fromId"] = fromId
	}

	res, err := bClient.Request(methods.GET, EP_HIS_TRADES, params, true, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get aggregate trades
func GetAggTrades(symbol string, fromId, startTime, endTime int64, limit int) ([]byte, error) {
	if limit <= 0 || limit > 1000 {
		limit = 500
	}
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["limit"] = limit
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}

	res, err := bClient.Request(methods.GET, EP_AGG_TRADES, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get klines
func GetKLines(symbol string, interval int, startTime int64, endTime int64, limit int) ([]byte, error) {
	if limit <= 0 || limit > 1000 {
		limit = 500
	}
	params := make(map[string]interface{})
	params["symbol"] = symbol
	params["interval"] = interval
	params["limit"] = limit
	if startTime > 0 {
		params["startTime"] = startTime
	}
	if endTime > 0 {
		params["endTime"] = endTime
	}

	res, err := bClient.Request(methods.GET, EP_KLINES, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get average price
func GetAvgPrice(symbol string) ([]byte, error) {
	params := make(map[string]interface{})
	params["symbol"] = symbol

	res, err := bClient.Request(methods.GET, EP_AVG_PRICE, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get 24hr ticker
func GetTicker24hr(symbol string) ([]byte, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}

	res, err := bClient.Request(methods.GET, EP_TICKER_24HR, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get ticker price
func GetTickerPrice(symbol string) ([]byte, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}

	res, err := bClient.Request(methods.GET, EP_TICKER_PRICE, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get ticker book
func GetTickerBook(symbol string) ([]byte, error) {
	params := make(map[string]interface{})
	if symbol != "" {
		params["symbol"] = symbol
	}

	res, err := bClient.Request(methods.GET, EP_TICKER_BOOK, params, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
