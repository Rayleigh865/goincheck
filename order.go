package goincheck

import (
	"errors"
	"regexp"
)

type Order struct {
	client *Goincheck
}

// Create a order object with given parameters.In live mode, this issues a transaction.
func (a Order) Create(param string) (string, error) {
	buy_btc, _ := regexp.Compile(`{"rate":"[0-9]*.[0-9]*","amount":"[0-9]*.[0-9]*","order_type":"buy","pair":"btc_jpy"}`)
	sel_btc, _ := regexp.Compile(`{"rate":"[0-9]*.[0-9]*","amount":"[0-9]*.[0-9]*","order_type":"sel","pair":"btc_jpy"}`)
	buy_fct, _ := regexp.Compile(`{"rate":"[0-9]*.[0-9]*","amount":"[0-9]*.[0-9]*","order_type":"buy","pair":"fct_jpy"}`)
	sel_fct, _ := regexp.Compile(`{"rate":"[0-9]*.[0-9]*","amount":"[0-9]*.[0-9]*","order_type":"sel","pair":"fct_jpy"}`)
	// If trading pair value is not "btc_jpy" or "fct_jpy", return error.
	if buy_btc.MatchString(param) || sel_btc.MatchString(param) || buy_fct.MatchString(param) || sel_fct.MatchString(param) {
		return a.client.Request("POST", "api/exchange/orders", param), nil
	}
	return "", errors.New(`trading pair value is not "btc_jpy" or "fct_jpy"`)
}

// cancel a created order specified by order id. Optional argument amount is to refund partially.
func (a Order) Cancel(id string) string {
	return a.client.Request("DELETE", "api/exchange/orders/"+id, "")
}

// List charges filtered by params
func (a Order) Opens() string {
	return a.client.Request("GET", "api/exchange/orders/opens", "")
}

// Get Order Transactions
func (a Order) Transactions() string {
	return a.client.Request("GET", "api/exchange/orders/transactions", "")
}
