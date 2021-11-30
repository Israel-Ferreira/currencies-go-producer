package domain

type Ticker struct {
	Symbol string `json:"symbol"`
	PriceChange string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	LastPrice string `json:"lastPrice"`
	PrevClosePrice string `json:"prevClosePrice"`
}