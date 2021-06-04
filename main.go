package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	BASE_URL          = "https://api.binance.com"
	API_KEY           = "**"
	API_CAPITAL_ALL   = "/sapi/v1/capital/config/getall"
	SECRET_KEY        = "**"
	API_SNAPSHOT      = "/sapi/v1/accountSnapshot"
	API_EXCHANGE_INFO = "/api/v3/exchangeInfo"
	API_DEPTH         = "/api/v3/depth"
)

func main() {
	apiTest()
}

type Client struct {
	baseUrl string
	secret  string
	apiKey  string
	clnt    *http.Client
}

func apiTest() {
	client := &Client{}
	// getExchangeInfo(client)
	getDepth(client)
}

func getExchangeInfo(client *Client) {
	body, err := client.request("GET", API_EXCHANGE_INFO, nil, false, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", body)
}

func getDepth(client *Client) {
	params := make(map[string]interface{})
	params["symbol"] = "ETHUSDT"
	params["limit"] = 5

	body, err := client.request("GET", API_DEPTH, params, false, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", body)
}

func (client *Client) request(method string, endpoint string, params map[string]interface{}, withApi bool, withSec bool) ([]byte, error) {
	// init client
	if client.clnt == nil {
		client.clnt = &http.Client{}
		client.secret = SECRET_KEY
		client.baseUrl = BASE_URL
		client.apiKey = API_KEY
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

	if withApi {
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
