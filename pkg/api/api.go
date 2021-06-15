package api

import (
	"log"

	"github.com/jeff2857/Binance-SDK/internal/client"
)

var bClient = &client.Client{}

// endpoint
const (
	E_SYSTEM_STATUS = "/sapi/v1/system/status"
	E_ALL_COIN      = "/sapi/v1/capital/config/getall"
)

const (
	GET  = "GET"
	POST = "POST"
)

func GetSystemStatus() ([]byte, error) {
	res, err := bClient.Request(GET, E_SYSTEM_STATUS, nil, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func GetAllCoin() ([]byte, error) {
	params := make(map[string]interface{})
	params["recvWindow"] = 50000
	res, err := bClient.Request(GET, E_ALL_COIN, params, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
