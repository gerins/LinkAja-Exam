package account

type Account struct {
	AccountNumber string `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}

type Transaction struct {
	SenderAccount   string
	ReceiverAccount string `json:"to_account_number"`
	Amount          int    `json:"amount"`
}
