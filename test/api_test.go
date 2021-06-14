package main

import (
	"testing"

	"github.com/jeff2857/Binance-SDK/pkg/api"
)

func TestApi(t *testing.T) {
	t.Run("system_status", func(t *testing.T) {
		res, err := api.GetSystemStatus()
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%s\n", res)
	})

	t.Run("all_coin", func(t *testing.T) {
		res, err := api.GetAllCoin()
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%s\n", res)
	})
}
