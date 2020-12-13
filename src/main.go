package main

import (
	"github.com/renandeandradevaz/golang-skills/src/worker"
	"time"
)

func main() {
	worker.InitWorker()
	for {
		time.Sleep(1 * time.Minute)
	}
}
