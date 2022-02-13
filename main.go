package main

import (
	"fmt"
	"time"
)

const INPUT_SIGNAL_FREQUENCY = 62500 * time.Microsecond
const FLIP_FLOPS_COUNT = 4
const RUN_DURATION = 10 * time.Second

func ntabs(count int) string {
	out := ""

	for i := 0; i < count; i++ {
		out += "\t"
	}

	return out
}

func firstFlipFlop(input <-chan time.Time, output chan time.Time) {
	for {
		<-input
		fmt.Printf("[S] Tick\n") // Log the original signal
		<-input
		fmt.Printf("[S] Tick\n")
		output <- time.Now()
		fmt.Printf("\t[0] Tick\n")
	}
}

func flipFlop(i int, input chan time.Time, output chan time.Time) {
	for {
		<-input
		<-input
		output <- time.Now()
		tabs := ntabs(i + 1)
		fmt.Printf("%s[%d] Tick\n", tabs, i)
	}
}

func consumeOutputSignal(input chan time.Time) {
	for {
		<-input
	}
}

func main() {
	var chans [FLIP_FLOPS_COUNT]chan time.Time

	for i := range chans {
		chans[i] = make(chan time.Time)
	}

	// Input signal
	signal := time.Tick(INPUT_SIGNAL_FREQUENCY)

	// First flip-flop
	go firstFlipFlop(signal, chans[0])

	// Other flip-flops
	for i := 0; i < FLIP_FLOPS_COUNT-1; i++ {
		go flipFlop(i+1, chans[i], chans[i+1])
	}

	// Consume output signal
	go consumeOutputSignal(chans[FLIP_FLOPS_COUNT-1])

	time.Sleep(RUN_DURATION)
}
