package main

import (
	"fmt"
	"time"

	"github.com/Israel-Ferreira/currencies-go-producer/jobs"
	"github.com/go-co-op/gocron"
)

func main() {
	fmt.Println("Kafka + Go = : - )")

	sch := gocron.NewScheduler(time.UTC)

	bnj := jobs.BinanceJob{Scheduler: sch}

	bnj.RunTask()

	sch.StartBlocking()

}
