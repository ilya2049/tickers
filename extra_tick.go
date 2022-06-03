package tickers

import (
	"time"
)

type ExtraTickTicker struct {
	ticker *time.Ticker
	c      chan time.Time
}

func NewExtraTickTicker(d time.Duration) *ExtraTickTicker {
	timeTicker := time.NewTicker(d)

	extraTickTicker := ExtraTickTicker{
		ticker: timeTicker,
		c:      make(chan time.Time, 1),
	}

	go func() {
		for t := range timeTicker.C {
			extraTickTicker.c <- t
		}
	}()

	return &extraTickTicker
}

func (t *ExtraTickTicker) C() <-chan time.Time {
	return t.c
}

func (t *ExtraTickTicker) Stop() {
	t.ticker.Stop()
}

func (t *ExtraTickTicker) Reset(d time.Duration) {
	t.ticker.Reset(d)
}

func (t *ExtraTickTicker) AddExtraTick() {
	go func() {
		select {
		case t.c <- time.Now():
		default:
		}
	}()
}
