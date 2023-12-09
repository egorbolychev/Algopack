package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"algopack/internal/model"
)

func GetPredictByTicket(ctx context.Context, latestTicketData []interface{}) ([]byte, error) {
	ticketData := model.TicketData{
		Ticket: latestTicketData,
	}

	jsonData, err := json.Marshal(ticketData)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(jsonData)
	apiUrl := fmt.Sprintf("http://localhost:8080/api-data")
	res, err := http.Post(apiUrl, "application/json", buffer)
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

	var ticketPredict model.TicketPredict
	err = json.Unmarshal(bodyData, &ticketPredict)
	if err != nil {
		return nil, err
	}

	ticketPredictJSON, err := json.Marshal(ticketPredict)
	if err != nil {
		return nil, err
	}
	return ticketPredictJSON, nil
}
