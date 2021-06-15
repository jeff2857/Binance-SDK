package wallet_test

import (
	"testing"

	"github.com/jeff2857/Binance-SDK/pkg/api/wallet"
)

func TestWallet(t *testing.T) {
	t.Run("system_status", func(t *testing.T) {
		res, err := wallet.GetSystemStatus()
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%s\n", res)
	})

	t.Run("all_capital", func(t *testing.T) {
		res, err := wallet.GetAllCapital()
		if err != nil {
			t.Error(err)
			return
		}
		t.Logf("%s\n", res)
	})
}
