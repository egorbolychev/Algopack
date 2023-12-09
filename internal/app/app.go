package app

import (
	"context"
	"sync"

	"algopack/internal/parse"
	"algopack/pkg/ctxtool"
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

		go func(ticket string) {
			defer wg.Done()

			_, err := parse.ParseTicketData(ctx, ticket)
			if err != nil {
				ctxtool.Logger(ctx).Error(err.Error())
			}
		}(ticket)
	}
	wg.Wait()
}
