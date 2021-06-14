package main

import "testing"
import "github.com/jeff2857/Binance-SDK/internal/apikey"

func TestApiKey(t *testing.T) {
	apiKey, secret, err := apikey.Read()
	if err != nil {
		t.Error(err)
	}
	t.Logf("apikey: %s\nsecret: %s", apiKey, secret)
}
