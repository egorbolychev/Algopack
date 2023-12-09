package app

import (
	"algopack/pkg/ctxtool"
	"context"
	"sync"

	"algopack/internal/parse"
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
		ticket := ticket
		go func() {
			_, err := parse.ParseTicketData(ctx, ticket, &wg)
			if err != nil {
				ctxtool.Logger(ctx).Error(err.Error())
			}
		}()
	}
	wg.Wait()
}
