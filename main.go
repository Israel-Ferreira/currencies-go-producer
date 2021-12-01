package main

import (
	"fmt"
	"time"

	"github.com/Israel-Ferreira/currencies-go-producer/config"
	"github.com/Israel-Ferreira/currencies-go-producer/jobs"
	"github.com/go-co-op/gocron"
)

func main() {
	fmt.Println("Kafka + Go = : - )")

	config.SetProducer("localhost:9092")

	sch := gocron.NewScheduler(time.UTC)

	bnj := jobs.BinanceJob{Scheduler: sch}

	bnj.RunTask()

	sch.StartBlocking()

}
