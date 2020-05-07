# goincheck

## Installation

```bash
go get github.com/Rayleigh865/goincheck
```

## Sample Code

```GoLang
package main

import (
	"fmt"
	"github.com/Rayleigh865/goincheck"
)

func main() {
	client := new(goincheck.Goincheck).NewClient("ACCESS_KEY", "API_SECRET")
	/** Public API */
	client.Ticker.All()
	client.Trade.All()
	client.OrderBook.All()

	/** Private API */
	// New order(If trading pair value is not "btc_jpy" or "fct_jpy", return error.)
	body, _ := client.Order.Create(`{"rate":"28500","amount":"0.00508771", "order_type":"buy", "pair":"btc_jpy"}`)
	fmt.Println(body)

	// List of pending orders
	client.Order.Opens()

	// Order Cancellation
	client.Order.Cancel("12345")

	//Transaction History
	client.Order.Transactions()

	// Position List
	client.Leverage.Positions()

	// Balance
	client.Account.Balance()

	// Balance of the leveraged count
	client.Account.LeverageBalance()

	// account information
	client.Account.Info()

	// Transfer Bitcoin
	client.Send.Create(`{"address":"1Gp9MCp7FWqNgaUWdiUiRPjGqNVdqug2hY","amount":"0.0002"`)

	// Bitcoin Transfer History
	client.Send.All("currency=BTC")

	// Bitcoin receipt history
	client.Deposit.All("currency=BTC")

	// Bitcoin Fast Deposit
	client.Deposit.Fast("12345")

	// List of bank accounts
	client.BankAccount.All()

	// Register your bank account
	client.BankAccount.Create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)

	// Deleting a bank account
	client.BankAccount.Delete("25621")

	// Withdrawal history
	client.Withdraw.All()

	// Preparing a withdrawal request
	client.Withdraw.Create(`{"bank_account_id":"2222","amount":"50000", "currency":"JPY", "is_fast":"false"}`)

	// Cancellation of withdrawal request
	client.Withdraw.Cancel("12345")

	// Application for a loan
	client.Borrow.Create(`{"amount":"100","currency":"JPY"}`)

	// Current Loan List
	client.Borrow.Matches()

	// Repayment
	client.Borrow.Repay("1135")

	// Transfer to Leverage Count
	client.Transfer.ToLeverage(`{"amount":"100","currency":"JPY"}`)

	// Transfer from a leveraged account
	client.Transfer.FromLeverage(`{"amount":"100","currency":"JPY"}`)
}

```

## License

MIT
