package model

type Page struct {
	Current  float64 `json:"current"`
	Next     bool    `json:"next"`
	Previous bool    `json:"previous"`
	Total    float64 `json:"total"`
	Size     float64 `json:"size"`
}

type baseItem struct {
	Uuid string `json:"uuid"`
}

type BaseOrder struct {
	*baseItem
	Id     string  `json:"id"`
	Orders []Order `json:"orders"`
}
type Order struct {
	Uuid string `json:"uuid"`
}

type Result struct {
	Pages   Page    `json:"pages"`
	Count   float64 `json:"count"`
	Results []Order `json:"results"`
}
