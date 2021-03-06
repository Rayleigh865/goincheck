package goincheck

type BankAccount struct {
	client *Goincheck
}

// Create a new BankAccount.
func (a BankAccount) Create(param string) string {
	return a.client.Request("POST", "api/bank_accounts", param)
}

// Get account information.
func (a BankAccount) All() string {
	return a.client.Request("GET", "api/bank_accounts", "")
}

// Delete a BankAccount.
func (a BankAccount) Delete(id string) string {
	return a.client.Request("DELETE", "api/bank_accounts/"+id, "")
}
