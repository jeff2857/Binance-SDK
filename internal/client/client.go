package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"../apikey"
)

const (
	BASE_URL = "https://api.binance.com"
)

type Client struct {
	baseUrl string
	apiKey  string
	secret  string
	clnt    *http.Client
}

func (client *Client) request(method string, endpoint string, params map[string]interface{}, withApiKey bool, withSec bool) ([]byte, error) {
	if client.clnt == nil {
		apiKey, secret, err := apikey.Read()
		if err != nil {
			return nil, err
		}
		client.clnt = &http.Client{}
		client.baseUrl = BASE_URL
		client.apiKey = apiKey
		client.secret = secret
	}

	req, err := http.NewRequest(method, client.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	urlValues := make(url.Values)
	if params != nil {
		for k, v := range params {
			urlValues.Set(k, fmt.Sprintf("%v", v))
		}
	}
	req.URL.RawQuery = fmt.Sprintf("%s", urlValues.Encode())

	if withSec {
		mac := hmac.New(sha256.New, []byte(client.secret))
		mac.Write([]byte(urlValues.Encode()))
		signature := fmt.Sprintf("%x", mac.Sum(nil))

		timestamp := time.Now().UnixNano() / 1e6

		v := make(url.Values)
		v.Set("signature", signature)
		sigParams := fmt.Sprintf("&timestamp=%d&%s", timestamp, v.Encode())
		req.URL.RawQuery += sigParams
	}

	if withApiKey {
		req.Header.Set("X-MBX-APIKEY", client.apiKey)
	}

	log.Println(req.URL.String())

	resp, err := client.clnt.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
