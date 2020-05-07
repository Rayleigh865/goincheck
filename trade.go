package goincheck

type Trade struct {
	client *Goincheck
}

// 最新の取引履歴を取得できます。
func (a Trade) All() string {
	return a.client.Request("GET", "api/trades", "")
}
