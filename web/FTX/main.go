package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type FtxClient struct {
	Client *http.Client
	Api    string
	Secret []byte
}

type Order struct {
	Market     string  `json:"market"`
	Side       string  `json:"side"`
	Price      float64 `json:"price"`
	Type       string  `json:"type"`
	Size       float64 `json:"size"`
	ReduceOnly bool    `json:"reduceOnly",omitempty`
	Ioc        bool    `json:"ioc",omitempty`
	PostOnly   bool    `json:"postOnly",omitempty`
}

type Market struct {
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	BaseCurrency   string  `json:"baseCurrency",omitempty`
	quoteCurrency  string  `json:"quoteCurrency",omitempty`
	Underlying     string  `json:"underlying"`
	Enable         bool    `json:"enable"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
	Last           float64 `json:"last"`
	PriceIncrement float64 `json:"priceIncrement"`
	sizeIncrement  float64 `json:"sizeIncrement"`
	Restricted     bool    `json:"restricted"`
}

type TTT struct {
	Result  Market `json:"result"`
	Success bool   `json:"success"`
}

var URL = "https://ftx.com/api/"

func (client *FtxClient) signRequest(method string, path string, body []byte) *http.Request {
	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
	signaturePayload := ts + method + "/api/" + path + string(body)
	signature := client.sign(signaturePayload)
	req, _ := http.NewRequest(method, URL+path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", client.Api)
	req.Header.Set("FTX-SIGN", signature)
	req.Header.Set("FTX-TS", ts)
	return req
}

func (client *FtxClient) _get(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("GET", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *FtxClient) _post(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("POST", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *FtxClient) _delete(path string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("DELETE", path, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *FtxClient) getMarkets() (*http.Response, error) {
	return client._get("markets", []byte(""))
}

func (client *FtxClient) deleteOrder(orderId int64) (*http.Response, error) {
	path := "orders/" + strconv.FormatInt(orderId, 10)
	return client._delete(path, []byte(""))
}

func (client *FtxClient) deleteAllOrders() (*http.Response, error) {
	return client._delete("orders", []byte(""))
}

func (client *FtxClient) placeOrder(market string, side string, price float64, _type string, size float64) (*http.Response, error) {
	newOrder := Order{Market: market, Side: side, Price: price, Type: _type, Size: size}
	body, _ := json.Marshal(newOrder)
	resp, err := client._post("orders", body)
	return resp, err
}

func (client *FtxClient) sign(signaturePayload string) string {
	mac := hmac.New(sha256.New, client.Secret)
	mac.Write([]byte(signaturePayload))
	return hex.EncodeToString(mac.Sum(nil))
}

func main() {
	client := FtxClient{Client: &http.Client{}, Api: "API KEY", Secret: []byte("SECRET")}
	/* resp, _ := client.getMarkets()
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body)) */

	resp, _ := client._get("/markets/BTC/USD", []byte(""))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	test := TTT{}
	if err := json.Unmarshal(body, &test); err != nil {
		fmt.Println(err)
	}
	fmt.Println(time.Now().Unix())
}
