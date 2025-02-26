package models

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type ID struct {
	ID string `json:"id"`
}

type Points struct {
	Points int64 `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
