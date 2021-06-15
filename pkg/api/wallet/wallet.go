package wallet

import (
	"log"

	"github.com/jeff2857/Binance-SDK/internal/client"
	"github.com/jeff2857/Binance-SDK/pkg/constants/methods"
)

var bClient = &client.Client{}

// endpoint
const (
	EP_SYSTEM_STATUS = "/sapi/v1/system/status"
	EP_ALL_CAPITAL   = "/sapi/v1/capital/config/getall"
)

// get system status
func GetSystemStatus() ([]byte, error) {
	res, err := bClient.Request(methods.GET, EP_SYSTEM_STATUS, nil, false, false)
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
	res, err := bClient.Request(methods.GET, EP_ALL_CAPITAL, params, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
