package model

type TradeData struct {
	Data struct {
		Data [][]interface{} `json:"data"`
	} `json:"data"`
}

type TicketPredict struct {
	Title   string  `json:"title"`
	Predict float64 `json:"predict"`
}

type TicketData struct {
	Ticket []interface{} `json:"ticket"`
}