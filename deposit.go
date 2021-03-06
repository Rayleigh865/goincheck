package goincheck

type Deposit struct {
	client *Goincheck
}

// You Get Deposit history
func (a Deposit) All(param string) string {
	return a.client.Request("GET", "api/deposit_money", param)
}

// Deposit Bitcoin Faster
func (a Deposit) Fast(id string) string {
	return a.client.Request("POST", "api/deposit_money/"+id+"/fast", `{"id":"`+id+`"}`)
}
