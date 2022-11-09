package transactiondto

type RequestTransaction struct {
	UserID    int    `json:"user_id"` 
	DonaturID int    `json:"donatur_id"`
	Donate    int    `json:"donate"`
	Status    string `jspn:"status"`
}

type ResponseTransaction struct {
	ID        int    `json:"id"`
	Donation  string `json:"donation"`
	Recipient string `json:"recipient"`
	Donatur   string `json:"donatur"`
	Donate    int    `json:"donate"`
	Status    string `json:"status"`
}
