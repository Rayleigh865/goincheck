package goincheck

type Leverage struct {
	client *Goincheck
}

// Get a leverage positions list.
func (a Leverage) Positions() string {
	return a.client.Request("GET", "api/exchange/leverage/positions", "")
}
