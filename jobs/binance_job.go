package jobs

import (
	"github.com/Israel-Ferreira/currencies-go-producer/services"
	"github.com/go-co-op/gocron"
)

type BinanceJob struct {
	Scheduler *gocron.Scheduler
}

func (b BinanceJob) configure() {
	b.Scheduler.Every(30).Seconds().Do(func() {
		services.GetBinanceBtcCurrency()
	})

	b.Scheduler.Every(20).Seconds().Do(services.GetBinanceEthCurrency)

}

func (b BinanceJob) RunTask() {
	b.configure()
	// b.Scheduler.StartImmediately()
}
