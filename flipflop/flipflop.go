package flipflop

import (
	"time"

	"github.com/MrNeocore/go-quartz-clock/logger"
)

func FirstFlipFlop(input <-chan time.Time, output chan time.Time) {
	for {
		<-input
		logger.LogInputSignal()

		<-input
		logger.LogInputSignal()

		output <- time.Now()
		logger.LogOutputSignal(1)
	}
}

func FlipFlop(i int, input chan time.Time, output chan time.Time) {
	for {
		<-input
		<-input
		output <- time.Now()
		logger.LogOutputSignal(i)
	}
}
