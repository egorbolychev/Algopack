package app

import (
	"algopack/internal/parse"
	"context"
	"sync"
)

func TradingIteration(ctx context.Context) {
	tickets := []string{"SBER", "MOEX", "MGNT"}
	var wg sync.WaitGroup

	for _, ticket := range tickets {
		wg.Add(1)
		go parse.ParseTicketData(ticket, &wg, ctx)
	}
	wg.Wait()
}
