package clock

import (
	"fmt"
	"time"

	"github.com/MrNeocore/go-quartz-clock/config"
	"github.com/MrNeocore/go-quartz-clock/flipflop"
	"github.com/MrNeocore/go-quartz-clock/logger"
)

func RunClock(inputSignal <-chan time.Time, forDuration time.Duration) {
	startClock(inputSignal)

	// Add some additionnal time to allow the last signal to show up
	duration := forDuration*time.Second + time.Duration(100)*time.Millisecond
	time.Sleep(duration)
}

func startClock(inputSignal <-chan time.Time) {
	var chans [config.FLIP_FLOPS_COUNT]chan time.Time

	for i := range chans {
		chans[i] = make(chan time.Time)
	}

	// First flip-flop
	go flipflop.FirstFlipFlop(inputSignal, chans[0])

	// Other flip-flops
	for i := 0; i < config.FLIP_FLOPS_COUNT-1; i++ {
		go flipflop.FlipFlop(i+1, chans[i], chans[i+1])
	}

	// Consume output signal
	go consumeOutputSignal(chans[config.FLIP_FLOPS_COUNT-1])
}

func consumeOutputSignal(input chan time.Time) {
	for {
		second := <-input
		logger.LogOutputSignal(config.FLIP_FLOPS_COUNT)
		fmt.Printf("=> %v\n", second)
	}
}
