package model

type Page struct {
	Current  uint8 `json:"current"`
	Next     uint8 `json:"next"`
	Previous uint8 `json:"previous"`
	Total    uint8 `json:"total"`
	Size     uint8 `json:"size"`
}

type BaseItem struct {
	Uuid string `json:"uuid"`
}

type BaseOrder struct {
	*BaseItem
	Id     string  `json:"id"`
	Orders []Order `json:"orders"`
}
type Order struct {
	*BaseItem
}

type OrderDetail struct {
	Total     float32     `json:"total"`
	Breakdown []Breakdown `json:"breakdown"`
}

type Total struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

type Breakdown struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Total float32 `json:"total"`
}

type Result struct {
	Pages   Page        `json:"pages"`
	Count   int         `json:"count"`
	Results []BaseOrder `json:"results"`
}
