package model

type TradeData struct {
	Data struct {
		Data [][]interface{} `json:"data"`
	} `json:"data"`
}
