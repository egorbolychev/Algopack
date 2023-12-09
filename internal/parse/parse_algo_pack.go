package parse

import (
	"algopack/internal/model"
	"algopack/pkg/ctxtool"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func ParseTicketData(title string, wg *sync.WaitGroup, ctx context.Context) []byte {
	defer wg.Done()
	apiUrl := fmt.Sprintf("https://iss.moex.com/iss/datashop/algopack/eq/tradestats/%s.json", title)
	res, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	bodyData, err := io.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)

	var tradeData model.TradeData
	err = json.Unmarshal(bodyData, &tradeData)
	if err != nil {
		panic(err)
	}

	ticketMap := buildingTicketMap(tradeData)
	ctxtool.Logger(ctx).Info("etiquette information " + title + " collected")
	latestTicketDataJSON, err := json.Marshal(ticketMap)
	if err != nil {
		panic(err)
	}
	return latestTicketDataJSON
}

func buildingTicketMap(tradeData model.TradeData) map[string]interface{} {
	ticketMap := make(map[string]interface{})

	fields := []string{
		"tradedate", "tradetime", "secid", "pr_open", "pr_high", "pr_low",
		"pr_close", "pr_std", "vol", "val", "trades", "pr_vwap", "pr_change",
		"trades_b", "trades_s", "val_b", "val_s", "vol_b", "vol_s", "disb",
		"pr_vwap_b", "pr_vwap_s", "SYSTIME",
	}

	lastIndex := len(tradeData.Data.Data) - 1
	for i, field := range fields {
		ticketMap[field] = tradeData.Data.Data[lastIndex][i]
	}
	return ticketMap
}
