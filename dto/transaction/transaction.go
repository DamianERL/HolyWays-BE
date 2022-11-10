package transactiondto

type RequestTransaction struct {
	UserID int    `json:"donatur_id"`
	FundID int    `json:"fund_id"`
	Donate int    `json:"donate"`
	Status string `jspn:"status"`
}

type ResponseTransaction struct {
	ID     int    `json:"id"`
	Fund   string `json:"fund"`
	User   string `json:"user"`
	Donate int    `json:"donate"`
	Status string `json:"status"`
}
