package api

import (
	"log"

	"github.com/jeff2857/Binance-SDK/internal/client"
)

var bClient = &client.Client{}

// endpoint
const (
	E_SYSTEM_STATUS = "/sapi/v1/system/status"
	E_ALL_CAPITAL   = "/sapi/v1/capital/config/getall"
)

const (
	GET  = "GET"
	POST = "POST"
)

// get system status
func GetSystemStatus() ([]byte, error) {
	res, err := bClient.Request(GET, E_SYSTEM_STATUS, nil, false, false)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get all capital
func GetAllCapital() ([]byte, error) {
	params := make(map[string]interface{})
	params["recvWindow"] = 50000
	res, err := bClient.Request(GET, E_ALL_CAPITAL, params, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
