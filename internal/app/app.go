package app

import (
	"algopack/internal/parse"
	"context"
	"sync"
)

const (
	SBER = "SBER"
	MOEX = "MOEX"
	MGNT = "MGNT"
)

var tickets = [...]string{SBER, MOEX, MGNT}

func TradingIteration(ctx context.Context) {
	var wg sync.WaitGroup

	for _, ticket := range tickets {
		wg.Add(1)
		go parse.ParseTicketData(ticket, &wg, ctx)
	}
	wg.Wait()
}
