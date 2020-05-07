package goincheck

type Ticker struct {
	client *Goincheck
}

// 各種最新情報を簡易に取得することができます。
func (a Ticker) All() string {
	return a.client.Request("GET", "api/ticker", "")
}
