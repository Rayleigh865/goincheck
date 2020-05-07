package goincheck

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Goincheck struct {
	AccessKey   string
	SecretKey   string
	Account     Account
	BankAccount BankAccount
	Borrow      Borrow
	Deposit     Deposit
	Leverage    Leverage
	Order       Order
	OrderBook   OrderBook
	Send        Send
	Ticker      Ticker
	Trade       Trade
	Transfer    Transfer
	Withdraw    Withdraw
}

func (g Goincheck) NewClient(accessKey string, secretKey string) Goincheck {
	g.AccessKey = accessKey
	g.SecretKey = secretKey
	g.Account = Account{&g}
	g.BankAccount = BankAccount{&g}
	g.Borrow = Borrow{&g}
	g.Deposit = Deposit{&g}
	g.Leverage = Leverage{&g}
	g.Order = Order{&g}
	g.OrderBook = OrderBook{&g}
	g.Send = Send{&g}
	g.Ticker = Ticker{&g}
	g.Trade = Trade{&g}
	g.Transfer = Transfer{&g}
	g.Withdraw = Withdraw{&g}
	return g
}

func (g Goincheck) Request(method string, path string, param string) string {
	if param != "" && method == "GET" {
		path = path + "?" + param
		param = ""
	}
	url := "https://coincheck.jp/" + path
	nonce := strconv.FormatInt(CreateNonce(), 10)
	message := nonce + url + param
	req := &http.Request{}
	if method == "POST" {
		payload := strings.NewReader(param)
		req, _ = http.NewRequest(method, url, payload)
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	signature := ComputeHmac256(message, g.SecretKey)
	req.Header.Add("access-key", g.AccessKey)
	req.Header.Add("access-nonce", nonce)
	req.Header.Add("access-signature", signature)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	return string(body)
}

//create nonce by milliseconds
func CreateNonce() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

//create signature
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
