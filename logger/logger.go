package logger

import (
	"fmt"

	"github.com/MrNeocore/go-quartz-clock/config"
)

func LogInputSignal() {
	if config.DEBUG_LOGS {
		fmt.Printf("[S] Tick\n")
	}
}

func LogOutputSignal(signal int) {
	if config.DEBUG_LOGS {
		tabs := ntabs(signal)
		fmt.Printf("%s[%d] Tick\n", tabs, signal-1)
	}
}

func ntabs(count int) string {
	out := ""

	for i := 0; i < count; i++ {
		out += "\t"
	}

	return out
}
