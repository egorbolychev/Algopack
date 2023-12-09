package model

type TradeData struct {
	Data struct {
		Metadata struct {
			Tradedate struct {
				Type    string `json:"type"`
				Bytes   int    `json:"bytes"`
				MaxSize int    `json:"max_size"`
			} `json:"tradedate"`
			Tradetime struct {
				Type    string `json:"type"`
				Bytes   int    `json:"bytes"`
				MaxSize int    `json:"max_size"`
			} `json:"tradetime"`
			Secid struct {
				Type    string `json:"type"`
				Bytes   int    `json:"bytes"`
				MaxSize int    `json:"max_size"`
			} `json:"secid"`
			PrOpen struct {
				Type string `json:"type"`
			} `json:"pr_open"`
			PrHigh struct {
				Type string `json:"type"`
			} `json:"pr_high"`
			PrLow struct {
				Type string `json:"type"`
			} `json:"pr_low"`
			PrClose struct {
				Type string `json:"type"`
			} `json:"pr_close"`
			PrStd struct {
				Type string `json:"type"`
			} `json:"pr_std"`
			Vol struct {
				Type string `json:"type"`
			} `json:"vol"`
			Val struct {
				Type string `json:"type"`
			} `json:"val"`
			Trades struct {
				Type string `json:"type"`
			} `json:"trades"`
			PrVwap struct {
				Type string `json:"type"`
			} `json:"pr_vwap"`
			PrChange struct {
				Type string `json:"type"`
			} `json:"pr_change"`
			TradesB struct {
				Type string `json:"type"`
			} `json:"trades_b"`
			TradesS struct {
				Type string `json:"type"`
			} `json:"trades_s"`
			ValB struct {
				Type string `json:"type"`
			} `json:"val_b"`
			ValS struct {
				Type string `json:"type"`
			} `json:"val_s"`
			VolB struct {
				Type string `json:"type"`
			} `json:"vol_b"`
			VolS struct {
				Type string `json:"type"`
			} `json:"vol_s"`
			Disb struct {
				Type string `json:"type"`
			} `json:"disb"`
			PrVwapB struct {
				Type string `json:"type"`
			} `json:"pr_vwap_b"`
			PrVwapS struct {
				Type string `json:"type"`
			} `json:"pr_vwap_s"`
			SYSTIME struct {
				Type    string `json:"type"`
				Bytes   int    `json:"bytes"`
				MaxSize int    `json:"max_size"`
			} `json:"SYSTIME"`
		} `json:"metadata"`
		Data [][]interface{} `json:"data"`
	} `json:"data"`
}
