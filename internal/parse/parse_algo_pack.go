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

func ParseTicketData(ctx context.Context, title string, wg *sync.WaitGroup) ([]byte, error) {
	defer wg.Done()
	apiUrl := fmt.Sprintf("https://iss.moex.com/iss/datashop/algopack/eq/tradestats/%s.json", title)
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(res.Body)
	bodyData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var tradeData model.TradeData
	err = json.Unmarshal(bodyData, &tradeData)
	if err != nil {
		return nil, err
	}

	ctxtool.Logger(ctx).Info("etiquette information " + title + " collected")
	latestTicketDataJSON, err := json.Marshal(tradeData.Data.Data[len(tradeData.Data.Data)-1])
	if err != nil {
		return nil, err
	}
	return latestTicketDataJSON, nil
}
