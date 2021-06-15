package wallet

import (
	"log"

	"github.com/jeff2857/Binance-SDK/internal/client"
	"github.com/jeff2857/Binance-SDK/pkg/constants/methods"
)

var bClient = &client.Client{}

// endpoint
const (
	EP_SYSTEM_STATUS         = "/sapi/v1/system/status"
	EP_ALL_CAPITAL           = "/sapi/v1/capital/config/getall"
	EP_ACCOUNT_SNAPSHOT      = "/sapi/v1/accountSnapshot"
	EP_ENABLE_FAST_WITHDRAW  = "/sapi/v1/account/enableFastWithdrawSwitch"
	EP_DISABLE_FAST_WITHDRAW = "/sapi/v1/account/disableFastWithdrawSwitch"
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
	params["recvWindow"] = 5000
	res, err := bClient.Request(methods.GET, EP_ALL_CAPITAL, params, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// get account snapshot
// t type "SPOT" | "MARGIN" | "FUTURES"
func GetAccountSnapshot(t string) ([]byte, error) {
	params := make(map[string]interface{})
	params["type"] = t
	res, err := bClient.Request(methods.GET, EP_ACCOUNT_SNAPSHOT, params, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// enable fast withdraw
func EnableFastWithdraw() ([]byte, error) {
	res, err := bClient.Request(methods.POST, EP_ENABLE_FAST_WITHDRAW, nil, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

// disable fast withdraw
func DisableFastWithdraw() ([]byte, error) {
	res, err := bClient.Request(methods.POST, EP_DISABLE_FAST_WITHDRAW, nil, true, true)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
