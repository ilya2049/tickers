package tickers

import (
	"time"
)

type SkipTicker struct {
	ticker *time.Ticker
	c      chan time.Time
}

func NewSkipTicker(d time.Duration) *SkipTicker {
	timeTicker := time.NewTicker(d)

	skipTicker := SkipTicker{
		ticker: timeTicker,
		c:      make(chan time.Time, 1),
	}

	go func() {
		for t := range timeTicker.C {
			skipTicker.c <- t
		}
	}()

	return &skipTicker
}

func (t *SkipTicker) C() <-chan time.Time {
	return t.c
}

func (t *SkipTicker) Stop() {
	close(t.c)

	t.ticker.Stop()
}

func (t *SkipTicker) Reset(d time.Duration) {
	t.ticker.Reset(d)
}

func (t *SkipTicker) Skip() {
	t.c <- time.Now()
}
