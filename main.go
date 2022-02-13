package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/MrNeocore/go-quartz-clock/clock"
	"github.com/MrNeocore/go-quartz-clock/config"
)

func parseCli() (duration time.Duration) {
	durationPtr := flag.Int("durationSeconds", 5, "Number of seconds to run the quartz clock")

	flag.Parse()

	return time.Duration(*durationPtr)
}

func main() {
	duration := parseCli()

	fmt.Println("Starting")

	inputSignal := time.NewTicker(config.INPUT_SIGNAL_FREQUENCY).C
	clock.RunClock(inputSignal, duration)

	fmt.Println("Exiting")
}
